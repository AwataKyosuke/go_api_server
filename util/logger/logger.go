package logger

import (
	"io"
	"log"
	"os"
)

// Setting ログ書き込みに関する設定を行う
func Setting(fileName string) {

	// ログファイルを開く
	logfile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file=logFile err=%s", err.Error())
	}

	// ログ出力と同時にファイルに書き込む
	multiLogFile := io.MultiWriter(os.Stdout, logfile)

	// ログ出力時の設定
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

// Fatal 致命的なエラーログの出力
func Fatal(err error) {
	log.SetPrefix("[FATAL] ")
	log.Printf("%+v\n", err)
}

// Error 重大なエラーログの出力
func Error(err error) {
	log.SetPrefix("[ERROR] ")
	log.Printf("%+v\n", err)
}

// Warn 注意が必要なエラーログの出力
func Warn(err error) {
	log.SetPrefix("[WARN] ")
	log.Printf("%+v\n", err)
}

// Info 一般的なログの出力
func Info(msg string) {
	log.SetPrefix("[INFO] ")
	log.Printf("%+v\n", msg)
}
