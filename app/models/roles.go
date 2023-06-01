package models

import (
	"log"
)

// 役割の構造体
type Role struct {
	ID   int
	Name string
}

type RoleCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
type Role_Info struct {
	ID         int
	Wedding_ID int
	Role_ID    int
	Pj_ID      int
}

// 入力ページで使う役割とそのPJを格納する構造体
type RoleInfoInTypingPage struct {
	RoleName string
	PjName   string
}

// 役割情報をデータベースに追加、更新する
func (dITP *DataInTypingPage) UpdateRoleInfoDB() (err error) {

	if len(dITP.WITPAM.Ampm) > 0 && len(dITP.WITPPM.Ampm) > 0 {
		err = dITP.WITPAM.UpdateRoleInfoDBByAmpm(dITP.RIITPsAM)
		if err != nil {
			log.Println(err)
		}
		err = dITP.WITPPM.UpdateRoleInfoDBByAmpm(dITP.RIITPsPM)
		if err != nil {
			log.Println(err)
		}
	} else if len(dITP.WITPAM.Ampm) > 0 {
		err = dITP.WITPAM.UpdateRoleInfoDBByAmpm(dITP.RIITPsAM)
		if err != nil {
			log.Println(err)
		}
	} else {
		err = dITP.WITPPM.UpdateRoleInfoDBByAmpm(dITP.RIITPsPM)
		if err != nil {
			log.Println(err)
		}
	}
	return err
}

// AMPMで分けて役割情報をデータベースに追加、更新する
func (w *WeddingInTypingPage) UpdateRoleInfoDBByAmpm(rIITPs []RoleInfoInTypingPage) (err error) {
	// fmt.Println(rIITPs)

	for _, rITTP := range rIITPs {
		//役割が新しく振り分けられているかどうか
		ok, err := w.isRoleInfoExistInDBByInput(rITTP)
		if err != nil {
			log.Println(err)
		}
		// fmt.Println(ok)
		if !ok {
			// 入力フォームが更新されたので、該当の入力フォームの役割をRoleInfoデータベースから削除する
			err = w.deleteRowToRoleInfoDB(rITTP.RoleName)
			if err != nil {
				log.Println(err)
			}
			//新しい入力フォームのPjをそれぞれ分割し、それぞれRole_InfoにInsertする
			splitedPjNames := splitPjsInSameRole(rITTP.PjName)
			for _, splitedPjName := range splitedPjNames {
				err = w.insertRowToRoleInfoDB(splitedPjName, rITTP.RoleName)
				if err != nil {
					log.Println(err)
				}
			}
			if err != nil {
				log.Println(err)
			}
		}

	}

	return err
}

// 入力フォームのPjがデータベース登録しているもの一致するかどうか
func (w *WeddingInTypingPage) isRoleInfoExistInDBByInput(r RoleInfoInTypingPage) (valid bool, err error) {
	// fmt.Println(w.Date, w.Ampm, r.RoleName, r.PjName)
	valid = true
	splitedPjNames := splitPjsInSameRole(r.PjName)
	count, err := w.getRolesCount(r.RoleName)
	if err != nil {
		log.Println()
	}
	if count != len(splitedPjNames) {
		valid = false
		return valid, err
	}
	for _, splitedPjName := range splitedPjNames {
		ok, err := w.isRoleInfoExistInDB(r.RoleName, splitedPjName)
		if err != nil {
			log.Println(err)
		}
		if !ok {
			valid = false
			return valid, err
		}
	}
	return valid, err
}

// その役割の人数を数える
func (w *WeddingInTypingPage) getRolesCount(rolename string) (count int, err error) {
	cmd := `SELECT 
			count(*) 
		FROM
			role_info AS ri
				INNER JOIN
			weddings AS w ON w.date = ? and w.ampm = ?
				INNER JOIN
			pjs AS p ON ri.pj_id = p.id
				INNER JOIN
			roles AS r ON ri.role_id = r.id and r.name = ?;`
	Db.QueryRow(cmd, w.Date, w.Ampm, rolename).Scan(&count)
	return count, err
}

