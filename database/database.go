
package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (db *gorm.DB) {
	dsn:= "root:@/shoppet?charset=utf8&parseTime=True&loc=Local"
	// dsn := "root:admin@/project?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error when connect to db ", err)
		return
	}

	if err != nil {
		log.Fatal("error when auto migrate table ", err)
	}
	return db
}
