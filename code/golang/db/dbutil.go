package dbutil

import (
	"database/sql"
	"fmt"
	"log"
)

// Conn ...
type Conn struct {
	db *sql.DB
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// New (dbHost, dbUser, dbPasswd, dbName string) *Conn
func New(dbHost, dbUser, dbPasswd, dbName string) *Conn {
	log.Println("hostname: ", dbHost)
	log.Println("username: ", dbUser)
	log.Println("password: ", dbPasswd)
	log.Println("database: ", dbName)

	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbUser, dbPasswd, dbName)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	return &Conn{db}
}

// Rows ...
func (conn *Conn) Rows(table string, columns string) *sql.Rows {
	rows, err := conn.db.Query(fmt.Sprintf("SELECT %s FROM %s", columns, table))
	checkErr(err)
	return rows
}

// Close ...
func (conn *Conn) Close() {
	conn.db.Close()
}
