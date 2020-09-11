package mysqlrepository

import (
	"github.com/AwataKyosuke/go_api_server/util/config"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// GetConnection データベースと接続する
func GetConnection() (*gorm.DB, error) {

	con, err := gorm.Open("mysql", config.Config.UserName+":"+config.Config.Password+"@tcp("+config.Config.Host+")/"+config.Config.DBName+"?parseTime=true&&loc=Asia%2FTokyo&charset=utf8")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return con, nil
}
