package class_file

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

type Method struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []Attribute
}

// Read a Method from the given file
func readMethod(file *os.File, method *Method) error {
	if err := binary.Read(file, binary.BigEndian, &method.AccessFlags); err != nil {
		return fmt.Errorf("reading access flags: %w", err)
	}

	if err := binary.Read(file, binary.BigEndian, &method.NameIndex); err != nil {
		return fmt.Errorf("reading name index: %w", err)
	}

	if err := binary.Read(file, binary.BigEndian, &method.DescriptorIndex); err != nil {
		return fmt.Errorf("reading descriptor index: %w", err)
	}

	if err := binary.Read(file, binary.BigEndian, &method.AttributesCount); err != nil {
		return fmt.Errorf("reading attributes count: %w", err)
	}

	method.Attributes = make([]Attribute, method.AttributesCount)
	for i := range method.Attributes {
		if err := readAttribute(file, &method.Attributes[i]); err != nil {
			return fmt.Errorf("reading attribute %d: %w", i, err)
		}
	}

	return nil
}

func (m Method) String() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "Access Flags: 0x%04X\n", m.AccessFlags)
	fmt.Fprintf(&builder, "Name Index: %d\n", m.NameIndex)
	fmt.Fprintf(&builder, "Descriptor Index: %d\n", m.DescriptorIndex)
	fmt.Fprintf(&builder, "Attributes Count: %d\n", m.AttributesCount)
	for i, attribute := range m.Attributes {
		fmt.Fprintf(&builder, "Attribute #%d: %s\n", i+1, attribute)
	}
	return builder.String()
}
