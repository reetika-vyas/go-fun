package util

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

const (
	DB_TYPE = "mysql"
)

func CreateDbConnection(env, path string) (db *gorm.DB, err error) {
	log.WithFields(log.Fields{"Path": path, "Env": env}).Info("Initing DB")
	dbConf, _ := goose.NewDBConf(path, env, DB_TYPE)
	if db, err = gorm.Open(dbConf.PgSchema, dbConf.Driver.OpenStr); err == nil {
		/** Print SQL */
		//db.LogMode(true)

		//TODO:From Config
		db.DB().SetMaxIdleConns(5)
		db.DB().SetMaxOpenConns(20)
	}
	return
}
