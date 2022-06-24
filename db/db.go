package db

import (
	"database/sql"	
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaBancoDados() *sql.DB {
				
	db, err := sql.Open("mysql", "usuario:senha@tcp(endpoint AWS:3306)/lojaGO")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}