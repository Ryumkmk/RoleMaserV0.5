package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"RMV0.5/app/config"
	_ "github.com/go-sql-driver/mysql"

	"github.com/google/uuid"
)

// データベース型
var Db *sql.DB

// エラー型
var err error

// 定数
const (
	shiftDateRowIndex = 0
	guestNumRowIndex  = 1
	ampmRowIndex      = 2
	// sheetName         = "6月シフト"
	shiftFileName = "pjシフト.xlsx"
)

// アプリ起動後に先ずデータベースを読み込む
func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
}

// ユニークなUUIDを作る
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// パスワードをハッシュ値で変換する
func Encrypt(plaintext string) (crypttext string) {
	crypttext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return crypttext
}
