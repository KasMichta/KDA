package main

import (
	"fmt"

	"KDA/pkg/greet"
)

func main() {
	message := greet.Hello("Kas")
	fmt.Println(message)
}
