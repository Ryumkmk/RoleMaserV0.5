package models

import (
	"fmt"
	"log"

	"RMV0.5/app/config"
	"RMV0.5/app/controllers"
	"github.com/xuri/excelize/v2"
)

type Pjs []string

func GetPjs(day string) {
	f := controllers.ReadXlsxFile()
	xf, err := excelize.OpenFile(config.Config.Xlsxpath + "/" + f.Name())
	defer xf.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(xf)
	// return
}
