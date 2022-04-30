package main

import (
	"database/sql"

	_ "github.com/dean2021/go-sqlite3"
)

func main() {
	for _, driver := range sql.Drivers() {
		println(driver)
	}
}
