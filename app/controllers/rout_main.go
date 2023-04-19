package controllers

import (
	"fmt"
	"log"
	"net/http"

	"RMV0.5/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	file := models.ReadXlsxFile()
	fmt.Println(file.Name())
	generateHTML(w, file.Name(), "layout", "top")

}

func index(w http.ResponseWriter, r *http.Request) {

	f := models.ReadXlsxFile()
	if f == nil {
		generateHTML(w, nil, "layout", "upload")
	} else {
		generateHTML(w, f.Name(), "layout", "top")
	}

}

func typingpage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	mm := r.PostFormValue("month")
	dd := r.PostFormValue("day")
	var pjs []models.Pj

	gotpjsName, gotpjsTime := models.GetPjs(dd)
	for i, _ := range gotpjsName {
		var pj models.Pj
		pj.Names = gotpjsName[i]
		pj.Time = gotpjsTime[i]
		pj.Date = fmt.Sprintf("%s月%s日", mm, dd)
		pjs = append(pjs, pj)
	}
	// fmt.Println(pjs)
	generateHTML(w, pjs, "layout", "typingpage")
}
