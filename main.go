package main

import (
	"github.com/AnkitDhawale/TodoListApp/app"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load ENV vars: %v", err)
	}
}

func main() {
	app.Start()
}
