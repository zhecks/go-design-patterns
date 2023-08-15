package interpreter

import (
	"strconv"
	"strings"
)

type Node interface {
	Interpret() int
}

type ValNode struct {
	val int
}

func (n *ValNode) Interpret() int {
	return n.val
}

type AddNode struct {
	left, right Node
}

func (n *AddNode) Interpret() int {
	return n.left.Interpret() + n.right.Interpret()
}

type SubNode struct {
	left, right Node
}

func (n *SubNode) Interpret() int {
	return n.left.Interpret() - n.right.Interpret()
}

type Parser struct {
	exp   []string
	index int
	prev  Node
}

func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")
	for p.index < len(p.exp) {
		switch p.exp[p.index] {
		case "+":
			p.prev = p.newAddNode()
		case "-":
			p.prev = p.newSubNode()
		default:
			p.prev = p.newValNode()
		}
	}
}

func (p *Parser) newAddNode() Node {
	p.index++
	return &AddNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

func (p *Parser) newSubNode() Node {
	p.index++
	return &SubNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

func (p *Parser) newValNode() Node {
	val, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &ValNode{
		val: val,
	}
}

func (p *Parser) Result() int {
	return p.prev.Interpret()
}
