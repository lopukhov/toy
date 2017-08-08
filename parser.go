// parser
package color

import (
	"strconv"
)

// Parser interface to build an Interpreter.
type Parser interface {
	Parse() Operator
}

// Te parser will return an Operator interface.
// This interface provides a method that the
// Interpreter can use to interpret the parsed
// expresion.
type Operator interface {
	Operate() int
}

// One implementation of the Operator interface.
// A Node of a tree structure, that contains an
// Operation to apply to it children.
type Node struct {
	Left      *Node
	Right     *Node
	Operation string
}

// Node Operate implementation.
func (n *Node) Operate() int {
	switch n.Operation {
	case "+":
		return n.Left.Operate() + n.Right.Operate()
	case "-":
		return n.Left.Operate() - n.Right.Operate()
	case "*":
		return n.Left.Operate() * n.Right.Operate()
	case "/":
		return n.Left.Operate() / n.Right.Operate()
	}
	i, _ := strconv.Atoi(n.Operation)
	return i
}

// One simple implementation of the Parser interface.
type SimpleParser struct {
	Lex       Lexer
	CurrToken Token
}

// Consumes a Token of Type tokenType.
func (sp *SimpleParser) eat(tokenType string) {
	if sp.CurrToken.Type == tokenType {
		sp.CurrToken = sp.Lex.Lex()
	}
}

// Implements the factor rule:
//
// factor: INT | LPAR expr RPAR
//
func (sp *SimpleParser) factor() *Node {
	if t := sp.CurrToken; t.Type == "INT" {
		sp.eat("INT")
		return &Node{nil, nil, t.Value}
	} else {
		sp.eat("LPAR")
		node := sp.expr()
		sp.eat("RPAR")
		return node
	}
}

// Implements the term rule:
//
// term: factor((MUL|DIV) factor)*
//
func (sp *SimpleParser) term() *Node {
	node := sp.factor()
	for value := sp.CurrToken.Value; value == "*" || value == "/"; value = sp.CurrToken.Value {
		sp.eat("OPR")
		node = &Node{node, sp.factor(), value}
	}
	return node
}

// Implements the expr rule:
//
// expr: term((SUM|SUBS) term)*
//
func (sp *SimpleParser) expr() *Node {
	node := sp.term()
	for value := sp.CurrToken.Value; value == "+" || value == "-"; value = sp.CurrToken.Value {
		sp.eat("OPR")
		node = &Node{node, sp.term(), value}
	}
	return node
}

// Parses the expresion into a tree.
// Return the root of the tree as an
// Operator
func (sp *SimpleParser) Parse() Operator {
	return sp.expr()
}