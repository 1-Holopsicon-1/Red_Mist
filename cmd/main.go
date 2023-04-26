package main

import (
	"RedMist/internal/app/db"
	"RedMist/internal/app/handler"
	"RedMist/server"
	"flag"
	"fmt"
	"log"
)

func main() {
	srv := new(server.Server)
	log.Println("Started program")
	defer log.Println("Ended Program")
	session := db.Connect()
	dbInstance, err := session.DB()
	defer dbInstance.Close()
	if err != nil {
		log.Fatalln(err, "error close db")
	}

	myHandler := handler.Handler{DB: session}
	migr := flag.Bool("migrate", false, fmt.Sprint("Migrating process"))
	start := flag.Bool("start", false, fmt.Sprint("Starting server"))
	flag.Parse()
	if *migr {
		log.Println("Migrating progress")
		db.Migrate(session)
	}
	if *start {
		log.Println("Openning server")
		if err := srv.Run(":4000", myHandler.InitRoutes()); err != nil {
			log.Fatalln(err)
		}
	}
}
