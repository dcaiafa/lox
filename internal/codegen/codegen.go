package codegen

import (
	gotoken "go/token"
	"io"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

const prefix = "_"

type Config struct {
	Errs           *errlogger.ErrLogger
	ImplDir        string
	Grammar        *grammar.AugmentedGrammar
	AnalysisWriter io.Writer
	AnalysisOnly   bool
}

func Generate(cfg *Config) {
	pgen := newParserGenState(cfg.ImplDir, cfg.Grammar, cfg.Errs)
	lgen := NewLexerGenState(cfg.ImplDir, cfg.Grammar)

	pgen.ConstructParseTables()
	if cfg.Errs.HasError() {
		return
	}

	if cfg.AnalysisWriter != nil {
		pgen.parserTable.Print(cfg.AnalysisWriter)
	}

	if cfg.AnalysisOnly {
		return
	}

	// TODO: check for conflicts

	err := lgen.Generate()
	if err != nil {
		cfg.Errs.Errorf(gotoken.Position{}, "failed to generate base.gen.go: %v", err)
		return
	}

	pgen.ParseGo()
	if cfg.Errs.HasError() {
		return
	}

	pgen.assignActions()
	if cfg.Errs.HasError() {
		return
	}

	err = pgen.Generate()
	if err != nil {
		cfg.Errs.Errorf(gotoken.Position{}, "failed to emit parser.gen.go: %v", err)
		return
	}
}
