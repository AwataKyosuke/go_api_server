package config

import (
	"github.com/AwataKyosuke/go_api_server/util/logger"
	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
)

// ConfigList 設定ファイルから取得したデータを保持する構造体
type ConfigList struct {
	LogFile  string
	Host     string
	UserName string
	Password string
	DBName   string
}

// Config 設定リスト保持変数
var Config ConfigList

// コンストラクタ
func init() {
	// ファイル読み込み
	cfg, err := ini.Load("config.ini")
	if err != nil {
		logger.Fatal(errors.WithStack(err))
	}

	// 変数に設定
	Config = ConfigList{
		LogFile:  cfg.Section("go_api_server").Key("log_file").String(),
		Host:     cfg.Section("db_setting").Key("host").String(),
		UserName: cfg.Section("db_setting").Key("user_name").String(),
		Password: cfg.Section("db_setting").Key("password").String(),
		DBName:   cfg.Section("db_setting").Key("db_name").String(),
	}
}
