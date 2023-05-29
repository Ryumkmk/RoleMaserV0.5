package controllers

import (
	"log"
	"net/http"

	"RMV0.5/app/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "signup")
		} else {
			http.Redirect(w, r, "/top", http.StatusFound)
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		au := models.Admin_User{
			Name:     r.PostFormValue("name"),
			Password: r.PostFormValue("password"),
		}
		if err := au.CreateUser(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/top", http.StatusFound)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "login")
	} else {
		http.Redirect(w, r, "/top", http.StatusFound)
	}

}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	au, err := models.GetUserByName(r.PostFormValue("name"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	if au.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := au.CreateSession()
		if err != nil {
			log.Println(err)
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/top", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}
