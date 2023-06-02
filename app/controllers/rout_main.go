package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"RMV0.5/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		generateHTML(w, nil, "layout", "top")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		http.Redirect(w, r, "/top", http.StatusFound)
	}
}

func typingpage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	yy := r.PostFormValue("year")
	mm := r.PostFormValue("month")
	if len(mm) == 1 {
		mm = "0" + mm
	}
	dd := r.PostFormValue("day")
	if len(dd) == 1 {
		dd = "0" + dd
	}
	date := fmt.Sprintf("%s-%s-%s", yy, mm, dd)

	wITPs := models.GetWeddingsByDateFromDB(date)
	if len(wITPs) == 0 {
		http.Redirect(w, r, "/top", http.StatusFound)
		return
	}

	pLTIs, err := models.GetPjsByDateFromDB(date)
	if err != nil {
		log.Println(err)
	}
	var rIITPsAM []models.RoleInfoInTypingPage
	var rIITPsPM []models.RoleInfoInTypingPage
	var wITPAM models.WeddingInTypingPage
	var wITPPM models.WeddingInTypingPage

	for _, wITP := range wITPs {
		if wITP.Ampm == "AM" {
			rIITPsAM, err = wITP.GetRoleInfoByDateFromDB()
			wITPAM = wITP
			if err != nil {
				log.Println(err)
			}
		} else {
			rIITPsPM, err = wITP.GetRoleInfoByDateFromDB()
			wITPPM = wITP
			if err != nil {
				log.Println(err)
			}
		}
	}

	var dITP = models.DataInTypingPage{
		PLITs:    pLTIs,
		WITPAM:   wITPAM,
		WITPPM:   wITPPM,
		RIITPsAM: rIITPsAM,
		RIITPsPM: rIITPsPM,
	}
	// fmt.Println(dITP)
	if len(wITPs) == 2 {
		generateHTML(w, dITP, "layout", "doubleTypingPage")
	} else {
		if dITP.WITPAM.Ampm == "AM" {
			generateHTML(w, dITP, "layout", "amTypingPage")
		} else {
			generateHTML(w, dITP, "layout", "pmTypingPage")
		}
	}
}

func cheakPj(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	var rIITPsAM []models.RoleInfoInTypingPage
	var rIITPsPM []models.RoleInfoInTypingPage
	var wITPAM models.WeddingInTypingPage
	var wITPPM models.WeddingInTypingPage
	var tTs []models.TrainerTrainee

	for n, v := range r.Form {
		if strings.Contains(n, "trainer") {
			key := string(n[len(n)-1])
			tT := models.TrainerTrainee{
				Key:     key,
				Trainer: v[0],
				Trainee: r.PostFormValue("trainee" + key),
			}
			tTs = append(tTs, tT)
		} else if strings.Contains(n, "date-form") {
			wITPAM.Date = v[0]
			wITPPM.Date = v[0]
		} else if strings.Contains(n, "am-form") {
			wITPAM.Ampm = v[0]
		} else if strings.Contains(n, "pm-form") {
			wITPPM.Ampm = v[0]
		} else if strings.Contains(n, "datetype2") {
			wITPAM.Date2 = v[0]
			wITPPM.Date2 = v[0]
		} else {
			if len(v[0]) > 0 && n[len(n)-1] == 'P' {
				rIITPPM := models.RoleInfoInTypingPage{
					RoleName: n,
					PjName:   v[0],
				}
				rIITPsPM = append(rIITPsPM, rIITPPM)
			} else if len(v[0]) > 0 {
				rIITPAM := models.RoleInfoInTypingPage{
					RoleName: n,
					PjName:   v[0],
				}
				rIITPsAM = append(rIITPsAM, rIITPAM)
			}
		}
	}
	pLTIs, err := models.GetPjsByDateFromDB(wITPAM.Date)
	if err != nil {
		log.Println()
	}

	var dITP = models.DataInTypingPage{
		WITPAM:   wITPAM,
		WITPPM:   wITPPM,
		RIITPsAM: rIITPsAM,
		RIITPsPM: rIITPsPM,
	}

	err = dITP.UpdateRoleInfoDB()
	if err != nil {
		log.Println(err)
	}

	dITP.TTs = tTs
	dITP.PLITs = pLTIs
	dITP.RICPs, err = dITP.WITPAM.MakeRest()
	if err != nil {
		log.Println(err)
	}

	if len(dITP.WITPAM.Ampm) > 0 && len(dITP.WITPPM.Ampm) > 0 {
		generateHTML(w, dITP, "layout", "doublecheckpage")
	} else if len(dITP.WITPAM.Ampm) > 0 {
		generateHTML(w, dITP, "layout", "amcheckpage")
	} else {
		generateHTML(w, dITP, "layout", "pmcheckpage")
	}

}

func shiftlist(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	mm := r.PostFormValue("month")
	if len(mm) == 1 {
		mm = "0" + mm
	}
	name := r.PostFormValue("pjName")
	shifts, err := models.GetAllShiftByName(name, mm)
	if err != nil {
		log.Println(err)
	}
	generateHTML(w, shifts, "layout", "shiftlist")
}
