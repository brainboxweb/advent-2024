// Package calculations contains the calculator
package calculations

import (
	"fmt"
	"strconv"
)

// NewCalculator returns a Calculator
func NewCalculator(operators []string) *Calculator {
	return &Calculator{Operators: operators}
}

// Calculator does the calculations
type Calculator struct {
	Operators []string
}

// Calculate performs the operations
func (c Calculator) Calculate(subtotal int, operands []int) []int {
	return calcRecursive([]int{subtotal}, operands, c.Operators)
}

func calcRecursive(subtotals []int, operands []int, operators []string) []int {
	for i, subtotal := range subtotals {
		for _, operator := range operators {
			switch operator {
			case "*":
				output := multiply(subtotal, operands[0])
				subtotals[i] = output // update
			case "+":
				output := add(subtotal, operands[0])
				subtotals = append(subtotals, output) // add new
			case "||":
				output := concatenate(subtotal, operands[0])
				subtotals = append(subtotals, output) // add new
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

func multiply(subtotal, operand int) int {
	return subtotal * operand
}

func add(subtotal, operand int) int {
	return subtotal + operand
}

func concatenate(subtotal, operand int) int {
	con := fmt.Sprintf("%d%d", subtotal, operand)
	concated, err := strconv.Atoi(con)
	if err != nil {
		panic("not expected")
	}

	return concated
}
