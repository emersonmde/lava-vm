package class_parser

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Class struct {
	Magic             uint32
	MinorVersion      uint16
	MajorVersion      uint16
	ConstantPoolCount uint16
	ConstantPool      []ConstantPoolEntry
}

// Add other types here

func Parse(filename string) (*Class, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	class := &Class{}
	if err := readVersionInfo(file, class); err != nil {
		return nil, fmt.Errorf("reading version info: %w", err)
	}

	if err := readConstantPool(file, class); err != nil {
		return nil, fmt.Errorf("reading constant pool: %w", err)
	}

	// Parse other parts of the class file here

	return class, nil
}

func readVersionInfo(file *os.File, class *Class) (err error) {
	if err = binary.Read(file, binary.BigEndian, &class.Magic); err != nil {
		return
	}

	if class.Magic != 0xCAFEBABE {
		return
	}

	if err = binary.Read(file, binary.BigEndian, &class.MinorVersion); err != nil {
		return
	}

	// 49 = Java SE 5.0 through 65 = Java SE 21
	if err = binary.Read(file, binary.BigEndian, &class.MajorVersion); err != nil {
		return
	}

	return nil
}

func (c *Class) Display() {
	fmt.Printf("Magic: 0x%X\n", c.Magic)
	fmt.Println("Minor Version:", c.MinorVersion)
	fmt.Println("Major Version:", c.MajorVersion)
	fmt.Println("Constant Pool Count:", c.ConstantPoolCount)

	for i := uint16(1); i < c.ConstantPoolCount; i++ {

		fmt.Printf("\nConstant #%d\n", i)
		fmt.Println("Tag:", c.ConstantPool[i].Tag)

		switch v := c.ConstantPool[i].Value.(type) {
		case ConstantUtf8Value:
			fmt.Printf("Type: %T", v)
			fmt.Printf("Value: %+v\n", v.String())
		default:
			fmt.Printf("Type: %T", v)
			fmt.Println("Value:", v)
		}
	}
}
