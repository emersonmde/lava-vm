package class

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

type ExceptionTableEntry struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

// See https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7
type Code struct {
	AttributeNameIndex   uint16
	AttributeLength      uint32
	MaxStack             uint16
	MaxLocals            uint16
	CodeLength           uint32
	Bytecode             []byte
	ExceptionTableLength uint16
	ExceptionTable       []ExceptionTableEntry
	AttributesCount      uint16
	Attributes           []Attribute
}

func parseCodeAttribute(attr *Attribute) (*Code, error) {
	reader := bytes.NewReader(attr.Info)
	code := &Code{}

	if err := binary.Read(reader, binary.BigEndian, &code.MaxStack); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &code.MaxLocals); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &code.CodeLength); err != nil {
		return nil, err
	}
	code.Bytecode = make([]byte, code.CodeLength)
	if err := binary.Read(reader, binary.BigEndian, &code.Bytecode); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &code.ExceptionTableLength); err != nil {
		return nil, err
	}
	code.ExceptionTable = make([]ExceptionTableEntry, code.ExceptionTableLength)
	for i := range code.ExceptionTable {
		if err := binary.Read(reader, binary.BigEndian, &code.ExceptionTable[i]); err != nil {
			return nil, err
		}
	}

	return code, nil
}

func bytecodeToHex(bytecode []byte) string {
	hexCodes := make([]string, len(bytecode))
	for i, code := range bytecode {
		hexCodes[i] = fmt.Sprintf("%02x", code)
	}
	return strings.Join(hexCodes, " ")
}

func (e *ExceptionTableEntry) String() string {
	return fmt.Sprintf("StartPc: %d, EndPc: %d, HandlerPc: %d, CatchType: %d", e.StartPc, e.EndPc, e.HandlerPc, e.CatchType)
}

func (c *Code) String() string {
	var exceptionTableEntries []string
	for _, entry := range c.ExceptionTable {
		exceptionTableEntries = append(exceptionTableEntries, entry.String())
	}
	exceptionTableStr := strings.Join(exceptionTableEntries, "\n")

	return fmt.Sprintf("MaxStack: %d\nMaxLocals: %d\nBytecode: %s\nExceptionTable:\n%s",
		c.MaxStack, c.MaxLocals, bytecodeToHex(c.Bytecode), exceptionTableStr)
}
