package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	file := ReadXlsxFile()
	fmt.Println(file.Name())
	generateHTML(w, file.Name(), "layout", "top")

}

func index(w http.ResponseWriter, r *http.Request) {

	f := ReadXlsxFile()
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
	// pjs := models.GetPjs(dd)
	// fmt.Println(pjs)
	date := fmt.Sprintf("%s月%s日", mm, dd)
	generateHTML(w, date, "layout", "typingpage")
}
