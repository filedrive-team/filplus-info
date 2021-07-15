package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func initDBForTest(t *testing.T) {
	var err error
	//db, err = gorm.Open(postgres.Open("postgresql://postgres:postgres@localhost:5432/datacap?sslmode=disable"))
	db, err = gorm.Open(mysql.Open("filplus:filplus@tcp(192.168.19.23:3306)/filplus?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		t.Fatalf("连接数据库失败 %v", err)
	}
}

func TestInsertNotary(t *testing.T) {
	initDBForTest(t)
	defer CloseDB()

	t.Fatalf("insert failed %v", InitNotaryList())
}

func TestDeleteNotaryById(t *testing.T) {
	initDBForTest(t)
	defer CloseDB()

	err := DeleteNotaryById(1)
	if err != nil {
		t.Fatalf("delete failed %v", err)
	}
}
