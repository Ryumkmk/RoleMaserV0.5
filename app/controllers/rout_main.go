package controllers

import (
	"fmt"
	"net/http"
	"os"
)

func top(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("./xlsx/")
	if err != nil {
		fmt.Println(err)
	}

	if len(files) == 0 {
		generateHTML(w,"ファイルがありせん", "layout", "top")
	} else {
		file := files[0]
		generateHTML(w, file.Name(), "layout", "top")
	}
}
