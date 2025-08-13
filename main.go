package main

import (
	"log"
	"runtime"
	"server-application/api"
	"server-application/database"
	"server-application/server"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	database.ConnectPostgres()
	s := server.NewRouter()
	api.Start(s.Group("/api"))
	s.Static("/", "./assets/")
	log.Println("Started on :8080")
	if err := s.ListenAndServe("127.0.0.1:8080"); err != nil {
		panic(err)
	}
}
