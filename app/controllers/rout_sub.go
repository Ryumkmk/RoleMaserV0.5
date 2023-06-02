package controllers

import (
	"encoding/json"
	"log"
	"net/http"

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
