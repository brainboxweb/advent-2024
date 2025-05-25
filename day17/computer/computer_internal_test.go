package computer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Combo operands 0 through 3 represent literal values 0 through 3.
// Combo operand 4 represents the value of register A.
// Combo operand 5 represents the value of register B.
// Combo operand 6 represents the value of register C.
// Combo operand 7 is reserved.
func TestOperand(t *testing.T) {
	c := Computer{regA: 16, regB: 14, regC: 12}
	tests := []struct {
		op       int
		expected int
	}{
		{
			0,
			0,
		},
		{
			1,
			1,
		},
		{
			2,
			2,
		},
		{
			3,
			3,
		},
		{
			4,
			16, // Register A
		},
		{
			5,
			14, // Register B
		},
		{
			6,
			12, // Register C
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			result := c.comboOperand(tt.op)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// The adv instruction (opcode 0) performs division. The numerator is the value
//
//	in the A register. The denominator is found by raising 2 to the power of
//	the instruction's combo operand. (So, an operand of 2 would divide A
//
// by 4 (2^2); an operand of 5 would divide A by 2^B.)
// The result of the division operation is truncated to an integer and then
// written to the A register.
func TestAdv(t *testing.T) {
	tests := []struct {
		start         Computer
		operand       int
		expectedState Computer
	}{
		{
			Computer{regA: 16},
			2,
			Computer{regA: 4},
		},
		{
			Computer{regA: 16, regB: 2},
			3,
			Computer{regA: 2, regB: 2},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			c := tt.start
			c.adv(tt.operand)
			assert.Equal(t, tt.expectedState, c)
		})
	}
}

// The bxl instruction (opcode 1) calculates the bitwise XOR of register B
// and the instruction's literal operand, then stores the result in register B.
func TestBxl(t *testing.T) {
	tests := []struct {
		start         Computer
		operand       int
		expectedState Computer
	}{
		{
			Computer{regB: 1001},
			1000,
			Computer{regB: 1},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			c := tt.start
			c.bxl(tt.operand)
			assert.Equal(t, tt.expectedState, c)
		})
	}
}

// The bst instruction (opcode 2) calculates the value of its combo operand modulo 8
// (thereby keeping only its lowest 3 bits), then writes that value to the B register.
func TestBst(t *testing.T) {
	tests := []struct {
		start         Computer
		operand       int
		expectedState Computer
	}{
		{
			Computer{},
			10,
			Computer{regB: 2},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			c := tt.start
			c.bst(tt.operand)
			assert.Equal(t, tt.expectedState, c)
		})
	}
}

// The bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C,
// then stores the result in register B. (For legacy reasons, this instruction reads an
// operand but ignores it.)
func TestBxc(t *testing.T) {
	tests := []struct {
		start         Computer
		operand       int
		expectedState Computer
	}{
		{
			Computer{regB: 1000, regC: 1001},
			2, // ignored!
			Computer{regB: 1, regC: 1001},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			c := tt.start
			c.bxc(tt.operand)
			assert.Equal(t, tt.expectedState, c)
		})
	}
}

// The out instruction (opcode 5) calculates the value of its combo operand modulo 8,
// then outputs that value. (If a program outputs multiple values, they are separated by commas.)
func TestOut(t *testing.T) {
	tests := []struct {
		start         Computer
		operand       int
		expectedState Computer
	}{
		{
			Computer{},
			10,
			Computer{output: []string{"2"}},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			c := tt.start
			c.out(tt.operand)
			assert.Equal(t, tt.expectedState, c)
		})
	}
}

// The bdv instruction (opcode 6) works exactly like the adv instruction except
// that the result is stored in the B register. (The numerator is still read from the A register.)
func TestBdv(t *testing.T) {
	tests := []struct {
		start         Computer
		operand       int
		expectedState Computer
	}{
		{
			Computer{regA: 16},
			2,
			Computer{regA: 16, regB: 4},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			c := tt.start
			c.bdv(tt.operand)
			assert.Equal(t, tt.expectedState, c)
		})
	}
}

// The cdv instruction (opcode 7) works exactly like the adv instruction except that the
// result is stored in the C register. (The numerator is still read from the A register.)
func TestCdv(t *testing.T) {
	tests := []struct {
		start         Computer
		operand       int
		expectedState Computer
	}{
		{
			Computer{regA: 16},
			2,
			Computer{regA: 16, regC: 4},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			c := tt.start
			c.cdv(tt.operand)
			assert.Equal(t, tt.expectedState, c)
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name           string
		start          Computer
		program        []int
		expectedState  Computer
		expectedOutput string
	}{
		{
			"If register C contains 9, the program 2,6 would set register B to 1.",
			Computer{regC: 9},
			[]int{2, 6},
			Computer{regB: 1, regC: 9, instructionPointer: 2},
			"",
		},
		{
			"If register A contains 10, the program 5,0,5,1,5,4 would output 0,1,2.",
			Computer{regA: 10},
			[]int{5, 0, 5, 1, 5, 4},
			Computer{regA: 10, instructionPointer: 6},
			"0,1,2",
		},
		{
			"If register A contains 2024, the program 0,1,5,4,3,0 would output 4,2,5,6,7,7,7,7,3,1,0 and leave 0 in register A.",
			Computer{regA: 2024},
			[]int{0, 1, 5, 4, 3, 0},
			Computer{regA: 0, instructionPointer: 6},
			"4,2,5,6,7,7,7,7,3,1,0",
		},
		{
			"If register B contains 29, the program 1,7 would set register B to 26.",
			Computer{regB: 29},
			[]int{1, 7},
			Computer{regB: 26, instructionPointer: 2},
			"",
		},

		{
			"If register B contains 2024 and register C contains 43690, the program 4,0 would set register B to 44354.",
			Computer{regB: 2024, regC: 43690},
			[]int{4, 0},
			Computer{regB: 44354, regC: 43690, instructionPointer: 2},
			"",
		},
		{
			"Reg A 729, program: 0,1,5,4,3,0 --> 4,6,3,5,6,3,5,2,1,0.",
			Computer{regA: 729},
			[]int{0, 1, 5, 4, 3, 0},
			Computer{instructionPointer: 6},
			"4,6,3,5,6,3,5,2,1,0",
		},
		{
			"Register A: 52042868, 2,4,1,7,7,5,0,3,4,4,1,7,5,5,3,0 --> ?",
			Computer{regA: 52042868},
			[]int{2, 4, 1, 7, 7, 5, 0, 3, 4, 4, 1, 7, 5, 5, 3, 0},
			Computer{regB: 3, instructionPointer: 16},
			"2,1,0,1,7,2,5,0,3", // <-- Today's solution
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.start
			result := c.Run(tt.program)
			assert.Equal(t, tt.expectedState, c)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}
