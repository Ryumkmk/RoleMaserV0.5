package models

import (
	"log"
)

type Role struct {
	ID   int
	Name string
}

type RoleCount struct {
	Name      string `json:"name"`
	CountAll  int    `json:"countall"`
	Count3Mon int    `json:"count3mon"`
}
type Role_Info struct {
	ID         int
	Wedding_ID int
	Role_ID    int
	Pj_ID      int
}

type RoleInfoInTypingPage struct {
	RoleName string
	PjName   string
}

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

func (w *WeddingInTypingPage) UpdateRoleInfoDBByAmpm(rIITPs []RoleInfoInTypingPage) (err error) {

	for _, rITTP := range rIITPs {
		ok, err := w.isRoleInfoExistInDBByInput(rITTP)
		if err != nil {
			log.Println(err)
		}
		// fmt.Println(ok)
		if !ok {
			err = w.deleteRowToRoleInfoDB(rITTP.RoleName)
			if err != nil {
				log.Println(err)
			}
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

func (w *WeddingInTypingPage) isRoleInfoExistInDBByInput(r RoleInfoInTypingPage) (valid bool, err error) {

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

func (w *WeddingInTypingPage) getRolesCount(rolename string) (count int, err error) {
	cmd := `SELECT 
				count(*)
			FROM
				role_info AS ri
					INNER JOIN
				weddings AS w ON w.id = ri.wedding_id
					INNER JOIN
				pjs AS p ON ri.pj_id = p.id
					INNER JOIN
				roles AS r ON ri.role_id = r.id
			WHERE
				w.date = ? and r.name = ? ;`
	Db.QueryRow(cmd, w.Date, rolename).Scan(&count)
	return count, err
}

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
	err = Db.QueryRow(cmd, w.Date, w.Ampm, rolename, pjname).Scan(
		&valid,
	)
	if err != nil {
		log.Println(err)
	}
	return valid, err
}

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
	if err != nil {
		log.Println(err)
	}
	return err
}

func (w *WeddingInTypingPage) deleteRowToRoleInfoDB(rolename string) (err error) {
	cmd := `DELETE ri FROM role_info AS ri
				JOIN
			roles AS r ON r.id = ri.role_id
				JOIN
			weddings AS w ON ri.wedding_id = w.id 
			WHERE
			r.name = ?
			AND w.date = ?;`
	_, err = Db.Exec(cmd, rolename, w.Date)
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetRoleCountFromPast(pjname string) (rCs []RoleCount, err error) {
	cmd := `SELECT 
				REGEXP_REPLACE(r.name, 'P$', '') AS counted_name,
				COUNT(*) AS count_all,
				COUNT(CASE
					WHEN w.date > DATE_SUB(CURDATE(), INTERVAL 1 MONTH) THEN 1
				END) AS count_3months
			FROM
				pjs AS p
					JOIN
				role_info AS ri ON p.id = ri.pj_id
					JOIN
				roles AS r ON ri.role_id = r.id
					JOIN
				weddings AS w ON w.id = ri.wedding_id
			WHERE
				p.name = ?
			GROUP BY counted_name
			ORDER BY CASE
				WHEN r.name REGEXP 'P$' THEN r.id - 18
				ELSE r.id
			END;`
	rows, err := Db.Query(cmd, pjname)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		rC := RoleCount{}
		rows.Scan(
			&rC.Name,
			&rC.CountAll,
			&rC.Count3Mon)
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
