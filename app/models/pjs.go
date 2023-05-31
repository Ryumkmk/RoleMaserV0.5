package models

import (
	"log"
	"regexp"
	"strings"
)

// Pjの構造体
type Pj struct {
	ID           int
	Name         string
	Level        string
	Gatekeeper   string
	Toilet       string
	Cloak        string
	Silver       string
	Wash         string
	Ape          string
	Coffee       string
	Champagne    string
	Drinkcounter string
	Leader       string
}

type PjListInTyping struct {
	Name      string
	Level     string
	ShiftTime string
	Ampm      string
}

type TrainerTrainee struct {
	Key     string
	Trainer string
	Trainee string
}

type DataInTypingPage struct {
	PLITs    []PjListInTyping
	WITPAM   WeddingInTypingPage
	WITPPM   WeddingInTypingPage
	RIITPsAM []RoleInfoInTypingPage
	RIITPsPM []RoleInfoInTypingPage
	TTs      []TrainerTrainee
}

// 全てのPjをデートベースから取得
func GetAllPjsByDB() (pjs []Pj, err error) {

	cmd := `select * from pjs`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var pj Pj
		err = rows.Scan(
			&pj.ID,
			&pj.Name,
			&pj.Level,
			&pj.Gatekeeper,
			&pj.Toilet,
			&pj.Cloak,
			&pj.Silver,
			&pj.Wash,
			&pj.Ape,
			&pj.Coffee,
			&pj.Champagne,
			&pj.Drinkcounter,
			&pj.Leader,
		)
		if err != nil {
			log.Println(err)
		}
		pjs = append(pjs, pj)
	}
	rows.Close()
	return pjs, err
}

// 同じ文字列を文字列スライスから削除
func removeDuplicates(slice []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, item := range slice {
		if !encountered[item] {
			encountered[item] = true
			result = append(result, item)
		}
	}

	return result
}

// AMPM、ダブルか判定する
func isAmpm(shifttime string) (ampm string) {
	//-が現れる位置を取得
	index := strings.Index(shifttime, "-")
	//出勤時間と退勤時間を取得
	startString := strings.TrimSpace(shifttime[:index])
	endString := strings.TrimSpace(shifttime[index+1:])
	startT := string(startString[0])
	endT := string(endString[0])
	endT2 := string(endString[1])
	// fmt.Printf("%v,%v\n", startT, endT)
	if (startT == "8" || startT == "9" || startT == "0") && (endT == "2" || endT2 == "9") {
		return "ダブル"

	} else if (startT == "8" || startT == "9" || startT == "0") && (endT == "1") {
		return "AM"
	} else if (startT == "1") && (endT == "2") {
		return "PM"
	} else {
		return "試食会"
	}
}

// 日付でデータベースから出勤PJ一覧を取得
func GetPjsByDateFromDB(date string) (pLITs []PjListInTyping, err error) {
	cmd := `SELECT 
				p.name,p.level,s.shifttime,s.ampm
			FROM
    			shifts AS s
        	INNER JOIN
    			pjs AS p ON s.pj_id = p.id
			WHERE date = ?;`
	rows, err := Db.Query(cmd, date)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var PLIT PjListInTyping
		rows.Scan(
			&PLIT.Name,
			&PLIT.Level,
			&PLIT.ShiftTime,
			&PLIT.Ampm,
		)
		pLITs = append(pLITs, PLIT)
	}
	return pLITs, err
}

// Weddingの日付とAMPMからその日の役割とそれに対応するPJを取得
func (w *WeddingInTypingPage) GetRoleInfoByDateFromDB() (rIITPs []RoleInfoInTypingPage, err error) {
	cmd := `SELECT 
    			r.name,p.name
			FROM
    			role_info AS ri
        	INNER JOIN
    			weddings AS w ON ri.wedding_id = w.id
        	INNER JOIN
    			pjs AS p ON ri.pj_id = p.id
        	INNER JOIN
    			roles AS r ON ri.role_id = r.id
    		where w.date = ? and w.ampm = ?;`
	rows, err := Db.Query(cmd, w.Date, w.Ampm)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var rIITP RoleInfoInTypingPage
		rows.Scan(
			&rIITP.RoleName,
			&rIITP.PjName,
		)
		rIITPs = append(rIITPs, rIITP)
	}
	return rIITPs, err
}

// 同じ役割に二人以上のPJがいる場合、それを分割する
func splitPjsInSameRole(input string) (result []string) {
	re := regexp.MustCompile(`[\p{Hiragana}\p{Katakana}\wー]+`)
	result = re.FindAllString(input, -1)
	return result
}


