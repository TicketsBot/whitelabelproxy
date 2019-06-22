package database

import (
	"fmt"
	"github.com/TicketsBot/whitelabelproxy/config"
	"github.com/jinzhu/gorm"
)

type Database struct {
	*gorm.DB
}

var Db Database

func Connect() {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.Database.Username,
		config.Conf.Database.Password,
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.Database,
	)

	db, err := gorm.Open("mysql", uri); if err != nil {
		panic(err)
	}

	db.DB().SetMaxOpenConns(config.Conf.Database.Pool.MaxConnections)
	db.DB().SetMaxIdleConns(config.Conf.Database.Pool.MaxIdle)

	db.Set("gorm:table_options", "charset=utf8mb4")

	Db = Database { db }
}

func Setup() {
	Db.AutoMigrate()
}

func IsConnected(ch chan bool) {
	ch <- Db.DB.DB().Ping() == nil
}
