package infrastructure

import "github.com/jinzhu/gorm"

// DB データベース接続情報
type DB struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Connection *gorm.DB
}

// NewDB TODO COnfigファイル等に設定値を外出し
func NewDB() *DB {
	config := &DB{
		Host:     "mysql",
		Username: "root",
		Password: "password",
		DBName:   "development",
	}
	return config
}

// Connect コネクションを生成する
func Connect(db *DB) *gorm.DB {
	con, err := gorm.Open("mysql", db.Username+":"+db.Password+"@tcp("+db.Host+")/"+db.DBName+"?parseTime=true&&loc=Asia%2FTokyo&charset=utf8")
	if err != nil {
		// TODO エラー処理
	}
	db.Connection = con
	return db.Connection
}
