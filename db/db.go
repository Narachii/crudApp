package db

import (
	"github.com/Narachii/crudApp/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Init() {
	dbms := "mysql"
	user := "root"
	protocol := "tcp(172.31.20.143:3306)"
	password := ""
	dbName := "test"

	connect := user + ":" + password + "@" + protocol + "/" + dbName
	database, err := gorm.Open(dbms, connect)
	if err != nil {
		panic(err)
	}

	db = database
	db.AutoMigrate(&entity.Book{})
}

func GetDB() *gorm.DB {
	return db
}

func Close()  {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
