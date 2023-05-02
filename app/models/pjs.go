package models

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
	"strings"

	"RMV0.5/app/config"
	"github.com/xuri/excelize/v2"
)

// Pjとその仕事の構造体
type WhatJob struct {
	Roles    []Role
	Pjs      []Pj
	Trainers []Trainer
}

// Pjの構造体
type Pj struct {
	Date    string //出勤日付
	Names   string //出勤Pj名
	Time    string //出勤時間
	AmPm    string //午前か午後
	CheckAM bool   //AM入力済チェック
	CheckPM bool   //PM入力済チェック

}

// 仕事の構造体
type Role struct {
	RoleName string //仕事の名前
	PjName   string //仕事に割り振られたPjの名前
}

type Trainer struct {
	Key         string //トレイナー、トレイナーのセットがわかるKey
	TrainerName string //トレイナーの名前
	TraineeName string //トレイニーの名前
}

// 日付から出勤するpj一覧を取得する、返り値は(出勤するPjの名前,出勤時間)
func GetPjs(month string, day string) (pjsNames []string, time []string, ampm []string) {

	//xlsx（シフト）ファイルを読み込む
	f := ReadXlsxFile()
	xf, err := excelize.OpenFile(config.Config.Xlsxpath + "/" + f.Name())
	if err != nil {
		log.Println(err)
	}
	defer xf.Close()

	//%s月のシフトシートを取得する
	sheetName := fmt.Sprintf("PJシフト%s月", month)
	sheetIndex, _ := xf.GetSheetIndex(sheetName)
	if sheetIndex == -1 {

		//シートが存在しない場合
		nopj := "Pjを取得出来ませんでした"
		notime := "日付をもう一度確認して下さい"
		noampm := "日付をもう一度確認して下さい"
		pjsNames = append(pjsNames, nopj)
		time = append(time, notime)
		ampm = append(ampm, noampm)
		return pjsNames, time, ampm
	} else {

		//シートが存在する場合
		allPjsNames := getAllPjsNames(xf, sheetName)
		Num, time, ampm := getShiftDayPjNum(xf, sheetName, day)
		if Num == nil {

			//出勤日が存在しない場合
			nopj := "Pjを取得出来ませんでした"
			notime := "日付をもう一度確認して下さい"
			noampm := "日付をもう一度確認して下さい"
			pjsNames = append(pjsNames, nopj)
			time = append(time, notime)
			ampm = append(ampm, noampm)

		} else {

			//出勤日が存在する場合
			for _, v := range Num {
				if v >= 0 {
					pjsNames = append(pjsNames, allPjsNames[v])
				}
			}
		}
		return pjsNames, time, ampm
	}
}

// 全てのPjの名前の文字列スライスを取得する
func getAllPjsNames(f *excelize.File, sheetName string) []string {
	cols, err := f.GetCols(sheetName)
	if err != nil {
		log.Println(err)
	}
	//xlsxファイルの三列目が全てのPj名、そして、14行名から
	pjsNames := cols[2]
	pjsNames = pjsNames[13:]
	return pjsNames
}

// シートから出勤のある日の文字列スライスを取得する
func getAllShiftDays(f *excelize.File, sheetName string) []string {
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Println(err)
	}
	//６行目が出勤日
	row := rows[5]
	return row
}

// 引数の日にちからその日出勤するPjのIdと出勤時間を取得
func getShiftDayPjNum(f *excelize.File, sheetName string, day string) (Nums []int, time []string, ampm []string) {
	cols, err := f.GetCols(sheetName)
	if err != nil {
		log.Println(err)
	}
	allShiftDays := getAllShiftDays(f, sheetName)
	dayInt := getShiftColNum(day, allShiftDays)
	//8:30-16:00のフォーマットの文字列を取得するためのRegex
	re := regexp.MustCompile(`\d{1,2}:\d{2}-\d{1,2}:\d{2}`)
	if dayInt == -1 {

		//出勤日が存在しない場合
		return nil, nil, nil
	} else {

		//出勤日が存在する場合
		col := cols[dayInt]
		for i, v := range col {
			matches := re.FindStringSubmatch(v)
			if len(matches) > 0 {
				Nums = append(Nums, i-13)
				//出勤時間を一つの文字列に
				timeString := strings.Join(matches, "")
				//-が現れる位置を取得
				index := strings.Index(timeString, "-")
				//出勤時間と退勤時間を取得
				startString := strings.TrimSpace(timeString[:index])
				endString := strings.TrimSpace(timeString[index+1:])
				startT := string(startString[0])
				endT := string(endString[0])
				// fmt.Printf("%v,%v\n", startT, endT)
				if (startT == "8" || startT == "9" || startT == "0") && (endT == "2") {
					ampm = append(ampm, "ダブル")
				} else if (startT == "8" || startT == "9" || startT == "0") && (endT == "1") {
					ampm = append(ampm, "AM")
				} else {
					ampm = append(ampm, "PM")
				}
				time = append(time, timeString)
			}
		}
		// fmt.Println(time)
		return Nums, time, ampm
	}
}

