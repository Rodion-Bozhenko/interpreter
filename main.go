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

	printfln("Hello %s! This is the Ligma programming language\n", user.Username)
	printfln("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}

func printfln(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
