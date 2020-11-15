package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // justificiation for blank import : initialisaion of the sqlite driver
)

// SetupModels connects to the sqlite database
func SetupModels(logMode bool, dbFileName string) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbFileName)

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.LogMode(logMode)

	db.AutoMigrate(&Struct{}, &Field{}, &Action{}, &Diagram{}, &Const{}, &EnumDB{})

	return db
}
