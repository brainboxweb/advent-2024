package day7

import (
	"fmt"
	"strconv"
	"strings"
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

func run(equations []Equation, operators []string) int {
	for i, equ := range equations {
		results := Calculate(equ, operators)
		for _, res := range results {
			if res == equ.result {
				equations[i].works = true // mark as true
			}
		}
	}
	total := 0
	for _, equ := range equations {
		if equ.works {
			total += equ.result
		}
	}

	return total
}

func Calculate(equ Equation, operators []string) []int {
	subtotal, operands := equ.Operands[0], equ.Operands[1:]
	return calcRecursive([]int{subtotal}, operands, operators)
}

func calcRecursive(subtotals []int, operands []int, operators []string) []int {
	for i, subtotal := range subtotals {
		for _, operator := range operators {
			output := 0
			switch operator {
			case "*":
				output = multiply(subtotal, operands[0])
				subtotals[i] = output // update one
			case "+":
				output = add(subtotal, operands[0])
				subtotals = append(subtotals, output) // and a new one
			case "||":
				output = concatenate(subtotal, operands[0])
				subtotals = append(subtotals, output) // and a new one
			default:
				panic("not implemented")
			}
		}
	}
	operands = operands[1:]
	if len(operands) < 1 {
		return subtotals
	}

	return calcRecursive(subtotals, operands, operators)
}

type Equation struct {
	result   int
	Operands []int
	works    bool
}

func multiply(subtotal, operand int) int {
	return subtotal * operand
}

func add(subtotal, operand int) int {
	return subtotal + operand
}

func concatenate(subtotal, operand int) int {
	con := fmt.Sprintf("%d%d", subtotal, operand)
	concated, _ := strconv.Atoi(con)
	return concated
}

func parse(data []string) []Equation {
	var ret []Equation
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
		eq := Equation{result: result, Operands: operands}
		ret = append(ret, eq)
	}

	return ret
}
