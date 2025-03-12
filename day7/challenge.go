package day7

import (
	"strconv"
	"strings"

	"github.com/brainboxweb/advent-2024/day7/calculations"
)

func ChallengeOne(data []string) int {
	equations := parse(data)
	operators := []string{"*", "+"}

	return run(equations, operators)
}

func ChallengeTwo(data []string) int {
	equations := parse(data)
	operators := []string{"*", "+", "||"}

	return run(equations, operators)
}

func run(equations []equation, operators []string) int {
	calc := calculations.NewCalculator(operators)
	for i, equ := range equations {
		subtotal, operands := equ.Operands[0], equ.Operands[1:]
		results := calc.Calculate(subtotal, operands)
		for _, res := range results {
			if res == equ.Result {
				equations[i].Works = true
				continue
			}
		}
	}
	total := 0
	for _, equ := range equations {
		if equ.Works {
			total += equ.Result
		}
	}

	return total
}

func parse(data []string) []equation {
	var ret []equation
	for _, line := range data {
		parts := strings.Split(line, ": ")
		res := parts[0]
		result, _ := strconv.Atoi(res)
		opers := strings.Split(parts[1], " ")
		operands := []int{}
		for _, oper := range opers {
			operand, _ := strconv.Atoi(oper)
			operands = append(operands, operand)
		}
		eq := equation{Result: result, Operands: operands}
		ret = append(ret, eq)
	}

	return ret
}

type equation struct {
	Result   int
	Operands []int
	Works    bool
}
