package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"RMV0.5/app/models"
)

// TopページのHtmlを生成
func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		// generateHTML(w, nil, "layout", "top")
	} else {
		generateHTML(w, nil, "layout", "top")
	}
}

// アクセス時のHtmlを生成
func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		http.Redirect(w, r, "/top", http.StatusFound)
	}
}

// 入力ページのHtmlを生成
func typingpage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	mm := r.PostFormValue("month")
	dd := r.PostFormValue("day")
	var pjs []models.Pj

	gotpjsName, gotpjsTime, gotpjsAmPm, amguest, pmguest := models.GetPjs(mm, dd)
	for i, v := range gotpjsName {
		var pj models.Pj
		pj.Names = v
		pj.Time = gotpjsTime[i]
		pj.AmPm = gotpjsAmPm[i]
		pj.Date = fmt.Sprintf("%s月%s日", mm, dd)
		pj.CheckAM = false
		pj.CheckPM = false
		pj.Amguest = amguest
		pj.Pmguest = pmguest
		pj.IsNewPj()
		pjs = append(pjs, pj)
	}
	generateHTML(w, pjs, "layout", "typingpage")
}

// 確認ページのHtmlを生成
func cheakPj(w http.ResponseWriter, r *http.Request) {
	var whatJob models.WhatJob
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	date := r.PostFormValue("date-form")
	for n, v := range r.Form {
		if strings.Contains(n, "trainer") || strings.Contains(n, "trainee") {
			models.WhosTrainee(n, v[0], &whatJob.Trainers)
		} else {
			var role models.Role
			role.PjName = v[0]
			role.RoleName = n
			whatJob.Roles = append(whatJob.Roles, role)
		}
	}
	// fmt.Println(whatJob.Trainers)

	dateDay := date[2:]
	dateMonth := date[:1]
	dateDaySt := ""
	dateMonthSt := ""

	for _, ch := range dateDay {
		if '0' <= ch && ch <= '9' {
			dateDaySt += string(ch)
		}
	}
	for _, ch := range dateMonth {
		if '0' <= ch && ch <= '9' {
			dateMonthSt += string(ch)
		}
	}
	gotpjsName, gotpjsTime, gotpjsAmPm, _, _ := models.GetPjs(dateMonthSt, dateDaySt)
	// fmt.Println(len(gotpjsName), len(gotpjsTime))
	for i, v := range gotpjsName {
		var pj models.Pj
		pj.Names = v
		pj.Time = gotpjsTime[i]
		pj.AmPm = gotpjsAmPm[i]
		pj.Date = date
		pj.CheckAM = false
		pj.CheckPM = false
		whatJob.Pjs = append(whatJob.Pjs, pj)
	}
	// fmt.Println(whatJob.Roles)
	// for _, v := range whatJob.Roles {
	// 	fmt.Println(v.PjName)
	// }

	models.IsInputPjs(gotpjsName, &whatJob)
	generateHTML(w, whatJob, "layout", "checkpage")
}

func attendancelist(w http.ResponseWriter, r *http.Request) {
	var List models.AttendanceList
	List.PjName = r.FormValue("pjName")
	List.AttendanceDaysList, List.AttendanceTimeList = models.GetAttendanceList(List.PjName)
	// fmt.Println(List)
	generateHTML(w, List, "layout", "attendancelist")
}
