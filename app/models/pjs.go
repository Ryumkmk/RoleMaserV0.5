package models

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"RMV0.5/app/config"
	"github.com/xuri/excelize/v2"
)

type Pjs struct {
	Date  string
	Names []string
}

func GetPjs(day string) []string {
	f := ReadXlsxFile()
	xf, err := excelize.OpenFile(config.Config.Xlsxpath + "/" + f.Name())
	defer xf.Close()
	if err != nil {
		log.Println(err)
	}
	sheetName := "PJシフト4月"

	allPjsNames := getAllPjsNames(xf, sheetName)
	// fmt.Println(allPjsNames)
	Num := getShiftDayPjNum(xf, sheetName, day)
	pjsNames := make([]string, 0)

	if Num == nil {
		nopj := "Pjを取得出来ませんでした"
		pjsNames = append(pjsNames, nopj)
	} else {
		// fmt.Println(Num)
		for _, v := range Num {
			if v >= 0 {
				pjsNames = append(pjsNames, allPjsNames[v])
			}
		}
	}
	return pjsNames

}

func getAllPjsNames(f *excelize.File, sheetName string) []string {
	cols, err := f.GetCols(sheetName)
	if err != nil {
		log.Println(err)
	}
	pjsNames := cols[2]
	pjsNames = pjsNames[13:]
	return pjsNames
}

func getAllShiftDays(f *excelize.File, sheetName string) []string {
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Println(err)
	}
	row := rows[5]
	return row
}

func getShiftDayPjNum(f *excelize.File, sheetName string, day string) (Nums []int) {
	cols, err := f.GetCols(sheetName)
	if err != nil {
		log.Println(err)
	}
	allShiftDays := getAllShiftDays(f, sheetName)
	dayInt := getShiftColNum(day, allShiftDays)
	if dayInt == -1 {
		return nil
	} else {
		col := cols[dayInt]
		for i, okOrno := range col {
			if len(okOrno) >= 8 {
				Nums = append(Nums, i-13)
			}
		}
		return Nums
	}
}

func getShiftColNum(day string, allShiftDays []string) int {
	for i, shiftDay := range allShiftDays {
		if day == shiftDay {
			return i
		}
	}
	return -1
}

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
