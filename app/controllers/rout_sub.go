package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"RMV0.5/app/models"
)

func makeresttypingpage(w http.ResponseWriter, r *http.Request) {
	var wITP models.WeddingInTypingPage
	wITP.Date = r.URL.Query().Get("date")
	rICPs, err := wITP.MakeRest()
	if err != nil {
		log.Println(err)
	}
	response := struct {
		RICPs []models.RestInCheckPage
	}{
		RICPs: rICPs,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func uploadPj(w http.ResponseWriter, r *http.Request) {
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
		if strings.Contains(n, "trainer") || strings.Contains(n, "trainee") {
			if strings.Contains(n, "trainer") {
				key := string(n[7:])
				tT := models.TrainerTrainee{
					Key:     key,
					Trainer: v[0],
					Trainee: r.PostFormValue("trainee" + key),
				}
				if len(tT.Trainer) > 0 && len(tT.Trainee) > 0 {
					tTs = append(tTs, tT)
				}
			}
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
			} else if n[len(n)-1] == 'P' {
				rIITPPM := models.RoleInfoInTypingPage{
					RoleName: n,
					PjName:   "",
				}
				rIITPsPM = append(rIITPsPM, rIITPPM)

			} else {
				rIITPAM := models.RoleInfoInTypingPage{
					RoleName: n,
					PjName:   "",
				}
				rIITPsAM = append(rIITPsAM, rIITPAM)
			}
		}
	}
	var dITP = models.DataInTypingPage{
		WITPAM:   wITPAM,
		WITPPM:   wITPPM,
		RIITPsAM: rIITPsAM,
		RIITPsPM: rIITPsPM,
		TTs:      tTs,
	}
	if err = dITP.UpdateRoleInfoDB(); err != nil {
		log.Println(err)
	}
	if err = models.UpdateTrainerPjDB(dITP.WITPAM.Date, dITP.TTs); err != nil {
		log.Println(err)
	}
}

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
