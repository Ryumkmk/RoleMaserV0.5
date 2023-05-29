package controllers

import (
	"log"
	"net/http"

	"RMV0.5/app/models"
)

//サインアップ
func signup(w http.ResponseWriter, r *http.Request) {

	//GETでアクセスした時
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			//未ログイン状態ならサインアップページの作成
			generateHTML(w, nil, "layout", "signup")
		} else {
			//ログイン状態ならTopへ飛ぶ
			http.Redirect(w, r, "/top", http.StatusFound)
		}
		
		//POSTでアクセスした時
	} else if r.Method == "POST" {
		//POSTした値を取得
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		au := models.Admin_User{
			Name:     r.PostFormValue("name"),
			Password: r.PostFormValue("password"),
		}
		//ユーザー名とパスワードからユーザーを作成し、データベースに登録
		if err := au.CreateUser(); err != nil {
			log.Println(err)
		}
		//ユーザー作成後、Topへ飛ぶ
		http.Redirect(w, r, "/top", http.StatusFound)
	}
}

//ログイン
func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		//未ログイン状態ならloginページを作成
		generateHTML(w, nil, "layout", "login")
	} else {
		//ログイン状態ならTopへ飛ぶ
		http.Redirect(w, r, "/top", http.StatusFound)
	}

}

//ユーザー認証
func authenticate(w http.ResponseWriter, r *http.Request) {
	//POSTした値をParse
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	//入力したユーザー名で、データベースから該当ユーザーを取得
	au, err := models.GetUserByName(r.PostFormValue("name"))
	if err != nil {
		log.Println(err)
		//ユーザーがデータベースに存在しないならもう一度ログインページへ飛ぶ
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	
	if au.Password == models.Encrypt(r.PostFormValue("password")) {

		//ユーザーが存在し＆パスワードが一致するなら、ログイン状態セッション情報を登録
		session, err := au.CreateSession()
		if err != nil {
			log.Println(err)
		}
		//セッション情報をクッキーに登録
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		//ログインしたならTopへ飛ぶ
		http.Redirect(w, r, "/top", http.StatusFound)
	} else {
		//パスワードが違うならもう一度ログインページへ飛ぶ
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

//ログアウト
func logout(w http.ResponseWriter, r *http.Request) {
	//クッキーのセッション情報を取得
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		//セッション情報をデータベースから削除
		session.DeleteSessionByUUID()
	}
	//ログインページへ飛ぶ
	http.Redirect(w, r, "/login", http.StatusFound)
}
