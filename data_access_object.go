package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func DAO() (*sql.DB, error) {
	var user = os.Getenv("PG_USER")
	var pass = os.Getenv("PG_PASS")
	var conn_string = "user=" + user + " password=" + pass + " dbname=swipelong sslmode=disable"
	return sql.Open("postgres", conn_string)
}
