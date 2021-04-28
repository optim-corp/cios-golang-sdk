package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type (
	DBInfo struct {
		Host            string
		Port            string
		Driver          string
		User            string
		Pass            string
		DBName          string
		SslMode         string
		Etc             []string
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxIdleTime time.Duration
		ConnMaxLifetime time.Duration
	}
)

func ConnectDB(info DBInfo) (*sql.DB, error) {
	addNotEmpty := func(s, t, d string) string {
		if d != "" {
			s += fmt.Sprintf("%s%s ", t, d)
		}
		return s
	}
	template := addNotEmpty("", "user=", info.User)
	template = addNotEmpty(template, "password=", info.Pass)
	template = addNotEmpty(template, "dbname=", info.DBName)
	template = addNotEmpty(template, "sslmode=", info.SslMode)
	template = addNotEmpty(template, "host=", info.Host)
	template = addNotEmpty(template, "port=", info.Port)
	db, err := sql.Open(info.Driver, fmt.Sprintf(
		"%s %s",
		template, strings.Join(info.Etc, " ")))
	if err != nil {
		return nil, err
	}
	if info.ConnMaxLifetime != 0 {
		db.SetConnMaxLifetime(info.ConnMaxLifetime)
	}
	if info.ConnMaxIdleTime != 0 {
		db.SetConnMaxIdleTime(info.ConnMaxIdleTime)
	}
	if info.MaxOpenConns != 0 {
		db.SetMaxOpenConns(info.MaxOpenConns)
	}
	if info.MaxIdleConns != 0 {
		db.SetMaxIdleConns(info.MaxIdleConns)
	}
	return db, err
}

func MustConnectDB(info DBInfo) *sql.DB {
	db, err := ConnectDB(info)
	if err != nil {
		panic(err)
	}
	return db
}
