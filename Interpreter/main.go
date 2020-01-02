package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
   Type: Behavioral
   Purpose: Given a language, define a representation for its grammar along with an interpreter that uses the representation to interpret sentences in the language.
   Additional: - when there is a language to interpret, and you can represent statements in the language as tree-like data structure.
*/

// Abstract Expression
type Expression interface {
	interpret() int
}

// Terminal Expression
type NumberExpression int

func (e *NumberExpression) interpret() int {
	return int(*e)
}

// NonTerminal Expression
type SubtractExpression struct {
	leftExpr  Expression
	rightExpr Expression
}

func (e *SubtractExpression) interpret() int {
	return e.leftExpr.interpret() - e.rightExpr.interpret()
}

type AddExpression struct {
	leftExpr  Expression
	rightExpr Expression
}

func (e *AddExpression) interpret() int {
	return e.leftExpr.interpret() + e.rightExpr.interpret()
}

// Parser generate the expressions, it is not part of the Interpreter design pattern
func parsePostfix(targetString string) Expression {
	var stack []Expression
	expressions := strings.Split(targetString, " ")
	for _, expr := range expressions {
		// is number (terminal expression)
		if v, err := strconv.Atoi(expr); err == nil {
			num := NumberExpression(v)
			stack = append(stack, &num)
		} else {
			// non terminal expression
			if expr != "+" && expr != "-" {
				continue
			}
			rightExpr := stack[len(stack)-1]
			leftExpr := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if expr == "+" {
				stack = append(stack, &AddExpression{leftExpr, rightExpr})
			} else if expr == "-" {
				stack = append(stack, &SubtractExpression{leftExpr, rightExpr})
			}
		}
	}
	return stack[0]
}

func main() {
	targetString := "5 4 - 10 +" // 5 - 4 + 10 = 11
	expr := parsePostfix(targetString)
	fmt.Println(expr.interpret())
}
