package config

import (
	"log"

	"RMV0.5/app/xlsx"
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	Static    string
	Sheetname string
	Xlsxpath  string
}

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		Static:    cfg.Section("web").Key("static").String(),
		Sheetname: cfg.Section("web").Key("sheetname").String(),
	}
	Config.Xlsxpath = xlsx.GetXlsxPath()
	// fmt.Println(Config.Xlsxpath)
}
