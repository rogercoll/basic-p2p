package main

import (
	"fmt"
	"log"
	"github.com/rogercoll/p2p/pkg/client"

)

func main() {
	fmt.Println("hey broh")
	err := client.Run()
	if err != nil {
		log.Fatal(err)
	}
}