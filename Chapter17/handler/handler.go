package handler

import (
	"fmt"
	"github.com/i-coder-robot/book_final_code/Chapter17/conf"
	"github.com/i-coder-robot/book_final_code/Chapter17/config"
	"github.com/i-coder-robot/book_final_code/Chapter17/model"
	"github.com/i-coder-robot/book_final_code/Chapter17/repository"
	"github.com/i-coder-robot/book_final_code/Chapter17/service"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

var (
	DB            *gorm.DB
	AccountHandle AccountHandler
)

func InitViper() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
}

func InitDB() {
	fmt.Println("数据库 init")
	var err error
	conf := &conf.DBConf{
		Host:     viper.GetString("database.host"),
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DbName:   viper.GetString("database.name"),
	}

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DbName,
		true,
		"Local")

	DB, err = gorm.Open("mysql", config)
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	DB.SingularTable(true)
	fmt.Println("数据库 init 结束...")
}

func InitHandler() {
	AccountHandle = AccountHandler{
		Srv: &service.AccountService{
			Repo: &repository.AccountModelRepo{
				DB: model.DataBase{MyDB: DB},
			},
		}}
}
