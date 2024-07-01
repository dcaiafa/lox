package main

import (
	"io"
	"log"
	"os"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	program, err := Parse("program.bolox", data)
	if err != nil {
		log.Fatal(err)
	}

	err = program.Run()
	if err != nil {
		log.Fatal(err)
	}
}
