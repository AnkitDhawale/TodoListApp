package main

import (
	"github.com/AnkitDhawale/TodoListApp/app"
)

/*
init() func is not required since we are running it via docker compose and
have exported env already, so os.Getenv("KEY") is sufficient.
func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load ENV vars: %v", err)
	}
}
*/

func main() {
	app.Start()
}