// 出勤する日の列番号を取得する
func getShiftColNum(day string, allShiftDays []string) int {
	for i, shiftDay := range allShiftDays {
		if day == shiftDay {
			return i
		}
	}
	return -1
}

// ファイルを読み込む
func ReadXlsxFile() (f fs.DirEntry) {
	files, err := os.ReadDir(config.Config.Xlsxpath)
	if err != nil {
		fmt.Println(err)
	}
	for _, f = range files {
		if strings.HasSuffix(f.Name(), ".xlsx") {
			return f
		}
	}
	return nil
}

// 入力済みかチェックする
func IsInputPjs(n []string, j *WhatJob) {
	var amExistPjs []string
	var pmExistPjs []string
	var dubleExsitPjs []string
	for _, v := range j.Roles {
		lastRoleName := string(v.RoleName[len(v.RoleName)-1])
		if lastRoleName != "P" {
			for _, v2 := range n {
				if strings.Contains(v.PjName, v2) {
					amExistPjs = append(amExistPjs, v2)
				}
			}
		} else {
			for _, v2 := range n {
				if strings.Contains(v.PjName, v2) {
					pmExistPjs = append(pmExistPjs, v2)
				}
			}
		}
	}
	for _, v := range amExistPjs {
		for _, v2 := range pmExistPjs {
			if v == v2 {
				dubleExsitPjs = append(dubleExsitPjs, v)
			}
		}
	}
	for _, v := range dubleExsitPjs {
		for i, v2 := range j.Pjs {
			if v == v2.Names {
				j.Pjs[i].CheckAM = true
				j.Pjs[i].CheckPM = true
			}
		}
	}
	for _, v := range amExistPjs {
		for i, v2 := range j.Pjs {
			if v == v2.Names {
				j.Pjs[i].CheckAM = true
			}
		}
	}
	for _, v := range pmExistPjs {
		for i, v2 := range j.Pjs {
			if v == v2.Names {
				j.Pjs[i].CheckPM = true
			}
		}
	}
}

// トレイナー、トレイニーをセットで登録する
func WhosTrainee(key string, trainer_neeName string, trainers *[]Trainer) {
	// fmt.Printf("key=%v,name = %v\n", key, trainer_neeName)
	if strings.Contains(key, "trainer") {
		if len(*trainers) == 0 {
			var trainer Trainer
			trainer.Key = key
			trainer.TrainerName = trainer_neeName
			*trainers = append(*trainers, trainer)
			return
		}
		for i, v := range *trainers {
			if v.Key == key {
				(*trainers)[i].TrainerName = trainer_neeName
				return
			}
		}
		var trainer Trainer
		trainer.Key = key
		trainer.TrainerName = trainer_neeName
		*trainers = append(*trainers, trainer)
		// fmt.Println(trainers)
		return
	} else {
		key = strings.Replace(key, "nee", "ner", 1)
		if len(*trainers) == 0 {
			var trainer Trainer
			trainer.Key = key
			trainer.TraineeName = trainer_neeName
			*trainers = append(*trainers, trainer)
			return
		}
		for i, v := range *trainers {
			if v.Key == key {
				(*trainers)[i].TraineeName = trainer_neeName
				return
			}
		}
		var trainer Trainer
		trainer.Key = key
		trainer.TraineeName = trainer_neeName
		*trainers = append(*trainers, trainer)
		// fmt.Println(trainers)
		return
	}
}
