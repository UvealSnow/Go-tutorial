package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("hello: ")
	log.SetFlags(0)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a name or comma separated list of names: ")
	ioIn, _ := reader.ReadString('\n')
	names, err := greetings.ParseNames(ioIn)

	if err != nil {
		log.Fatal(err)
	}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, value := range messages {
		fmt.Println(value)
	}
}
