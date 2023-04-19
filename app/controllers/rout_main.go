package controllers

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"RMV0.5/app/config"
)

func top(w http.ResponseWriter, r *http.Request) {
	file := ReadXlsxFile()
	fmt.Println(file.Name())
	generateHTML(w, file.Name(), "layout", "top")

}

func index(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(config.Config.Xlsxpath)
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
	// pjs := models.GetPjs(dd)
	// fmt.Println(pjs)
	date := fmt.Sprintf("%s月%s日", mm, dd)
	generateHTML(w, date, "layout", "typingpage")
}
