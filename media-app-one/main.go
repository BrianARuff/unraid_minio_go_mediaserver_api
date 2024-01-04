package main

import (
	"log"
	server "main/src/Server"
)

func main() {

	servErr := server.RunAPI()

	if servErr != nil {
		log.Fatalf("Server failed: %v", servErr)
	}
}
