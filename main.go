package main

import (
	"eco-journal/config"
	"eco-journal/route"
	"log"
)

func main() {
	config.ConnectDB()
	config.MigrateDB()
	r := route.SetupRouter()
	if err := r.Run(":8000"); err != nil {
		log.Fatal("Server Run Failed:", err)
	}
}
