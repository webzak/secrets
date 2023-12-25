package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}
	command := os.Args[1]
	os.Args = os.Args[1:]
	var err error
	switch command {
	case "prepare":
		err = prepare()
	case "read":
		err = read()
	case "help":
		help()
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		help()
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func help() {
	fmt.Println(`
Usage: secrets command [options]

Commands:

prepare - prepare secrets
    options:
        -storage <value>  - storage type, allowed values: enc-env
        -prefix  <value>  - prefix for secret name
        -uppercase        - uppercase name

read - read secret
    options:
        -name <value>  - secret name

help - show help`)
}