// データベースにPjが存在するかどうか
func (w *WeddingInTypingPage) isRoleInfoExistInDB(rolename string, pjname string) (valid bool, err error) {
	cmd := `SELECT
				case
				when count(*) > 0 then 'true'
				else 'false'
				end as 'IsExist'
			FROM
				role_info AS ri
					INNER JOIN
					weddings AS w ON w.date = ? and w.ampm = ?
					INNER JOIN
				pjs AS p ON ri.pj_id = p.id
					INNER JOIN
				roles AS r ON ri.role_id = r.id and r.name = ?
			WHERE
				p.name = ?;`
	// fmt.Println(w.Date, w.Ampm, r.RoleName, r.PjName)
	err = Db.QueryRow(cmd, w.Date, w.Ampm, rolename, pjname).Scan(
		&valid,
	)
	if err != nil {
		log.Println(err)
	}
	return valid, err
}

// 新しいRole_InfoをInsert
func (w *WeddingInTypingPage) insertRowToRoleInfoDB(pjname string, rolename string) (err error) {
	cmd := `insert into role_info(wedding_id, role_id, pj_id)
					SELECT
					w.id,r.id,p.id
				FROM
					roles AS r
						JOIN
					pjs AS p ON p.name = ?
						JOIN
					weddings AS w ON w.date = ? and w.ampm = ?
				WHERE
					r.name = ?;`
	_, err = Db.Exec(cmd, pjname, w.Date, w.Ampm, rolename)
	// fmt.Println(r.PjName, w.Date, w.Ampm, r.RoleName)
	// fmt.Printf("Insert:%s", r.PjName)
	if err != nil {
		log.Println(err)
	}
	return err
}

// Role_Infoから役割名で振られた役割の全ての情報を削除
func (w *WeddingInTypingPage) deleteRowToRoleInfoDB(rolename string) (err error) {
	cmd := `delete from role_info where (wedding_id,role_id) IN(
		SELECT 
		w.id,r.id
	FROM
		role_info as ri
		join
		roles as r on r.id = ri.role_id
			JOIN
		weddings AS w ON w.date = ? and w.ampm = ?
	WHERE
		r.name = ?
		group by w.id);`
	_, err = Db.Exec(cmd, w.Date, w.Ampm, rolename)
	if err != nil {
		log.Println(err)
	}
	return err
}

// 名前から、今までどの役割を何回やったかカウントする
func GetRoleCountFromPast(pjname string) (rCs []RoleCount, err error) {
	cmd := `SELECT 
			REGEXP_REPLACE(r.name, 'P$', '') AS counted_name,
			COUNT(*) AS count
		FROM
			pjs AS p
				JOIN
			role_info AS ri ON p.id = ri.pj_id
				JOIN
			roles AS r ON ri.role_id = r.id
		WHERE
			p.name = ?
		GROUP BY counted_name;`
	rows, err := Db.Query(cmd, pjname)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		rC := RoleCount{}
		rows.Scan(
			&rC.Name,
			&rC.Count)
		rCs = append(rCs, rC)
	}
	return rCs, err
}

/*
// 同じ文字列を削除し、新しい文字列を返す
func deleteSameRoleName(rIITPs []RoleInfoInTypingPage) (result []string, err error) {
	uniqueStrings := make(map[string]bool)
	for _, rIITP := range rIITPs {
		if !uniqueStrings[rIITP.RoleName] {
			result = append(result, rIITP.RoleName)
			uniqueStrings[rIITP.RoleName] = true
		}
	}
	return result, err
}

// 特定の文字列削除する
func removeString(slice []string, target string) []string {
	result := []string{}
	for _, str := range slice {
		if str != target {
			result = append(result, str)
		}
	}
	return result
}
*/
