package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
)

func GetConnection() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     viper.GetString("database.addr"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.pass"),
		Database: viper.GetString("database.database"),
	})

	return db
}
