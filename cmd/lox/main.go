package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dcaiafa/lox/internal/codegen"
)

func main() {
	err := realMain()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func realMain() error {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		return fmt.Errorf("<path> required")
	}
	dir := flag.Arg(0)

	state := codegen.NewState(dir, dir)
	err := state.ParseGrammar()
	if err != nil {
		return err
	}
	state.ConstructParseTables()
	err = state.ParseGo()
	if err != nil {
		return err
	}
	err = state.MapReduceActions()
	if err != nil {
		return err
	}
	err = state.Generate()
	if err != nil {
		return err
	}

	return nil
}