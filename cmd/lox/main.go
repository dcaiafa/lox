package main

import (
	"flag"
	"fmt"
	gotoken "go/token"
	"io"
	"os"
	"runtime/pprof"

	"github.com/dcaiafa/lox/internal/base/errlogger"
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
	var (
		flagReport = flag.Bool("report", false, "")
		flagProf   = flag.String("cpu-prof", "", "")
	)

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		return fmt.Errorf("<path> required")
	}
	dir := flag.Arg(0)

	if *flagProf != "" {
		f, err := os.Create(*flagProf)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	fset := gotoken.NewFileSet()
	errs := errlogger.New(fset, os.Stderr)

	var reportOut io.Writer
	if *flagReport {
		reportOut = os.Stdout
	}

	ok := codegen.Generate(&codegen.Config{
		Fset:   fset,
		Errs:   errs,
		Dir:    dir,
		Report: reportOut,
	})
	if !ok {
		return fmt.Errorf("errors ocurred")
	}

	return nil
}
