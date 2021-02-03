package mysql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(user, password, host, port, name string) *gorm.DB {

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
