package config

import (
	"log"

	"RMV0.5/app/utils"
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	Static    string
	LogFile   string
	SQLDriver string
	DbName    string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {

	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		Static:    cfg.Section("web").Key("static").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
	}
}
