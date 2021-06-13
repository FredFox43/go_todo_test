package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//os.O_RDWR(読み書き）|os.O_CREATE（作成）|os.O_APPEND（追記）
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)   //出力先を標準出力とlogfileに指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //フォーマット指定
	log.SetOutput(multiLogFile)
}
