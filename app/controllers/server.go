package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"RMV0.5/app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/top", top)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/", index)
	http.HandleFunc("/typingpage", typingpage)

	// fmt.Println("Stated Server")
	port := os.Getenv("PORT")
	return http.ListenAndServe(":"+port, nil)
}
