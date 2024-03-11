package main

import (
	"flag"
	"fmt"

	"github.com/webzak/secrets"
)

func prepare() error {
	storage := flag.String("storage", "", "")
	prefix := flag.String("prefix", "", "")
	uppercase := flag.Bool("uppercase", false, "")

	flag.Parse()
	var err error
	switch *storage {
	case "enc-env":
		err = prepareEncEnv(*prefix, *uppercase)
	case "enc-dump":
		err = prepareEncDump()
	default:
		err = fmt.Errorf("Storage '%s' is not supported", *storage)
	}
	return err
}

func prepareEncEnv(prefix string, uppercase bool) error {
	fmt.Println()
	master, err := secrets.ReadPassword("storage password: ")
	if err != nil {
		return err
	}
	es, err := secrets.NewEncEnvStorage(master, prefix, uppercase)
	if err != nil {
		return err
	}
	var name string
	var secret string
	data := make(map[string]string)
	fmt.Printf("\nEnter secret names and values below (empty name to stop)\n\n")
	for {
		name, err = secrets.ReadLine("secret name: ")
		if err != nil {
			return err
		}
		if name == "" {
			break
		}
		secret, err = secrets.ReadPassword("secret value: ")
		if err != nil {
			return err
		}
		if secret == "" {
			break
		}
		if name, secret, err = es.Prepare(name, secret); err != nil {
			return err
		}
		data[name] = secret
	}
	fmt.Printf("\n\nResult: \n\n")
	for n, s := range data {
		fmt.Printf("export %s=%s\n", n, s)
	}
	fmt.Println()
	return nil
}

func prepareEncDump() error {
	fmt.Println()
	master, err := secrets.ReadPassword("storage password: ")
	if err != nil {
		return err
	}
	es, err := secrets.NewEncMapStorage(master, "")
	if err != nil {
		return err
	}
	var name string
	var secret string
	fmt.Printf("\nEnter secret names and values below (empty name to stop)\n\n")
	for {
		name, err = secrets.ReadLine("secret name: ")
		if err != nil {
			return err
		}
		if name == "" {
			break
		}
		secret, err = secrets.ReadPassword("secret value: ")
		if err != nil {
			return err
		}
		if secret == "" {
			break
		}
		es.Set(name, secret)
	}
	dump, err := es.Dump()
	if err != nil {
		return err
	}
	fmt.Println(dump)
	return nil
}
