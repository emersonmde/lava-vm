package execution_engine

import (
	"errors"
	"fmt"
	"lava-vm/pkg/class"
)

type Code = class.Code

type Instruction struct {
	Opcode   byte
	Operands []byte
}

func ParseInstructions(c *Code) ([]Instruction, error) {
	instructions := []Instruction{}
	pc := 0
	for pc < len(c.Bytecode) {
		opcode := c.Bytecode[pc]
		pc += 1

		var operands []byte
		switch opcode {
		case 0xbb: // new
			if pc+2 > len(c.Bytecode) {
				return nil, errors.New("unexpected end of bytecode")
			}
			operands = c.Bytecode[pc : pc+2]
			pc += 2
		// TODO: MOAR OPCODES
		default:
			fmt.Printf("Unknown opcode 0x%02X\n", opcode)
			continue
		}

		instructions = append(instructions, Instruction{
			Opcode:   opcode,
			Operands: operands,
		})
	}

	return instructions, nil
}
