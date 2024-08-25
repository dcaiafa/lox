package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: bolox <source.bolox>")
		os.Exit(1)
	}

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
	}

	filename := flag.Arg(0)

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	program, err := Parse(filename, data)
	if err != nil {
		log.Fatal(err)
	}

	ctx := NewContext()

	err = program.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
