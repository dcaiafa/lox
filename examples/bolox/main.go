package main

import (
	"flag"
	"fmt"
	gotoken "go/token"
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

	fset := gotoken.NewFileSet()

	program, err := Parse(fset, filename, data)
	if err != nil {
		log.Fatal(err)
	}

	ctx := NewContext(fset)

	err = program.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
