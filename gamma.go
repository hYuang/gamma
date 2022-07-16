package main

import (
	"fmt"
	"gamma/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s ! this is new language\n", user.Username)
	fmt.Printf("Feel free to type in commands \n")
	repl.Start(os.Stdin, os.Stdout)
}
