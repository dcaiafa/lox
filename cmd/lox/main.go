package main

import (
	"flag"
	"fmt"
	gotoken "go/token"
	"os"
	"runtime/pprof"

	"github.com/dcaiafa/lox/internal/codegen"
	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/parsergen/lr1"
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
		flagAnalyze = flag.Bool("analyze", false, "")
		flagProf    = flag.String("cpu-prof", "", "")
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

	errLogger := errlogger.New()

	grammar := codegen.ParseGrammar(fset, dir, errLogger)
	if errLogger.HasError() {
		return fmt.Errorf("failed to parse grammar")
	}

	parserTable := lr1.ConstructLALR(grammar)
	if *flagAnalyze {
		parserTable.Print(os.Stdout)
		return nil
	}

	if parserTable.HasConflicts {
		return fmt.Errorf("conflicts detected. Re-run with -analyze for details")
	}

	cfg := &codegen.Config{
		Errs:        errLogger,
		ImplDir:     dir,
		Grammar:     grammar,
		ParserTable: parserTable,
	}

	if *flagAnalyze {
		cfg.AnalysisWriter = os.Stdout
		cfg.AnalysisOnly = true
	}

	codegen.Generate(cfg)

	if cfg.Errs.HasError() {
		return fmt.Errorf("errors ocurred")
	}

	return nil
}
