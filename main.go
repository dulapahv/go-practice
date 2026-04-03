package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"gogram/parser"
)

func main() {
	expr := strings.TrimSpace(strings.Join(os.Args[1:], " "))
	if expr == "" {
		fmt.Print("Enter expression: ")
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "error: failed to read input: %v\n", err)
			} else {
				fmt.Fprintln(os.Stderr, "error: no input provided")
			}
			os.Exit(1)
		}
		expr = strings.TrimSpace(scanner.Text())
	}

	result, err := parseAndEval(expr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}

func parseAndEval(expr string) (int, error) {
	if strings.TrimSpace(expr) == "" {
		return 0, fmt.Errorf("expression is empty")
	}

	input := antlr.NewInputStream(expr)
	lexer := parser.NewCalcLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCalcParser(stream)

	start := p.Start_()

	return evalExpr(start.Expression())
}

func evalExpr(ctx parser.IExpressionContext) (int, error) {
	switch node := ctx.(type) {

	// Number literal
	case *parser.NumberContext:
		v, err := strconv.Atoi(node.NUMBER().GetText())
		if err != nil {
			return 0, fmt.Errorf("invalid number %q: %w", node.NUMBER().GetText(), err)
		}
		return v, nil

	// Multiplication or division
	case *parser.MulDivContext:
		left, err := evalExpr(node.Expression(0))
		if err != nil {
			return 0, err
		}
		right, err := evalExpr(node.Expression(1))
		if err != nil {
			return 0, err
		}

		switch node.GetOp().GetTokenType() {

		// Multiplication
		case parser.CalcParserMUL:
			return left * right, nil

		// Division
		case parser.CalcParserDIV:
			if right == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			return left / right, nil

		default:
			return 0, fmt.Errorf("unsupported operator %q", node.GetOp().GetText())
		}

	// Addition or subtraction
	case *parser.AddSubContext:
		left, err := evalExpr(node.Expression(0))
		if err != nil {
			return 0, err
		}
		right, err := evalExpr(node.Expression(1))
		if err != nil {
			return 0, err
		}

		switch node.GetOp().GetTokenType() {

		// Addition
		case parser.CalcParserADD:
			return left + right, nil

		// Subtraction
		case parser.CalcParserSUB:
			return left - right, nil

		default:
			return 0, fmt.Errorf("unsupported operator %q", node.GetOp().GetText())
		}

	default:
		return 0, fmt.Errorf("unsupported expression node %T", ctx)
	}
}
