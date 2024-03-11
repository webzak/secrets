package secrets

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

// ReadLine reads line with provided hint
func ReadLine(hint string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(hint)
	s, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return s[:len(s)-1], nil
}

// ReadPassword reads password without echo
func ReadPassword(hint string) (string, error) {
	fmt.Print(hint)
	data, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return string(data), err
}
