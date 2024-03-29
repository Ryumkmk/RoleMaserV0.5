package controllers

import (
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
	dd := r.PostFormValue("day")
	date := models.ChangeDateFormatToDBFormat(yy, mm, dd)
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
	tTs, err := models.GetAllTrainersTraineesFromDB(date)
	if err != nil {
		log.Println(err)
	}

	var dITP = models.DataInTypingPage{
		PLITs:    pLTIs,
		WITPAM:   wITPAM,
		WITPPM:   wITPPM,
		RIITPsAM: rIITPsAM,
		RIITPsPM: rIITPsPM,
		TTs:      tTs,
	}
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

	var wITPAM models.WeddingInTypingPage
	var wITPPM models.WeddingInTypingPage
	var rIITPsAM []models.RoleInfoInTypingPage
	var rIITPsPM []models.RoleInfoInTypingPage
	var rICPs []models.RestInCheckPage

	date := r.PostFormValue("date")
	wITPAM.Date = date
	wITPPM.Date = date

	for n, v := range r.Form {
		if strings.Contains(n, "datetype2") {

			wITPAM.Date2 = v[0]
			wITPPM.Date2 = v[0]
		} else if strings.Contains(n, "date") {
			continue
		} else if strings.Contains(n, "ampmAM") {
			wITPAM.Ampm = v[0]
			rIITPsAM, err = wITPAM.GetRoleInfoByDateFromDB()
			if err != nil {
				log.Println(err)
			}
		} else if strings.Contains(n, "ampmPM") {
			wITPPM.Ampm = v[0]
			rIITPsPM, err = wITPPM.GetRoleInfoByDateFromDB()
			if err != nil {
				log.Println(err)
			}
		} else {
			var rICP = models.RestInCheckPage{
				RoleName: n,
				PjName:   v[0],
			}
			rICPs = append(rICPs, rICP)
		}

	}
	pLTIs, err := models.GetPjsByDateFromDB(date)
	if err != nil {
		log.Println()
	}
	tTs, err := models.GetAllTrainersTraineesFromDB(date)
	if err != nil {
		log.Println(err)
	}
	var dITP = models.DataInTypingPage{
		PLITs:    pLTIs,
		WITPAM:   wITPAM,
		WITPPM:   wITPPM,
		RIITPsAM: rIITPsAM,
		RIITPsPM: rIITPsPM,
		TTs:      tTs,
		RICPs:    rICPs,
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

func changeshift(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	password := r.PostFormValue("password")
	if password != "4649" {
		http.Redirect(w, r, "/top", http.StatusFound)
		return
	}
	yy := r.PostFormValue("year")
	mm := r.PostFormValue("month")
	dd := r.PostFormValue("day")
	var date string
	if len(yy) != 0 {
		date = models.ChangeDateFormatToDBFormat(yy, mm, dd)
	} else {
		date = r.URL.Query().Get("date")

	}
	wITP := models.GetWeddingsByDateFromDB(date)
	if len(wITP) == 0 {
		http.Redirect(w, r, "/top", http.StatusFound)
		return
	}
	pLITs, err := models.GetAllPjsShiftByDateFromDB(date, yy, mm, dd)
	sIICP := models.ShiftInfoInChangePage{
		WITP:  wITP[0],
		PLITs: pLITs,
	}
	if err != nil {
		log.Println(err)
	}

	generateHTML(w, sIICP, "layout", "changeshift")

}

func updatepj(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	pjname := r.PostFormValue("pjname")
	pj, err := models.GetPjInfoFromDB(pjname)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "top", http.StatusFound)
		return
	}

	for n, v := range r.Form {
		switch n {
		case "Level":
			pj.Level = v[0]
		case "Gatekeeper":
			pj.Gatekeeper = v[0]
		case "Toilet":
			pj.Toilet = v[0]
		case "Cloak":
			pj.Cloak = v[0]
		case "Silver":
			pj.Silver = v[0]
		case "Wash":
			pj.Wash = v[0]
		case "Ape":
			pj.Ape = v[0]
		case "Coffee":
			pj.Coffee = v[0]
		case "Champagne":
			pj.Champagne = v[0]
		case "Drinkcounter":
			pj.Drinkcounter = v[0]
		case "Leader":
			pj.Leader = v[0]
		}
	}

	if err := pj.UpdatePjDb(); err != nil {
		log.Println(err)
	}

	generateHTML(w, pj, "layout", "updatepj")
}
