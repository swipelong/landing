package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func DAO() (*sql.DB, error) {
	return sql.Open("postgres", "user=postgres dbname=swipelong sslmode=disable")
}
