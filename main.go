package main

import (
	"fmt"
	"log"
	"github.com/rogercoll/p2p/pkg/server"

)

func main() {
	fmt.Println("hey broh")
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}