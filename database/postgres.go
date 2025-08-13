package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Postgre *sqlx.DB

func ConnectPostgres() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=new_password dbname=udemylifeserver sslmode=disable"
	Postgre, err = sqlx.Open("postgres", connStr)
	//	Postgre, err = sqlx.Open("postgres", "user=postgres dbname=udemylifeserver sslmode=disable")
	if err != nil {
		panic(err)
	}

	Postgre.SetMaxIdleConns(1)
	Postgre.SetMaxOpenConns(8)
}
