package database

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct{
	db *gorm.DB
}

var data *Database

func(self *Database) CreateConnection(url string, tables ...interface{}) error{
	if self.db!=nil{
		return nil
	}
	dba, err := gorm.Open(sqlite.Open(url), &gorm.Config{});
	if err != nil{
		return errors.New("failed to connect to database")
	}
	self.db = dba
	self.db.AutoMigrate(tables...)
	return nil
}

func GetDatabase() *Database{
	if data==nil{
		data = &Database{}
	}
	return data
}

func(self *Database) Get(obj interface{}, conditions ...interface{}) (error){
	if result := self.db.Find(obj, conditions...); result.Error!=nil{
		return result.Error
	}
	return nil
}

func(self *Database) Create(obj interface{}, conditions ...interface{}) (error){
	if result := self.db.Create(obj); result.Error!=nil{
		return result.Error
	}
	return nil
}

func(self *Database) Update(obj interface{}, newVal interface{}) error {
	if result := self.db.Model(obj).Updates(newVal); result.Error!=nil{
		return result.Error
	}
	return nil
}