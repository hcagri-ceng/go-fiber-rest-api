package database

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("myDatabase.db"))
	if err != nil {
		panic("failed to connect database")
	}
	DB = db // global değişkene atadık.

	fmt.Println("CONNECTED TO DATABASE")
}
