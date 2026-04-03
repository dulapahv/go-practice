# go-practice

Practice project for learning parser generation in Go using ANTLR4.

This repository contains:

- A custom grammar in `Calc.g4`
- Generated lexer/parser code in `parser/`
- A Go CLI evaluator in `main.go`

The language is a tiny arithmetic DSL with custom tokens:

- Numbers: `ny` followed by zero or more `a` characters (for example `ny`, `nya`, `nyaa`), and `nyu` for zero
- Operators:
  - `nyya` for addition
  - `nyyya` for subtraction
  - `nyyyya` for multiplication
  - `nyyyyya` for division

## Generate parser code

`java -jar .\antlr-4.13.1-complete.jar -Dlanguage=Go -o parser Calc.g4`

## Run

`go run . "nya nyya nyaa nyyyya nyaaa"`

Or run without arguments and enter an expression interactively:

`go run .`
