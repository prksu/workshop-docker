package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// Default database configuration
var (
	Host     = "127.0.0.1"
	User     = "todos"
	Password = "secret"
	Name     = "todos"
)

// Database interface
type Database interface {
	Connect() *sql.DB
}

// Options implement database
type Options struct {
	DSN string
}

// NewDatabase create new database configuration
func NewDatabase() Database {
	database := new(Options)
	database.DSN = fmt.Sprintf("%s:%s@tcp(%s)/%s?autocommit=true&parseTime=true", User, Password, Host, Name)
	return database
}

// Connect open database connection.
func (o *Options) Connect() *sql.DB {
	database, err := sql.Open("mysql", o.DSN)
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatal(err)
	}

	database.SetMaxOpenConns(5)
	database.SetMaxIdleConns(5)
	database.SetConnMaxLifetime(time.Hour)
	return database
}
