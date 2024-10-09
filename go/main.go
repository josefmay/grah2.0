package main

import (
	"log"
)


func main() {
	http_server := NewHttpServer(":42069")
	err := http_server.Run()
	if err != nil{
		log.Fatal("Server failed to launch.")
	}
}