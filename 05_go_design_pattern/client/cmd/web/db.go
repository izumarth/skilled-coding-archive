package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	maxOpenDbConn = 25
	maxIdelDbConn = 25
	maxDBLifetime = 5 * time.Minute
)

func initMySQLDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// test our databse
	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdelDbConn)
	db.SetConnMaxLifetime(maxDBLifetime)

	return db, nil
}
