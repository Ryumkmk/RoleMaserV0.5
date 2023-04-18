package controllers

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
)

func top(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("./xlsx/")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".xlsx") {
			generateHTML(w, file.Name(), "layout", "top")
			break
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("./xlsx/")
	if err != nil {
		log.Println(err)
	}
	var xlsxfile fs.DirEntry
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".xlsx") {
			xlsxfile = file
			break
		}
	}
	if xlsxfile == nil {
		generateHTML(w, nil, "layout", "upload")
	} else {
		generateHTML(w, xlsxfile.Name(), "layout", "top")
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
	generateHTML(w, date, "layout", "typingpage")
}
