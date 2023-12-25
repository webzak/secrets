package main

import (
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/webzak/secrets/crypto"
	"golang.org/x/term"
)

func generate() error {
	genType := flag.String("t", "rand", "generator type")
	size := flag.Int("s", 0, "size in bytes")
	format := flag.String("f", "hex", "output type, can be comma separated")

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
	var res string
	switch *format {
	case "hex":
		res = hex.EncodeToString(r)
	case "b64":
		res = crypto.BytesToB64(r)
	case "bip39":
		res, err = crypto.BytesToBIP39(r)
	default:
		err = fmt.Errorf("format %s is not supported. Use one of hex, b64, bip39", *format)
	}

	if err != nil {
		return err
	}
	fmt.Print(res)
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
