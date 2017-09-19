package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"	
)


type DataOperations interface {
	RegisterUser(Username, Password, Mail string) bool
	GetAllUsers() []*User
	Login(username, password string) (bool, *User)
	CheckUser(username string) bool
}

type DB struct {
	*sql.DB
}

func InitDB(sourcename, sorceInfo string) *DB {
	db, err:= sql.Open(sourcename, sorceInfo)
	if err!=nil {
		panic(err)
	
	}
	err = db.Ping()

	if err!=nil {
		panic(err)
	}

	return &DB{db}
}