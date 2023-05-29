package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"RMV0.5/app/config"
)

// xlsxファイルをアップロード
func upload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	defer file.Close()

	f, err := os.OpenFile(config.Config.Xlsxpath+"/"+filepath.Base(header.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	io.Copy(f, file)
	http.Redirect(w, r, "/top", http.StatusFound)
}

// xlsxファイルを削除
func delete(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	file := r.PostFormValue("filename")
	err = os.Remove(config.Config.Xlsxpath + "/" + file)
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
