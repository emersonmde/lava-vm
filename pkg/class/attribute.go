package class

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

type Attribute struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	Info               []byte
}

// Read an Attribute from the given file
func readAttribute(file *os.File, attribute *Attribute) error {
	if err := binary.Read(file, binary.BigEndian, &attribute.AttributeNameIndex); err != nil {
		return fmt.Errorf("reading attribute name index: %w", err)
	}

	if err := binary.Read(file, binary.BigEndian, &attribute.AttributeLength); err != nil {
		return fmt.Errorf("reading attribute length: %w", err)
	}

	attribute.Info = make([]byte, attribute.AttributeLength)
	if err := binary.Read(file, binary.BigEndian, &attribute.Info); err != nil {
		return fmt.Errorf("reading attribute info: %w", err)
	}

	return nil
}

func (a Attribute) String() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "Attribute Name Index: %d\n", a.AttributeNameIndex)
	fmt.Fprintf(&builder, "Attribute Length: %d\n", a.AttributeLength)
	fmt.Fprintf(&builder, "Info: %x\n", a.Info)
	return builder.String()
}
