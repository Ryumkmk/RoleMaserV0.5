package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func top(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("./xlsx/")
	if err != nil {
		fmt.Println(err)
	}
	file := files[0]
	generateHTML(w, file.Name(), "layout", "top")
}

func index(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("./xlsx/")
	if err != nil {
		log.Println(err)
	}

	if len(files) == 0 {
		generateHTML(w, "ファイルがありせん", "layout", "upload")
	} else {
		file := files[0]
		generateHTML(w, file.Name(), "layout", "top")
	}
}

func typingpage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	mm := r.PostFormValue("month")
	dd := r.PostFormValue("day")
	date := fmt.Sprintf("%s月%s日", mm, dd)
	fmt.Println(date)
	generateHTML(w, date, "layout", "typingpage")
}
