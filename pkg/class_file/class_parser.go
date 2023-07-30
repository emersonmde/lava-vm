package class_file

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

type Class struct {
	Magic             uint32
	MinorVersion      uint16
	MajorVersion      uint16
	ConstantPoolCount uint16
	ConstantPool      []ConstantPoolEntry
	AccessFlags       uint16
	ThisClass         uint16
	SuperClass        uint16
	// TODO: finish below attributes
	InterfacesCount uint16
	Interfaces      []uint16
	FieldsCount     uint16
	Fields          []Field
	MethodsCount    uint16
	Methods         []Method
	AttributesCount uint16
	Attributes      []Attribute
}

func Parse(filename string) (*Class, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	class := &Class{}
	if err = binary.Read(file, binary.BigEndian, &class.Magic); err != nil {
		return nil, fmt.Errorf("reading magic number: %w", err)
	}

	if class.Magic != 0xCAFEBABE {
		return nil, fmt.Errorf("magic number check: %w", err)
	}

	if err = binary.Read(file, binary.BigEndian, &class.MinorVersion); err != nil {
		return nil, fmt.Errorf("reading minor version: %w", err)
	}

	// 49 = Java SE 5.0 through 65 = Java SE 21
	if err = binary.Read(file, binary.BigEndian, &class.MajorVersion); err != nil {
		return nil, fmt.Errorf("reading major version: %w", err)
	}

	if err := readConstantPool(file, class); err != nil {
		return nil, fmt.Errorf("reading constant pool: %w", err)
	}

	if err = binary.Read(file, binary.BigEndian, &class.AccessFlags); err != nil {
		return nil, fmt.Errorf("reading access flags: %w", err)
	}

	if err = binary.Read(file, binary.BigEndian, &class.ThisClass); err != nil {
		return nil, fmt.Errorf("reading this class: %w", err)
	}

	if err = binary.Read(file, binary.BigEndian, &class.SuperClass); err != nil {
		return nil, fmt.Errorf("reading super class: %w", err)
	}

	// Continue from the "Parse other parts of the class file here" comment

	if err = binary.Read(file, binary.BigEndian, &class.InterfacesCount); err != nil {
		return nil, fmt.Errorf("reading interfaces count: %w", err)
	}

	class.Interfaces = make([]uint16, class.InterfacesCount)
	if err = binary.Read(file, binary.BigEndian, &class.Interfaces); err != nil {
		return nil, fmt.Errorf("reading interfaces: %w", err)
	}

	if err = binary.Read(file, binary.BigEndian, &class.FieldsCount); err != nil {
		return nil, fmt.Errorf("reading fields count: %w", err)
	}

	class.Fields = make([]Field, class.FieldsCount)
	for i := range class.Fields {
		if err = readField(file, &class.Fields[i]); err != nil {
			return nil, fmt.Errorf("reading field %d: %w", i, err)
		}
	}

	if err = binary.Read(file, binary.BigEndian, &class.MethodsCount); err != nil {
		return nil, fmt.Errorf("reading methods count: %w", err)
	}

	class.Methods = make([]Method, class.MethodsCount)
	for i := range class.Methods {
		if err = readMethod(file, &class.Methods[i]); err != nil {
			return nil, fmt.Errorf("reading method %d: %w", i, err)
		}
	}

	if err = binary.Read(file, binary.BigEndian, &class.AttributesCount); err != nil {
		return nil, fmt.Errorf("reading attributes count: %w", err)
	}

	class.Attributes = make([]Attribute, class.AttributesCount)
	for i := range class.Attributes {
		if err = readAttribute(file, &class.Attributes[i]); err != nil {
			return nil, fmt.Errorf("reading attribute %d: %w", i, err)
		}
	}

	return class, nil

	// Parse other parts of the class file here

	return class, nil
}

func (c Class) String() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "Magic: 0x%X\n", c.Magic)
	fmt.Fprintf(&builder, "Minor Version: %d\n", c.MinorVersion)
	fmt.Fprintf(&builder, "Major Version: %d\n", c.MajorVersion)
	fmt.Fprintf(&builder, "Constant Pool Count: %d\n", c.ConstantPoolCount)

	for i := uint16(1); i < c.ConstantPoolCount; i++ {
		fmt.Fprintf(&builder, "\n\tConstant #%d\n", i)
		fmt.Fprintf(&builder, "\tTag: %v\n", c.ConstantPool[i].Tag)

		switch v := c.ConstantPool[i].Value.(type) {
		case fmt.Stringer:
			fmt.Fprintf(&builder, "\tType: %T Value: %s\n", v, v.String())
		default:
			fmt.Fprintf(&builder, "\tType: %T Value: %v\n", v, v)
		}
	}

	fmt.Fprintf(&builder, "\nAccess Flags: 0x%04X\n", c.AccessFlags)
	fmt.Fprintf(&builder, "This Class: %d: %s\n", c.ThisClass, c.ConstantPool[c.ThisClass].Value)
	fmt.Fprintf(&builder, "Super Class: %d\n", c.SuperClass)

	fmt.Fprintf(&builder, "Interfaces Count: %d\n", c.InterfacesCount)
	for i, iface := range c.Interfaces {
		fmt.Fprintf(&builder, "Interface #%d: %v\n", i+1, iface)
	}

	fmt.Fprintf(&builder, "\nFields Count: %d\n\n", c.FieldsCount)
	for i, field := range c.Fields {
		fmt.Fprintf(&builder, "\tField #%d: %s\n", i+1, field)
	}

	fmt.Fprintf(&builder, "\nMethods Count: %d\n\n", c.MethodsCount)
	for i, method := range c.Methods {
		fmt.Fprintf(&builder, "\tMethod #%d: %s\n", i+1, method)
	}

	fmt.Fprintf(&builder, "\nAttributes Count: %d\n\n", c.AttributesCount)
	for i, attr := range c.Attributes {
		fmt.Fprintf(&builder, "\tAttribute #%d: %s\n", i+1, attr)
	}

	return builder.String()
}
