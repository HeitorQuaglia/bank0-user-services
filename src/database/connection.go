package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type DbConnection struct {
	*sql.DB
}

var (
	dbInstance *DbConnection
	dbOnce     sync.Once
)

func newDBConnection() *DbConnection {
	err := godotenv.Load()

	if err != nil {
		_ = fmt.Errorf("error loading .env file: %w", err)
		panic(err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		_ = fmt.Errorf("error connecting to database: %w", err)
		panic(err)
	}

	return &DbConnection{db}
}

func GetDBConnection() *DbConnection {

	dbOnce.Do(func() {
		dbInstance = newDBConnection()
	})

	return dbInstance
}
