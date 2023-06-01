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

var Db *sql.DB

var err error

const (
	shiftDateRowIndex = 0
	guestNumRowIndex  = 1
	ampmRowIndex      = 2
	shiftFileName = "pjシフト.xlsx"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (crypttext string) {
	crypttext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return crypttext
}
