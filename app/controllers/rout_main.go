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
	// fmt.Println(len(gotpjsName), len(gotpjsTime))
	for i, v := range gotpjsName {
		var pj models.Pj
		pj.Names = v
		pj.Time = gotpjsTime[i]
		pj.Date = fmt.Sprintf("%s月%s日", mm, dd)
		pj.Check = false
		pjs = append(pjs, pj)
	}
	// fmt.Println(pjs)
	generateHTML(w, pjs, "layout", "typingpage")
}

func cheakPj(w http.ResponseWriter, r *http.Request) {
	var whatJob models.WhatJob
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	date := r.PostFormValue("date-form")

	for n, v := range r.Form {
		var role models.Role
		role.PjName = v[0]
		role.RoleName = n
		whatJob.Roles = append(whatJob.Roles, role)
	}
	dateDay := date[2:]
	dateDaySt := ""
	for _, ch := range dateDay {
		if '0' <= ch && ch <= '9' {
			dateDaySt += string(ch)
		}
	}
	gotpjsName, gotpjsTime := models.GetPjs(dateDaySt)
	// fmt.Println(len(gotpjsName), len(gotpjsTime))
	for i, v := range gotpjsName {
		var pj models.Pj
		pj.Names = v
		pj.Time = gotpjsTime[i]
		pj.Date = date
		pj.Check = false
		whatJob.Pjs = append(whatJob.Pjs, pj)
	}
	whatJob = models.IsInputPjs(gotpjsName, &whatJob)
	// fmt.Println(whatJob.Pjs)
	generateHTML(w, whatJob, "layout", "checkpage")
}
