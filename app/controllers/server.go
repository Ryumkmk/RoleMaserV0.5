package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"RMV0.5/app/config"
	"RMV0.5/app/models"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/top", top)
	http.HandleFunc("/", index)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/typingpage", typingpage)
	http.HandleFunc("/shiftlist", shiftlist)
	http.HandleFunc("/checkPj", cheakPj)
	http.HandleFunc("/getRoleCount", getRoleCount)
	http.HandleFunc("/changeshift", changeshift)
	http.HandleFunc("/deletepjshift", deletepjshift)
	http.HandleFunc("/addpjshift", addpjshift)
	http.HandleFunc("/updatepj", updatepj)
	http.HandleFunc("/uploadPj", uploadPj)
	http.HandleFunc("/makeresttypingpage", makeresttypingpage)
	http.HandleFunc("/ispjinputed", ispjinputed)
	http.HandleFunc("/ispjinputeddouble", ispjinputeddouble)

	fmt.Println("Stated Server")
	port := os.Getenv("PORT")
	return http.ListenAndServe(":"+port, nil)
	// return http.ListenAndServe(":"+config.Config.Port, nil)
}
