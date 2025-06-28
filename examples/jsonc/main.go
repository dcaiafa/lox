package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %v <input.json>\n", args[0])
		os.Exit(1)
	}

	inputFile := args[1]
	var jsonData []byte
	var err error

	if inputFile == "-" {
		jsonData, err = io.ReadAll(os.Stdin)
	} else {
		jsonData, err = os.ReadFile(inputFile)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	res, err := Parse(string(jsonData))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	err = enc.Encode(res)
	if err != nil {
		panic(err)
	}
}
