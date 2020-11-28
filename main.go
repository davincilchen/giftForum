package main

import (
	"giftForum/server"
	"log"
)

//main function
func main() {

	s := server.NewServer()
	err := s.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	s.Serve()

	s.Close()
}
