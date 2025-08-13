package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Postgre *sqlx.DB

func ConnectPostgres() {
	var err error
	Postgre, err = sqlx.Open("postgres", "user=postgres dbname=udemylifeserver sslmode=disable")
	if err != nil {
		panic(err)
	}

	Postgre.SetMaxIdleConns(1)
	Postgre.SetMaxOpenConns(8)
}
