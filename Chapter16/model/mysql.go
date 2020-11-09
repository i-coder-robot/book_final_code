package model

import (
	"fmt"
	"github.com/i-coder-robot/book_final_code/Chapter16/MyLog"
	"github.com/jinzhu/gorm"
	// MySQL driver.
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type DataBase struct {
	MyDB *gorm.DB
}

var DB *DataBase

func (db *DataBase) Init() {
	DB = &DataBase{
		MyDB: GetMySqlDB(),
	}
}

func (db *DataBase) Close() {
	DB.MyDB.Close()
}

func InitSelfDB() *gorm.DB {
	db := openDB(viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.addr"),
		viper.GetString("database.name"))
	return db
}

func GetMySqlDB() *gorm.DB {
	return InitSelfDB()
}

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		MyLog.Log.Errorf("Database connection failed. Database name: %s,Eroor:%s", name, err.Error())
	}

	setupDB(db)
	db.SingularTable(true)
	return db
}

func setupDB(db *gorm.DB) {
	// 用于设置闲置的连接数.
	db.DB().SetMaxIdleConns(5)
}
