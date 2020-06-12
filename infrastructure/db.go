package infrastructure

import (
	"github.com/AwataKyosuke/go_api_server/util/config"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// DB データベース接続情報
type DB struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Connection *gorm.DB
}

// NewDB データベース情報を取得し新しく作成する
func NewDB() *DB {
	db := &DB{
		Host:     config.Config.Host,
		Username: config.Config.UserName,
		Password: config.Config.Password,
		DBName:   config.Config.DBName,
	}
	return db
}

// Connect コネクションを生成する
func Connect(db *DB) (*gorm.DB, error) {
	con, err := gorm.Open("mysql", db.Username+":"+db.Password+"@tcp("+db.Host+")/"+db.DBName+"?parseTime=true&&loc=Asia%2FTokyo&charset=utf8")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	db.Connection = con
	return db.Connection, nil
}
