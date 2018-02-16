package main

import (
	"github.com/frederic-gendebien/customer/commands"
	"log"
)

func main() {
	err := commands.Handle()
	if err != nil {
		log.Fatal(err)
	}
}
