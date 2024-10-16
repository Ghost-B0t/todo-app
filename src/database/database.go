package database

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func CreateConnection(url string, tables ...interface{}) error{
	if db!=nil{
		return nil
	}
	database, err := gorm.Open(sqlite.Open(url), &gorm.Config{});
	if err != nil{
		return errors.New("failed to connect to database")
	}
	db = database
	db.AutoMigrate(tables...)
	return nil
}

func Get(obj interface{}, conditions ...interface{}) (error){
	if result := db.Find(obj, conditions...); result.Error!=nil{
		return result.Error
	}
	return nil
}

func Create(obj interface{}, conditions ...interface{}) (error){
	if result := db.Create(obj); result.Error!=nil{
		return result.Error
	}
	return nil
}