package main

import (
	"fmt"
	gotoken "go/token"
	"io"
	"os"
	"runtime/pprof"

	"github.com/dcaiafa/lox/internal/base/errlogger"
	"github.com/dcaiafa/lox/internal/codegen"
	"github.com/spf13/pflag"
)

func usage() {
	out := pflag.CommandLine.Output()
	fmt.Fprintf(out, "Parser and lexer generator for Go.\n")
	fmt.Fprintf(out, "https://dcaiafa.github.io/lox\n")
	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "Usage:\n")
	fmt.Fprintf(out, "  lox [flags] <package-path>\n")
	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "Flags:\n")
	fmt.Fprintf(out, "  --report   Show detailed analysis report\n")
	fmt.Fprintf(out, "  --help/-h  Print this help\n")
	fmt.Fprintf(out, "\n")
}

func main() {
	err := realMain()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func realMain() error {
	var (
		flagReport = pflag.Bool("report", false, "")
		flagProf   = pflag.String("cpu-prof", "", "")
		flagHelp   = pflag.BoolP("help", "h", false, "")
	)

	pflag.Usage = usage
	pflag.Parse()

	if *flagHelp {
		pflag.Usage()
		return nil
	}

	if pflag.NArg() != 1 {
		pflag.Usage()
		return fmt.Errorf("<path> required")
	}
	dir := pflag.Arg(0)

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
