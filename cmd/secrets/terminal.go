package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

func readLine(hint string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(hint)
	s, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return s[:len(s)-1], nil
}

func readPassword(hint string) (string, error) {
	fmt.Print(hint)
	data, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return string(data), err
}
