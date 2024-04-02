package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/connor-ashton-dev/chad/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Yo %s! Welcome to Chadlang. The programming language for chads.\n", user.Username)
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
