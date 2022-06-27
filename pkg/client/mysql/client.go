package mysql

import (
	"go-clean-news-api/internal/domain/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewClient(connString string) (*gorm.DB, error){
	Instance, dbError := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
	}
	Migrate(Instance)
	log.Println("Connected to Database!")
	return Instance, dbError
}

func Migrate(Instance *gorm.DB) {
	Instance.AutoMigrate(&entity.User{})
	log.Println("Database Migration Completed!")
}