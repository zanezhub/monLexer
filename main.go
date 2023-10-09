package main

import (
	"fmt"
	"monLexer/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", user.Name)
	fmt.Printf("Type command\n")
	repl.Start(os.Stdin, os.Stdout)
}
