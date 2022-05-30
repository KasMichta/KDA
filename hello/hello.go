package main

import (
	"fmt"

	"github.com/KasMichta/KDA/greet"
)

func main() {
	message := greet.Hello("Kas")
	fmt.Println(message)
}
