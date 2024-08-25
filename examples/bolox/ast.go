package main

import (
	"fmt"
	gotoken "go/token"
	"reflect"
	"strings"
)

type Bounds struct {
	Begin gotoken.Pos
	End   gotoken.Pos
}

type AST interface {
	Bounds() Bounds
	SetBounds(b Bounds)
	Discard() bool
}

type baseAST struct {
	bounds Bounds
}

func (b *baseAST) Bounds() Bounds {
	return b.bounds
}

func (b *baseAST) SetBounds(v Bounds) {
	b.bounds = v
}

func (b *baseAST) Discard() bool { return false }

func errorWithPos(ctx *Context, ast AST, err error) error {
	return fmt.Errorf("%w\nfrom: %v", err, ctx.FileSet.Position(ast.Bounds().Begin))
}

type Control int

const (
	Step Control = iota
	Continue
	Return
)

type Statement interface {
	AST
	Run(ctx *Context) (Control, error)
}

type Program struct {
	baseAST
	Block *Block
}

func (p *Program) Run(ctx *Context) error {
	step, err := p.Block.Run(ctx)
	if err != nil {
		return err
	}

	switch step {
	case Continue:
		return fmt.Errorf("continue used outside of while")
	case Step:
	default:
		panic("unreachable")
	}

	return nil
}

type Block struct {
	baseAST
	Statements []Statement
}

func (b *Block) Run(ctx *Context) (Control, error) {
	for _, stmt := range b.Statements {
		ctrl, err := stmt.Run(ctx)
		if err != nil {
			return 0, err
		}
		if ctrl != Step {
			return ctrl, nil
		}
	}
	return Step, nil
}

type FuncCallStatement struct {
	baseAST
	FuncCall *FuncCall
}

func (s *FuncCallStatement) Run(ctx *Context) (Control, error) {
	_, err := s.FuncCall.Eval(ctx)
	if err != nil {
		return 0, errorWithPos(ctx, s.FuncCall, err)
	}
	return Step, err
}

type VarAssign struct {
	baseAST

	VarName string
	Value   Expr
}

func (a *VarAssign) Run(ctx *Context) (Control, error) {
	v, err := a.Value.Eval(ctx)
	if err != nil {
		return Step, err
	}
	ctx.SetGlobal(a.VarName, v)
	return Step, nil
}

type Expr interface {
	AST
	Eval(ctx *Context) (any, error)
}

type FuncCall struct {
	baseAST

	FuncName string
	Args     []Expr
}

func (c *FuncCall) Eval(ctx *Context) (any, error) {
	var err error
	vals := make([]any, len(c.Args))
	for i, arg := range c.Args {
		vals[i], err = arg.Eval(ctx)
		if err != nil {
			return nil, err
		}
	}
	return ctx.Call(c.FuncName, vals)
}

type VarRef struct {
	baseAST

	VarName string
}

func (r *VarRef) Eval(ctx *Context) (any, error) {
	v, ok := ctx.GetGlobal(r.VarName)
	if !ok {
		return nil, errorWithPos(ctx, r, fmt.Errorf("undefined: %v", r.VarName))
	}
	return v, nil
}

func doPrint(vals []any) (any, error) {
	fmt.Println(vals...)
	return nil, nil
}

type Literal struct {
	baseAST

	Val any
}

func (l *Literal) Eval(ctx *Context) (any, error) {
	return l.Val, nil
}

type Op int

const (
	OpPlus Op = iota
	OpMinus
	OpTimes
	OpDiv
	OpLT
	OpLE
	OpGT
	OpGE
	OpEq
	OpNE
	OpAnd
	OpOr
)

type BinaryExpr struct {
	baseAST

	Left  Expr
	Right Expr
	Op    Op
}

func (e *BinaryExpr) Eval(ctx *Context) (res any, err error) {
	defer func() {
		if err != nil {
			err = errorWithPos(ctx, e, err)
		}
	}()
	switch e.Op {
	case OpAnd:
		return e.evalAnd(ctx)
	case OpOr:
		return e.evalOr(ctx)
	}

	va, err := e.Left.Eval(ctx)
	if err != nil {
		return nil, err
	}

	vb, err := e.Right.Eval(ctx)
	if err != nil {
		return nil, err
	}

	switch e.Op {
	case OpEq:
		return va == vb, nil
	case OpNE:
		return va != vb, nil
	}

	if a, b, ok := castBinaryExpr[int](va, vb); ok {
		switch e.Op {
		case OpPlus:
			return a + b, nil
		case OpMinus:
			return a - b, nil
		case OpTimes:
			return a * b, nil
		case OpDiv:
			if b == 0 {
				return nil, fmt.Errorf("division by zero")
			}
			return a / b, nil
		case OpLT:
			return a < b, nil
		case OpLE:
			return a <= b, nil
		case OpGT:
			return a > b, nil
		case OpGE:
			return a >= b, nil
		default:
			panic("unreachable")
		}
	} else if a, b, ok := castBinaryExpr[string](va, vb); ok {
		switch e.Op {
		case OpPlus:
			return a + b, nil
		case OpLT:
			return a < b, nil
		case OpLE:
			return a <= b, nil
		case OpGT:
			return a > b, nil
		case OpGE:
			return a >= b, nil
		default:
			return nil, fmt.Errorf("operation not supported by string")
		}
	} else {
		return nil, fmt.Errorf(
			"operation not supported between %v and %v",
			reflect.TypeOf(va), reflect.TypeOf(vb))
	}
}

