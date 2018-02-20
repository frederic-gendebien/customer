package main

import (
	"log"

	"github.com/frederic-gendebien/customer/protocol"
)

func init() {
}

func main() {
	command, err := protocol.Parse()
	if err != nil {
		log.Fatal(err)
	}

	err = command.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
