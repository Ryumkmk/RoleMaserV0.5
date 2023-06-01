package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"RMV0.5/app/models"
)

// TopページのHtmlを生成
func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		//未ログイン情報ならloginページへ飛ぶ
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		//ログイン状態ならTopを作成
		generateHTML(w, nil, "layout", "top")
	}
}

// 初期アクセスページ
func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		//未ログイン情報ならloginページへ飛ぶ
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		//ログイン状態ならTopページへ飛ぶ
		http.Redirect(w, r, "/top", http.StatusFound)
	}
}

/*
func allpjs(w http.ResponseWriter, r *http.Request) {
	y, _ := models.GetAllPjsByDB()
	generateHTML(w, y, "layout", "pjs")
}
*/

// 入力ページのHtmlを生成
func typingpage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	//入力した日付を2023-03-12のフォーマットに変換し、dateに格納
	yy := r.PostFormValue("year")
	mm := r.PostFormValue("month")
	if len(mm) == 1 {
		mm = "0" + mm
	}
	dd := r.PostFormValue("day")
	if len(dd) == 1 {
		dd = "0" + dd
	}
	date := fmt.Sprintf("%s-%s-%s", yy, mm, dd)
	// fmt.Println(date)

	//日付から婚礼情報を取得
	wITPs := models.GetWeddingsByDateFromDB(date)
	if len(wITPs) == 0 {
		http.Redirect(w, r, "/top", http.StatusFound)
		return
	}

	//日付から出勤PJ情報を取得
	pLTIs, err := models.GetPjsByDateFromDB(date)
	if err != nil {
		log.Println(err)
	}
	//AMPMのそれぞれの役割を取得
	var rIITPsAM []models.RoleInfoInTypingPage
	var rIITPsPM []models.RoleInfoInTypingPage
	var wITPAM models.WeddingInTypingPage
	var wITPPM models.WeddingInTypingPage

	for _, wITP := range wITPs {
		if wITP.Ampm == "AM" {
			rIITPsAM, err = wITP.GetRoleInfoByDateFromDB()
			// fmt.Println(rIITPsAM)
			wITPAM = wITP
			if err != nil {
				log.Println(err)
			}
		} else {
			rIITPsPM, err = wITP.GetRoleInfoByDateFromDB()
			// fmt.Println(rIITPsPM)
			wITPPM = wITP
			if err != nil {
				log.Println(err)
			}
		}
	}
	//全ての情報をdITPに格納
	var dITP = models.DataInTypingPage{
		PLITs:    pLTIs,
		WITPAM:   wITPAM,
		WITPPM:   wITPPM,
		RIITPsAM: rIITPsAM,
		RIITPsPM: rIITPsPM,
	}

	//ダブル、つまり婚礼情報を二つ持っていたらダブルの入力ページを生成
	if len(wITPs) == 2 {
		generateHTML(w, dITP, "layout", "doubleTypingPage")
	} else {
		//AMならAM入力ページを生成
		if dITP.WITPAM.Ampm == "AM" {
			generateHTML(w, dITP, "layout", "amTypingPage")
		} else {
			//PM or 試食会ならPM入力ページを生成
			generateHTML(w, dITP, "layout", "pmTypingPage")
		}
	}
}

// 確認ページのHtmlを生成
func cheakPj(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	var rIITPsAM []models.RoleInfoInTypingPage
	var rIITPsPM []models.RoleInfoInTypingPage
	var wITPAM models.WeddingInTypingPage
	var wITPPM models.WeddingInTypingPage
	var tTs []models.TrainerTrainee

	//入力フォームの値を取得
	for n, v := range r.Form {
		if strings.Contains(n, "trainer") {
			//トレイナーなら、セットとするトレイニーを組み合わせる
			key := string(n[len(n)-1])
			tT := models.TrainerTrainee{
				Key:     key,
				Trainer: v[0],
				Trainee: r.PostFormValue("trainee" + key),
			}
			tTs = append(tTs, tT)
		} else if strings.Contains(n, "date-form") {
			//日付情報を格納
			wITPAM.Date = v[0]
			wITPPM.Date = v[0]
		} else if strings.Contains(n, "am-form") {
			//AMならWedding情報をAMを追加
			wITPAM.Ampm = v[0]
		} else if strings.Contains(n, "pm-form") {
			//PMならWedding情報をPMを追加
			wITPPM.Ampm = v[0]
		} else {
			if len(v[0]) > 0 && n[len(n)-1] == 'P' {
				//PMの役割なら、役割名と入力フォームの値を追加
				rIITPPM := models.RoleInfoInTypingPage{
					RoleName: n,
					PjName:   v[0],
				}
				rIITPsPM = append(rIITPsPM, rIITPPM)
			} else if len(v[0]) > 0 {
				//AMの役割なら、役割名と入力フォームの値を追加
				rIITPAM := models.RoleInfoInTypingPage{
					RoleName: n,
					PjName:   v[0],
				}
				rIITPsAM = append(rIITPsAM, rIITPAM)
			}
		}
	}
	pLTIs, err := models.GetPjsByDateFromDB(wITPAM.Date)
	if err != nil {
		log.Println()
	}
	wITPAM.Date2 = r.PostFormValue("datetype2")
	wITPPM.Date2 = r.PostFormValue("datetype2")
	//引き渡し用のデータに登録
	var dITP = models.DataInTypingPage{
		PLITs:    pLTIs,
		WITPAM:   wITPAM,
		WITPPM:   wITPPM,
		RIITPsAM: rIITPsAM,
		RIITPsPM: rIITPsPM,
		TTs:      tTs,
	}
	//データベース更新
	err = dITP.UpdateRoleInfoDB()
	if err != nil {
		log.Println(err)
	}
	//ダブルならダブル確認ページを生成
	if len(dITP.WITPAM.Ampm) > 0 && len(dITP.WITPPM.Ampm) > 0 {
		generateHTML(w, dITP, "layout", "doublecheckpage")
	} else if len(dITP.WITPAM.Ampm) > 0 {
		//AMならAM確認ページを生成
		generateHTML(w, dITP, "layout", "amcheckpage")
	} else {
		//PM or 試食会AMならPM確認ページを生成
		generateHTML(w, dITP, "layout", "pmcheckpage")
	}
}

func shiftlist(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	mm := r.PostFormValue("month")
	if len(mm) == 1 {
		mm = "0" + mm
	}
	name := r.PostFormValue("pjName")
	shifts, err := models.GetAllShiftByName(name, mm)
	if err != nil {
		log.Println(err)
	}
	generateHTML(w, shifts, "layout", "shiftlist")
}

// 役割のカウントを取得
func getRoleCount(w http.ResponseWriter, r *http.Request) {
	pjname := r.URL.Query().Get("pjname")
	rCs, err := models.GetRoleCountFromPast(pjname)
	if err != nil {
		log.Println(err)
	}
	response := struct {
		RoleCounts []models.RoleCount
	}{
		RoleCounts: rCs,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
