package main

import (
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/webzak/secrets/crypto"
	"golang.org/x/term"
)

func generate() error {
	genType := flag.String("t", "", "generator type")
	size := flag.Int("s", 0, "size in bytes")
	outType := flag.String("f", "hex", "output type, can be comma separated")

	flag.Parse()
	var err error
	var r []byte
	switch *genType {
	case "rand":
		r, err = cryptoRand(*size)
	case "coin":
		r, err = coin(*size)
		fmt.Printf("\n")
	default:
		err = errors.New("unsupported generation type")
	}
	if err != nil {
		return err
	}
	formats := strings.Split(*outType, ",")
	for _, f := range formats {
		f = strings.Trim(f, " ")
		switch f {
		case "hex":
			fmt.Println("hex:", hex.EncodeToString(r))
		case "b64":
			fmt.Println("b64:", crypto.BytesToB64(r))
		case "bip39":
			v, err := crypto.BytesToBIP39(r)
			if err != nil {
				return err
			}
			fmt.Println("bip39:", v)
		default:
			return fmt.Errorf("format %s is not supported. Use one of hex, b64, bip39", f)
		}
	}
	return nil
}

func cryptoRand(n int) ([]byte, error) {
	if n <= 0 {
		return nil, errors.New("size must be > 0")
	}
	return crypto.RandomBytes(n)
}

func coin(n int) ([]byte, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	ret := make([]byte, n)
	b := make([]byte, 1)
	read := 0
	for read < n {
		pos := 7
		var b4 byte
		for pos >= 0 {
			_, err = os.Stdin.Read(b)
			if err != nil {
				return nil, err
			}
			c := string(b[0])
			if c == "1" {
				b4 += 1 << pos
			} else if c != "0" {
				continue
			}
			pos--
		}
		fmt.Printf("%x", b4)
		ret[read] = b4
		read++
	}
	return ret, nil
}
