package main

import (
	"bufio"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/tyler-smith/go-bip39"
	"github.com/webzak/secrets/crypto"
)

func convert() error {
	fromType := flag.String("from", "", "from type")
	toType := flag.String("to", "", "to type")
	flag.Parse()

	// read from stdin
	reader := bufio.NewReader(os.Stdin)
	inbuf := make([]byte, 8192)
	n, err := reader.Read(inbuf)
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("empty input")
	}
	ins := string(inbuf[:n])

	var bin []byte

	switch *fromType {
	case "hex":
		bin, err = hex.DecodeString(ins)
	case "b64":
		bin, err = crypto.B64ToBytes(ins)
	case "bip39":
		bin, err = bip39.EntropyFromMnemonic(ins)
	default:
		err = fmt.Errorf("format %s is not supported. Use one of hex, b64, bip39", *toType)
	}
	if err != nil {
		return err
	}
	var res string
	switch *toType {
	case "hex":
		res = hex.EncodeToString(bin)
	case "b64":
		res = crypto.BytesToB64(bin)
	case "bip39":
		res, err = crypto.BytesToBIP39(bin)
		if err != nil {
			return err
		}
	default:
		err = fmt.Errorf("format %s is not supported. Use one of hex, b64, bip39", *toType)
	}
	if err != nil {
		return err
	}
	fmt.Print(res)
	return nil
}
