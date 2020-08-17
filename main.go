package main

import (
	"donkey/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hee-Haw %s! This is the donkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands.\n")

	repl.Start(os.Stdin, os.Stdout)
}
