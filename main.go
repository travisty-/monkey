package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/travisty-/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in a command below:\n")
	repl.Start(os.Stdin, os.Stdout)
}
