package main

import (
	"flag"
	"fmt"

	"github.com/webzak/secrets"
)

func read() error {
	storage := flag.String("storage", "", "")
	name := flag.String("name", "", "")

	flag.Parse()
	var err error
	switch *storage {
	case "enc-env":
		err = readEncEnv(*name)
	default:
		err = fmt.Errorf("Storage '%s' is not supported", *storage)
	}
	return err
}

func readEncEnv(name string) error {
	fmt.Println()
	master, err := secrets.ReadPassword("storage password: ")
	if err != nil {
		return err
	}
	es, err := secrets.NewEncEnvStorage(master, "", false)
	if err != nil {
		return err
	}
	secret, err := es.Get(name)
	if err != nil {
		return err
	}
	fmt.Printf("secret: %s\n\n", secret)
	return nil
}
