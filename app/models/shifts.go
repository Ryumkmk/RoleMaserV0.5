package models

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
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

type ShiftInfoInChangePage struct {
	WITP  WeddingInTypingPage
	PLITs []PjListInTyping
}

func GetShiftsByDateFromFile(date string, sheetName string) (pj_id []int, shiftTime []string, ampm []string) {

	f, err := excelize.OpenFile("./" + shiftFileName)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	sheetIndex, _ := f.GetSheetIndex(sheetName)
	if sheetIndex == -1 {
		log.Println(sheetIndex)
		return nil, nil, nil
	} else {

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
				break
			}
		}
		cols, err := f.GetCols(sheetName)
		if err != nil {
			log.Println(err)
		}

		col := cols[dateColNum]
		pjidcol := cols[pjidcolIndex]
		pjidcol = pjidcol[3:]
		for i, v := range col {
			matches := re.FindStringSubmatch(v)
			if len(matches) > 0 {
				id, _ := strconv.Atoi(pjidcol[i-3])
				pj_id = append(pj_id, id)
				stj := strings.Join(matches, "")
				shiftTime = append(shiftTime, stj)
				ampm = append(ampm, isAmpm(stj))
			}
		}
	}
	return pj_id, shiftTime, ampm
}

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
		}
		// fmt.Println("Doing...")
	}
}

func (s *Shift) insertRowToShiftsDB() (err error) {
	cmd := `insert into shifts (date,pj_id,shifttime,ampm) values (?,?,?,?)`
	_, err = Db.Exec(cmd, s.Date, s.Pj_ID, s.ShiftTime, s.Ampm)
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetAllShiftByName(name string, month string) (shifts []Shift, err error) {
	cmd := `SELECT 
				date_format(date,'%m月%d日'),s.shifttime,s.ampm
			FROM
				shifts AS s
					INNER JOIN
				pjs AS p ON p.id = s.pj_id
			WHERE
				p.name = ? and date_format(s.date,'%m') = ?
				order by s.date;`
	rows, err := Db.Query(cmd, name, month)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var shift Shift
		rows.Scan(
			&shift.Date,
			&shift.ShiftTime,
			&shift.Ampm,
		)
		shifts = append(shifts, shift)
	}
	return shifts, err

}
func ChangeDateFormatToDBFormat(year string, month string, day string) (date string) {

	if len(month) == 1 {
		month = "0" + month
	}
	if len(day) == 1 {
		day = "0" + day
	}
	return fmt.Sprintf("%s-%s-%s", year, month, day)
}

func GetAllPjsShiftByDateFromDB(date string, year string, month string, day string) (pLITs []PjListInTyping, err error) {
	cmd := `SELECT 
				p.name,
				s.shifttime,
                s.ampm
			FROM
				shifts AS s
					JOIN
				pjs AS p ON p.id = s.pj_id
			WHERE
				s.date = ?
				order by p.id;`
	rows, err := Db.Query(cmd, date)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var pLIT PjListInTyping
		rows.Scan(
			&pLIT.Name,
			&pLIT.ShiftTime,
			&pLIT.Ampm)
		pLITs = append(pLITs, pLIT)
	}
	return pLITs, err
}

func DeletePjShiftFromDB(date string, pjname string) (err error) {
	cmd := `DELETE FROM shifts 
			WHERE
				(date , pj_id) IN (SELECT 
					date, pj_id
				FROM
					(SELECT 
						s.date, p.id AS pj_id
					FROM
						shifts AS s
					JOIN pjs AS p ON p.id = s.pj_id
					
					WHERE
						s.date = ? AND p.name = ?) AS subquery);`

	_, err = Db.Exec(cmd, date, pjname)
	if err != nil {
		log.Println(err)
	}
	return err
}

func AddPjShiftFromDB(date string, pjname string, shifttime string, ampm string) (err error) {
	cmd := `insert shifts (date,pj_id,shifttime,ampm) values (
				?,
				(select p.id from pjs as p where p.name = ?),
				?,
				?);`

	_, err = Db.Exec(cmd, date, pjname, shifttime, ampm)
	if err != nil {
		log.Println(err)
	}
	return err
}
