// Package computer represents the computer
package computer

import (
	"math"
	"strconv"
	"strings"
)

// New returns a new computer
func New() *Computer {
	return &Computer{}
}

// Computer represents the computer
type Computer struct {
	regA               int
	regB               int
	regC               int
	output             []string
	instructionPointer int
}

// SetRegisterA sets the value of Register A
func (c *Computer) SetRegisterA(val int) {
	c.regA = val
}

//revive:disable:cyclomatic

// Run runs the specified program
func (c *Computer) Run(input []int) string {
	c.instructionPointer = 0
	for c.instructionPointer < len(input)-1 {
		opCode := input[c.instructionPointer]
		literalOperand := input[c.instructionPointer+1]
		comboOperand := c.comboOperand(literalOperand)
		switch opCode {
		case 0:
			c.adv(comboOperand)
		case 1:
			c.bxl(literalOperand)
		case 2:
			c.bst(comboOperand)
		case 3:
			if c.regA != 0 {
				c.instructionPointer = literalOperand
				continue
			}
		case 4:
			c.bxc(comboOperand) // operand ignored
		case 5:
			c.out(comboOperand)
		case 6:
			c.bdv(comboOperand)
		case 7:
			c.cdv(comboOperand)
		}
		c.instructionPointer += 2
	}
	ret := strings.Join(c.output, ",")
	c.output = nil
	return ret
}

//revive:enable:cyclomatic

func (c *Computer) comboOperand(val int) int {
	switch val {
	case 4:
		val = c.regA
	case 5:
		val = c.regB
	case 6:
		val = c.regC
	}
	return val
}

func (c *Computer) adv(val int) {
	c.regA = dv(c.regA, val)
}

func dv(regA, val int) int {
	numerator := regA
	denominator := math.Pow(2, float64(val))
	result := int(float64(numerator) / denominator)
	return result
}

// The bxl instruction (opcode 1) calculates the bitwise XOR of register B
// and the	instruction's literal operand, then stores the result in
// register B.
func (c *Computer) bxl(literalOperand int) {
	res := c.regB ^ literalOperand
	c.regB = res
}

// The bst instruction (opcode 2) calculates the value of its combo operand
// modulo 8 (thereby keeping only its lowest 3 bits), then writes that
// value to the B register.
func (c *Computer) bst(val int) {
	c.regB = val % 8
}

// The bxc instruction (opcode 4) calculates the bitwise XOR of register B
// and register C, then stores the result in register B. (For legacy reasons,
// this instruction reads an operand but ignores it.)
func (c *Computer) bxc(val int) {
	_ = val
	res := c.regB ^ c.regC
	c.regB = res
}

// The out instruction (opcode 5) calculates the value of its combo operand
// modulo 8, then outputs that value. (If a program outputs multiple values,
// they are separated by commas.)
func (c *Computer) out(val int) {
	res := val % 8
	c.output = append(c.output, strconv.Itoa(res))
}

// The bdv instruction (opcode 6) works exactly like the adv instruction
// except that	the result is stored in the B register. (The numerator is
// still read from the A register.)
func (c *Computer) bdv(val int) {
	c.regB = dv(c.regA, val)
}

// The cdv instruction (opcode 7) works exactly like the adv instruction
// except that the result is stored in the C register. (The numerator is
// still read from the A register.)
func (c *Computer) cdv(val int) {
	c.regC = dv(c.regA, val)
}
