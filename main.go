package main

import (
	"flag"
	"fmt"
	"os"
)

func replace(filepath string) error {
	if !exists(filepath) {
		return fmt.Errorf("%v is not found", filepath)
	}
	println("not implemented.")

	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()

	for _, filepath := range args {
		replace(filepath)
	}
}

func exists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}
