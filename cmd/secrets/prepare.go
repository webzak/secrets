package main

import (
	"flag"
	"fmt"

	"github.com/webzak/secrets"
)

func prepare() error {
	var storage = flag.String("storage", "", "")
	var prefix = flag.String("prefix", "", "")
	var uppercase = flag.Bool("uppercase", false, "")

	flag.Parse()
	var err error
	switch *storage {
	case "enc-env":
		err = prepareEncEnv(*prefix, *uppercase)
	default:
		err = fmt.Errorf("Storage '%s' is not supported", *storage)
	}
	return err
}

func prepareEncEnv(prefix string, uppercase bool) error {
	fmt.Println()
	master, err := readPassword("storage password: ")
	if err != nil {
		return err
	}
	es, err := secrets.NewEncEnvStorage(master, prefix, uppercase)
	if err != nil {
		return err
	}
	var name string
	var secret string
	secrets := make(map[string]string)
	fmt.Printf("\nEnter secret names and values below (empty name to stop)\n\n")
	for {
		name, err = readLine("secret name: ")
		if err != nil {
			return err
		}
		if name == "" {
			break
		}
		secret, err = readPassword("secret value: ")
		if err != nil {
			return err
		}
		if secret == "" {
			break
		}
		if name, secret, err = es.Prepare(name, secret); err != nil {
			return err
		}
		secrets[name] = secret
	}
	fmt.Printf("\n\nResult: \n\n")
	for n, s := range secrets {
		fmt.Printf("export %s=%s\n", n, s)
	}
	fmt.Println()
	return nil
}
