package models

import (
	"log"
)

type RestInCheckPage struct {
	RoleName string
	PjName   string
}

func (w *WeddingInTypingPage) MakeRest() (rICPs []RestInCheckPage, err error) {

	cmdAM := `SELECT
					r.name,
					p.name
				FROM
					(
						SELECT
							p.id,
							MIN(r.id) AS min_r_id
						FROM
							role_info AS ri
							JOIN weddings AS w ON w.id = ri.wedding_id
							JOIN roles AS r ON r.id = ri.role_id
							JOIN pjs AS p ON p.id = ri.pj_id
						WHERE
							w.date = ?
							AND r.name REGEXP 'P$' = FALSE
						GROUP BY
							p.id
					) AS sub
					JOIN pjs AS p ON p.id = sub.id
					JOIN roles AS r ON r.id = sub.min_r_id
				ORDER BY
					r.id;`
	rowsAM, err := Db.Query(cmdAM, w.Date)
	if err != nil {
		log.Println(err)
	}
	defer rowsAM.Close()

	for rowsAM.Next() {
		var rICP = RestInCheckPage{}
		rowsAM.Scan(
			&rICP.RoleName,
			&rICP.PjName,
		)
		rICPs = append(rICPs, rICP)
	}

	cmdPM := `SELECT 
					r.name, p.name
				FROM
					(SELECT 
						p.id, MIN(r.id) AS min_r_id
					FROM
						role_info AS ri
					JOIN weddings AS w ON w.id = ri.wedding_id
					JOIN roles AS r ON r.id = ri.role_id
					JOIN pjs AS p ON p.id = ri.pj_id
					WHERE
						w.date = ?
							AND r.name REGEXP 'P$' IS NOT FALSE
					GROUP BY p.id) AS sub
						JOIN
					pjs AS p ON p.id = sub.id
						JOIN
					roles AS r ON r.id = sub.min_r_id
				ORDER BY r.id;`
	rowsPM, err := Db.Query(cmdPM, w.Date)
	if err != nil {
		log.Println(err)
	}
	defer rowsPM.Close()

	for rowsPM.Next() {
		var rICP = RestInCheckPage{}
		rowsPM.Scan(
			&rICP.RoleName,
			&rICP.PjName,
		)
		rICPs = append(rICPs, rICP)
	}
	// fmt.Println(rICPs)
	return rICPs, err
}
