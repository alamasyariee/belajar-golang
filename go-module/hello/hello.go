package main

import "fmt"
import "log"

import "example.com/salam"

func main() {
	log.SetPrefix("salam: ")
	log.SetFlags(0)

	message, err := salam.Salam("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
