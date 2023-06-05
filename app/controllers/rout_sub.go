package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"RMV0.5/app/models"
)

func getRoleCount(w http.ResponseWriter, r *http.Request) {
	pjname := r.URL.Query().Get("pjname")
	rCs, err := models.GetRoleCountFromPast(pjname)
	if err != nil {
		log.Println(err)
	}
	response := struct {
		RoleCounts []models.RoleCount
	}{
		RoleCounts: rCs,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func deletepjshift(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	date := r.PostFormValue("date")
	pjname := r.PostFormValue("pjname")

	err = models.DeletePjShiftFromDB(date, pjname)
	if err != nil {
		log.Println(err)
	}

	redirectURL := fmt.Sprintf("/changeshift?date=%s", date)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func addpjshift(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	date := r.PostFormValue("date")
	redirectURL := fmt.Sprintf("/changeshift?date=%s", date)
	pjname := r.PostFormValue("pjname")
	shifttime := r.PostFormValue("shifttime")
	ampm := r.PostFormValue("ampm")

	re := regexp.MustCompile(`\d{1,2}:\d{2}-\d{1,2}:\d{2}`)
	fmt.Println(date, pjname, shifttime, ampm)
	if re.FindStringSubmatch(shifttime) == nil {
		http.Redirect(w, r, redirectURL, http.StatusFound)
		return
	}
	err = models.AddPjShiftFromDB(date, pjname, shifttime, ampm)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, redirectURL, http.StatusFound)
}
