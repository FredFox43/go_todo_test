package config

import (
	"log"
	"todo_app_heroku/utils"

	"gopkg.in/go-ini/ini.v1"
)

type Configlist struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}

var Config Configlist //グローバルな変数宣言

func init() { //main関数より先に行われる。
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = Configlist{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		Static:    cfg.Section("web").Key("static").String(),
	}
}
