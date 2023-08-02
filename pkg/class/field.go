package class

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

type Field struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []Attribute
}

// Read a Field from the given file
func readField(file *os.File, field *Field) error {
	if err := binary.Read(file, binary.BigEndian, &field.AccessFlags); err != nil {
		return fmt.Errorf("reading access flags: %w", err)
	}

	if err := binary.Read(file, binary.BigEndian, &field.NameIndex); err != nil {
		return fmt.Errorf("reading name index: %w", err)
	}

	if err := binary.Read(file, binary.BigEndian, &field.DescriptorIndex); err != nil {
		return fmt.Errorf("reading descriptor index: %w", err)
	}

	if err := binary.Read(file, binary.BigEndian, &field.AttributesCount); err != nil {
		return fmt.Errorf("reading attributes count: %w", err)
	}

	field.Attributes = make([]Attribute, field.AttributesCount)
	for i := range field.Attributes {
		if err := readAttribute(file, &field.Attributes[i]); err != nil {
			return fmt.Errorf("reading attribute %d: %w", i, err)
		}
	}

	return nil
}

func (f Field) String() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "Access Flags: 0x%04X\n", f.AccessFlags)
	fmt.Fprintf(&builder, "Name Index: %d\n", f.NameIndex)
	fmt.Fprintf(&builder, "Descriptor Index: %d\n", f.DescriptorIndex)
	fmt.Fprintf(&builder, "Attributes Count: %d\n", f.AttributesCount)
	for i, attribute := range f.Attributes {
		fmt.Fprintf(&builder, "Attribute #%d: %s\n", i+1, attribute)
	}
	return builder.String()
}
