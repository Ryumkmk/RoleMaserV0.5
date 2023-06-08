package models

import (
	"log"
	"regexp"
	"strings"
)

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
	RCS      []RoleCount
	RICPs    []RestInCheckPage
}

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

func isAmpm(shifttime string) (ampm string) {
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

func splitPjsInSameRole(input string) (result []string) {
	re := regexp.MustCompile(`[\p{Hiragana}\p{Katakana}\wー\s]+`)
	matches := re.FindAllString(input, -1)
	for _, match := range matches {
		pjs := strings.Fields(match)
		result = append(result, pjs...)
	}
	return result
}

func GetPjInfoFromDB(pjname string) (pj Pj, err error) {

	cmd := `SELECT 
				*
			FROM
				pjs AS p
			WHERE
				p.name = ?;`

	err = Db.QueryRow(cmd, pjname).Scan(
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
	return pj, err
}

func (p *Pj) UpdatePjDb() (err error) {
	cmd := `UPDATE pjs 
			SET 
			level = ?,
			gatekeeper = ?,
			toilet = ?,
			cloak = ?,
			silver = ?,
			wash = ?,
			ape = ?,
			champagne = ?,
			drinkcounter = ?,
			leader = ?
			where name = ?;`
	_, err = Db.Exec(
		cmd,
		p.Level,
		p.Gatekeeper,
		p.Toilet,
		p.Cloak,
		p.Silver,
		p.Wash,
		p.Ape,
		p.Champagne,
		p.Drinkcounter,
		p.Leader,
		p.Name)
	if err != nil {
		log.Println(err)
	}
	return err
}

func UpdateTrainerPjDB(date string, tTs []TrainerTrainee) (err error) {

	if err = deleteTrainerTraineeDB(date); err != nil {
		log.Println(err)
	}
	for _, tT := range tTs {
		if err = tT.insertTrainerTraineeDB(date); err != nil {
			log.Println(err)
		}
	}
	return err
}

func GetAllTrainersTraineesFromDB(date string) (tTs []TrainerTrainee, err error) {
	cmd := `SELECT 
			(select p.name from pjs as p where t.trainer_pj_id = p.id),
			(select p.name from pjs as p where t.trainee_pj_id = p.id)
		FROM
			trainers as t
			where t.date = ?
			order by id;`
	rows, err := Db.Query(cmd, date)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var tT TrainerTrainee
		rows.Scan(
			&tT.Trainer,
			&tT.Trainee,
		)
		tTs = append(tTs, tT)
	}
	return tTs, err
}

func (tT *TrainerTrainee) insertTrainerTraineeDB(date string) (err error) {
	cmd := `INSERT INTO trainers (date, trainer_pj_id, trainee_pj_id)
			SELECT ?, p1.id, p2.id
			FROM pjs AS p1, pjs AS p2
			WHERE p1.name = ? AND p2.name = ?
			AND p1.name IS NOT NULL AND p2.name IS NOT NULL;`

	_, err = Db.Exec(cmd, date, tT.Trainer, tT.Trainee)
	if err != nil {
		log.Println(err)
	}
	return err
}

func deleteTrainerTraineeDB(date string) (err error) {
	cmd := `DELETE FROM trainers 
			WHERE
				date = ?;`
	_, err = Db.Exec(cmd, date)
	if err != nil {
		log.Println(err)
	}
	return err
}
