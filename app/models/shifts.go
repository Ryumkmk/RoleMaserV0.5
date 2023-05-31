package models

import (
	"log"
	"regexp"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Shift struct {
	ID        int
	Date      string
	Pj_ID     int
	ShiftTime string
	Ampm      string
}

// 日付からシフト情報を取得、返り値は（出勤PJID、出勤時間、AMPM）
func GetShiftsByDateFromFile(date string, sheetName string) (pj_id []int, shiftTime []string, ampm []string) {

	// date = "2023-06-03"
	//xlsx（シフト）ファイルを読み込む
	f, err := excelize.OpenFile("./" + shiftFileName)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	sheetIndex, _ := f.GetSheetIndex(sheetName)
	if sheetIndex == -1 {
		//シートが存在しない場合
		log.Println(sheetIndex)
		return nil, nil, nil
	} else {

		//シートが存在する場合
		re := regexp.MustCompile(`\d{1,2}:\d{2}-\d{1,2}:\d{2}`)
		rows, err := f.GetRows(sheetName)
		if err != nil {
			log.Println(err)
		}
		row := rows[shiftDateRowIndex]
		var dateColNum int
		for i, v := range row {
			if v == date {
				dateColNum = i
			}
		}
		cols, err := f.GetCols(sheetName)
		if err != nil {
			log.Println(err)
		}
		col := cols[dateColNum]
		for i, v := range col {
			matches := re.FindStringSubmatch(v)
			if len(matches) > 0 {
				pj_id = append(pj_id, i-2)
				stj := strings.Join(matches, "")
				shiftTime = append(shiftTime, stj)
				ampm = append(ampm, isAmpm(stj))
			}
		}
	}
	return pj_id, shiftTime, ampm
}

// シート名前からシフトを構造体に登録※呼び出すと、データベースShiftsにその月のシフト情報を登録する
func InsertAllRowsShifts(sheetName string) {
	dates := GetWeddingsDateByFile(sheetName)
	for _, date := range dates {
		pj_ids, shifttimes, ampm := GetShiftsByDateFromFile(date, sheetName)
		for i, pj_id := range pj_ids {
			s := Shift{
				Date:      date,
				Pj_ID:     pj_id,
				ShiftTime: shifttimes[i],
				Ampm:      ampm[i],
			}
			s.insertRowToShiftsDB()
			// fmt.Println(s)
		}
	}
}

// データベースにShiftsを登録する
func (s *Shift) insertRowToShiftsDB() (err error) {
	cmd := `insert into shifts (date,pj_id,shifttime,ampm) values (?,?,?,?)`
	_, err = Db.Exec(cmd, s.Date, s.Pj_ID, s.ShiftTime, s.Ampm)
	if err != nil {
		log.Println(err)
	}
	return err
}
