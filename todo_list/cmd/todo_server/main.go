package main

import (
	"internal/todo_server"
	"log"
)

func main() {
	srv := new(todo_server.Server)
	if error := srv.Run("8080"); error != nil {
		log.Fatal(error)
	}
}
