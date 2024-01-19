package connection

import (
	"database/sql"
	"sync"
    _ "github.com/go-sql-driver/mysql"
)

type dbConnection struct {
    db *sql.DB
    err error
}

var (
    connection *dbConnection
    once sync.Once
)

func GetDBConnection() (*sql.DB, error) {
    once.Do(func() {
        db, err := sql.Open("mysql", "root:rootpass@/jyatodos")
        connection = &dbConnection{db: db, err: err}
    })
    return connection.db, connection.err
}
