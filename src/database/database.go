package database

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type database struct{
	db *gorm.DB
}

var data *database

func(database *database) CreateConnection(url string, tables ...interface{}) error{
	if database.db!=nil{
		return nil
	}
	dba, err := gorm.Open(sqlite.Open(url), &gorm.Config{});
	if err != nil{
		return errors.New("failed to connect to database")
	}
	database.db = dba
	database.db.AutoMigrate(tables...)
	return nil
}

func GetDatabase() *database{
	return data
}

func(database *database) Get(obj interface{}, conditions ...interface{}) (error){
	if result := database.db.Find(obj, conditions...); result.Error!=nil{
		return result.Error
	}
	return nil
}

func(database *database) Create(obj interface{}, conditions ...interface{}) (error){
	if result := database.db.Create(obj); result.Error!=nil{
		return result.Error
	}
	return nil
}

func(database *database) Update(obj interface{}, newVal interface{}) error {
	if result := database.db.Model(obj).Updates(newVal); result.Error!=nil{
		return result.Error
	}
	return nil
}