func (e *BinaryExpr) evalAnd(ctx *Context) (res any, err error) {
	defer func() {
		if err != nil {
			err = errorWithPos(ctx, e, err)
		}
	}()

	va, err := e.Left.Eval(ctx)
	if err != nil {
		return nil, err
	}
	ba, ok := va.(bool)
	if !ok {
		return nil, fmt.Errorf(
			"operation not supported for %v",
			reflect.TypeOf(va))
	}
	if !ba {
		return false, nil
	}
	vb, err := e.Right.Eval(ctx)
	if err != nil {
		return nil, err
	}
	bb, ok := vb.(bool)
	if !ok {
		return nil, fmt.Errorf(
			"operation not supported for %v",
			reflect.TypeOf(vb))
	}
	return ba && bb, nil
}

func (e *BinaryExpr) evalOr(ctx *Context) (res any, err error) {
	defer func() {
		if err != nil {
			err = errorWithPos(ctx, e, err)
		}
	}()

	va, err := e.Left.Eval(ctx)
	if err != nil {
		return nil, err
	}
	ba, ok := va.(bool)
	if !ok {
		return nil, fmt.Errorf(
			"operation not supported for %v",
			reflect.TypeOf(va))
	}
	if ba {
		return true, nil
	}
	vb, err := e.Right.Eval(ctx)
	if err != nil {
		return nil, err
	}
	bb, ok := vb.(bool)
	if !ok {
		return nil, fmt.Errorf(
			"operation not supported for %v",
			reflect.TypeOf(vb))
	}
	return ba || bb, nil
}

func castBinaryExpr[T any](a, b any) (ca, cb T, ok bool) {
	ca, ok = a.(T)
	if !ok {
		return ca, cb, false
	}
	cb, ok = b.(T)
	if !ok {
		return ca, cb, false
	}
	return ca, cb, true
}

type While struct {
	baseAST

	Pred  Expr
	Block *Block
}

func (w *While) Run(ctx *Context) (Control, error) {
	for {
		pred, err := evalPredicate(ctx, w.Pred)
		if err != nil {
			return 0, err
		}
		if !pred {
			break
		}
		ctrl, err := w.Block.Run(ctx)
		if err != nil {
			return 0, err
		}
		switch ctrl {
		case Continue:
			continue
		case Step:
		default:
			panic("unreachable")
		}
	}
	return Step, nil
}

type ContinueStatement struct {
	baseAST
}

func (s *ContinueStatement) Run(ctx *Context) (Control, error) {
	return Continue, nil
}

type IfStatement struct {
	baseAST

	Pred  Expr
	Block *Block
	Elifs []*Elif
	Else  *Else
}

func (s *IfStatement) Run(ctx *Context) (Control, error) {
	v, err := evalPredicate(ctx, s.Pred)
	if err != nil {
		return 0, err
	}
	if v {
		return s.Block.Run(ctx)
	}

	for _, elif := range s.Elifs {
		v, err := evalPredicate(ctx, elif.Pred)
		if err != nil {
			return 0, err
		}
		if v {
			return elif.Block.Run(ctx)
		}
	}

	if s.Else != nil {
		return s.Else.Block.Run(ctx)
	}

	return Step, nil
}

type Elif struct {
	baseAST

	Pred  Expr
	Block *Block
}

type Else struct {
	baseAST

	Block *Block
}

type Noop struct {
	baseAST
}

func (n *Noop) Run(ctx *Context) (Control, error) { return Step, nil }
func (n *Noop) Discard() bool                     { return true }

func evalPredicate(ctx *Context, p Expr) (res bool, err error) {
	defer func() {
		if err != nil {
			err = errorWithPos(ctx, p, err)
		}
	}()
	v, err := p.Eval(ctx)
	if err != nil {
		return false, err
	}

	vb, ok := v.(bool)
	if !ok {
		return false, fmt.Errorf(
			"predicate must be bool, not %T", v)
	}

	return vb, nil
}

type String struct {
	baseAST

	Parts []Expr
}

func (s *String) Eval(ctx *Context) (any, error) {
	var res strings.Builder
	for _, part := range s.Parts {
		v, err := part.Eval(ctx)
		if err != nil {
			return nil, err
		}
		fmt.Fprintf(&res, "%v", v)
	}
	return res.String(), nil
}

type StringCharSeq struct {
	baseAST
	Seq string
}

func (s *StringCharSeq) Eval(ctx *Context) (any, error) {
	return s.Seq, nil
}
