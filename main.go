package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is Interpreter Shell by Oleksandr Babilia!\n", user.Username)
	fmt.Printf("Feel Free to Type in Command\n")
	repl.Strat(os.Stdin, os.Stdout)
}
