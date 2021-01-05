package main

type Expr interface {
}

type Var string

type literal float64

type unary struct {
	op rune // '+', '-'
	x  Expr
}

type binary struct {
	op   rune // '+', '-', '*', '/'
	x, y Expr
}

type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

type Env map[Var]float64
