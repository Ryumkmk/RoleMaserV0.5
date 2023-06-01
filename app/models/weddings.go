package models

import (
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

type Wedding struct {
	ID    int
	Date  string
	Ampm  string
	Guest int
}

type WeddingInTypingPage struct {
	Date  string
	Date2 string
	Ampm  string
	Guest int
}

func GetWeddingsByDateFromDB(date string) (wITPs []WeddingInTypingPage) {
	cmd := `SELECT 
    			date, ampm, guest
			FROM
    			weddings
			WHERE
    			date = ?;`
	rows, err := Db.Query(cmd, date)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var wed Wedding
		rows.Scan(
			&wed.Date,
			&wed.Ampm,
			&wed.Guest,
		)
		date22, err := time.Parse("2006-01-02", wed.Date)
		if err != nil {
			log.Println(err)
		}
		date2 := date22.Format("1月2日")
		var wITP = WeddingInTypingPage{
			Date:  wed.Date,
			Date2: date2,
			Ampm:  wed.Ampm,
			Guest: wed.Guest,
		}
		wITPs = append(wITPs, wITP)
	}
	return wITPs
}

func GetWeddingsDateByFile(sheetName string) (date []string) {
	f, err := excelize.OpenFile("./" + shiftFileName)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	sheetIndex, _ := f.GetSheetIndex(sheetName)
	if sheetIndex == -1 {
		log.Println(sheetIndex)
		return nil
	} else {
		rows, err := f.GetRows(sheetName)
		if err != nil {
			log.Println(err)
		}
		row := rows[shiftDateRowIndex]
		re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
		for _, v := range row {
			mathces := re.FindStringSubmatch(v)
			if len(mathces) > 0 {
				date = append(date, mathces...)
			}
		}
	}
	dateResult := removeDuplicates(date)
	return dateResult
}

// シート名前からWeddingsを構造体に登録※呼び出すと、データベースWeddingsにその月の婚礼情報を登録する
func InsertAllRowsWeddings(sheetName string) {

	f, err := excelize.OpenFile("./" + shiftFileName)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	sheetIndex, _ := f.GetSheetIndex(sheetName)
	if sheetIndex == -1 {
		log.Println(sheetIndex)
		return
	} else {

		rows, err := f.GetRows(sheetName)
		if err != nil {
			log.Println(err)
		}
		alldateRow := rows[shiftDateRowIndex]
		dateRow := alldateRow[3:]
		allguestRow := rows[guestNumRowIndex]
		guestRow := allguestRow[3:]
		allampmRow := rows[ampmRowIndex]
		ampmRow := allampmRow[3:]

		for i := 0; i < len(dateRow); i++ {
			guest, _ := strconv.Atoi(guestRow[i])
			w := Wedding{
				Date:  dateRow[i],
				Guest: guest,
				Ampm:  ampmRow[i],
			}
			w.insertRowToWeddingsDB()
		}
	}
}

func (w *Wedding) insertRowToWeddingsDB() (err error) {
	cmd := `insert into weddings (date,ampm,guest) values (?,?,?)`
	_, err = Db.Exec(cmd, w.Date, w.Ampm, w.Guest)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (w *Wedding) InsertRowToWeddings() (wedding_id int, err error) {
	cmd := `SELECT 
   				id
			FROM
    			weddings
			WHERE
    			date = '2023-05-01' AND ampm = 'AM';`
	err = Db.QueryRow(cmd, w.Date, w.Ampm).Scan(
		&wedding_id,
	)
	return wedding_id, err
}
