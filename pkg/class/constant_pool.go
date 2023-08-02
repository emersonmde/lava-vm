package class

import (
	"encoding/binary"
	"fmt"
	"os"
	"unicode/utf8"
)

type ConstantPool struct {
	entries []ConstantPoolEntry
}

func (cp *ConstantPool) Get(index uint16) ConstantPoolEntry {
	return cp.entries[index-1]
}

func (cp *ConstantPool) GetConstantName(index uint16) string {
	if index == 0 || index >= uint16(len(cp.entries)) {
		return ""
	}
	entry := cp.entries[index]
	if utf8Entry, ok := entry.Value.(*ConstantUtf8Value); ok {
		return utf8Entry.String()
	}
	return ""
}

type ConstantPoolEntry struct {
	Tag   uint8             // The tag representing the type of constant pool entry
	Value ConstantPoolValue // Data specific to the constant pool entry type
}

type ConstantPoolValue interface{}

// ConstantUtf8Value represents a UTF-8 string in a Java class file.
type ConstantUtf8Value struct {
	Length uint16
	Bytes  []byte
}

// String returns the ConstantUtf8Value as a string.
func (c *ConstantUtf8Value) String() string {
	return string(c.Bytes)
}

// readConstantUtf8Value reads a ConstantUtf8Value from the provided file.
// It returns the ConstantPoolValue and any error encountered.
// readConstantUtf8Value reads a ConstantUtf8Value from the provided file.
// It returns the ConstantPoolValue and any error encountered.
func readConstantUtf8Value(file *os.File) (ConstantPoolValue, error) {
	value := ConstantUtf8Value{}
	if err := binary.Read(file, binary.BigEndian, &value.Length); err != nil {
		return nil, err
	}
	value.Bytes = make([]byte, value.Length)
	if err := binary.Read(file, binary.BigEndian, &value.Bytes); err != nil {
		return nil, err
	}
	if !utf8.Valid(value.Bytes) {
		return nil, fmt.Errorf("invalid UTF-8 sequence")
	}
	return &value, nil
}

// ConstantIntegerValue represents an integer constant in a Java class file.
type ConstantIntegerValue struct {
	Value int32
}

// readConstantIntegerValue reads a ConstantIntegerValue from the provided file.
// It returns the ConstantPoolValue and any error encountered.
func readConstantIntegerValue(file *os.File) (ConstantPoolValue, error) {
	value := ConstantIntegerValue{}
	if err := binary.Read(file, binary.BigEndian, &value.Value); err != nil {
		return nil, err
	}
	return &value, nil
}

// ConstantFloatValue represents a floating point constant in a Java class file.
type ConstantFloatValue struct {
	Value float32
}

// readConstantFloatValue reads a ConstantFloatValue from the provided file.
// It returns the ConstantPoolValue and any error encountered.
func readConstantFloatValue(file *os.File) (ConstantPoolValue, error) {
	value := ConstantFloatValue{}
	if err := binary.Read(file, binary.BigEndian, &value.Value); err != nil {
		return nil, err
	}
	return &value, nil
}

// ConstantLongValue represents a long integer constant in a Java class file.
type ConstantLongValue struct {
	Value int64
}

// readConstantLongValue reads a ConstantLongValue from the provided file.
// It returns the ConstantPoolValue and any error encountered.
func readConstantLongValue(file *os.File) (ConstantPoolValue, error) {
	value := ConstantLongValue{}
	if err := binary.Read(file, binary.BigEndian, &value.Value); err != nil {
		return nil, err
	}
	return &value, nil
}

// ConstantDoubleValue represents a double precision floating point constant in a Java class file.
type ConstantDoubleValue struct {
	Value float64
}

// readConstantDoubleValue reads a ConstantDoubleValue from the provided file.
// It returns the ConstantPoolValue and any error encountered.
func readConstantDoubleValue(file *os.File) (ConstantPoolValue, error) {
	value := ConstantDoubleValue{}
	if err := binary.Read(file, binary.BigEndian, &value.Value); err != nil {
		return nil, err
	}
	return &value, nil
}

// ConstantClassRefValue represents a Class reference in Java class files.
// It contains an index into the constant pool table.
type ConstantClassRefValue struct {
	Index uint16
}

// readConstantClassRefValue reads a ConstantClassRefValue from the provided file.
// It returns the ConstantPoolValue and any error encountered.
func readConstantClassRefValue(file *os.File) (ConstantPoolValue, error) {
	value := ConstantClassRefValue{}
	if err := binary.Read(file, binary.BigEndian, &value.Index); err != nil {
		return nil, err
	}
	return &value, nil
}

// ConstantStringRefValue represents a String reference in Java class files.
// It contains an index into the constant pool table.
type ConstantStringRefValue struct {
	Index uint16
}

// readConstantStringRefValue reads a ConstantStringRefValue from the provided file.
// It returns the ConstantPoolValue and any error encountered.
func readConstantStringRefValue(file *os.File) (ConstantPoolValue, error) {
	value := ConstantStringRefValue{}
	if err := binary.Read(file, binary.BigEndian, &value.Index); err != nil {
		return nil, err
	}
	return &value, nil
}

// ConstantFieldRefValue represents a field reference in Java class files.
// It contains two indexes into the constant pool table.
// The first index points to a Class reference entry and
// the second index points to a NameAndType descriptor entry.
type ConstantFieldRefValue struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

// readConstantFieldRefValue reads a ConstantFieldRefValue from the provided file.
// It returns the ConstantPoolValue and any error encountered.
func readConstantFieldRefValue(file *os.File) (ConstantPoolValue, error) {
	value := ConstantFieldRefValue{}
	if err := binary.Read(file, binary.BigEndian, &value.ClassIndex); err != nil {
		return nil, err
	}
	if err := binary.Read(file, binary.BigEndian, &value.NameAndTypeIndex); err != nil {
		return nil, err
	}
	return &value, nil
}

// ConstantMethodRefValue represents a method reference in Java class files.
// It contains two indexes into the constant pool table.
// The first index points to a Class reference entry and
// the second index points to a NameAndType descriptor entry.
type ConstantMethodRefValue struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

// readConstantMethodRefValue reads a ConstantMethodRefValue from the provided file.
// It returns the ConstantPoolValue and any error encountered.
func readConstantMethodRefValue(file *os.File) (ConstantPoolValue, error) {
	value := ConstantMethodRefValue{}
	if err := binary.Read(file, binary.BigEndian, &value.ClassIndex); err != nil {
		return nil, err
	}
	if err := binary.Read(file, binary.BigEndian, &value.NameAndTypeIndex); err != nil {
		return nil, err
	}
	return &value, nil
}

// ConstantInterfaceMethodRefValue represents an interface method reference in Java class files.
// It contains two indexes into the constant pool table.
// The first index points to a Class reference entry and
// the second index points to a NameAndType descriptor entry.
type ConstantInterfaceMethodRefValue struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

// readConstantInterfaceMethodRefValue reads a ConstantInterfaceMethodRefValue from the provided file.
// It returns the ConstantPoolValue and any error encountered.
func readConstantInterfaceMethodRefValue(file *os.File) (ConstantPoolValue, error) {
	value := ConstantInterfaceMethodRefValue{}
	if err := binary.Read(file, binary.BigEndian, &value.ClassIndex); err != nil {
		return nil, err
	}
	if err := binary.Read(file, binary.BigEndian, &value.NameAndTypeIndex); err != nil {
		return nil, err
	}
	return &value, nil
}

// ConstantNameAndTypeDescriptorValue represents a Name and Type descriptor in Java class files.
// It contains two indexes into the constant pool table.
// The first index points to a UTF-8 string entry that represents a name,
// and the second index points to a UTF-8 string entry that represents a type descriptor.
type ConstantNameAndTypeDescriptorValue struct {
	NameIndex       uint16
	DescriptorIndex uint16
}

// readConstantNameAndTypeDescriptorValue reads a ConstantNameAndTypeDescriptorValue from the provided file.
// It returns the ConstantPoolValue and any error encountered.
func readConstantNameAndTypeDescriptorValue(file *os.File) (ConstantPoolValue, error) {
	value := ConstantNameAndTypeDescriptorValue{}
	if err := binary.Read(file, binary.BigEndian, &value.NameIndex); err != nil {
		return nil, err
	}
	if err := binary.Read(file, binary.BigEndian, &value.DescriptorIndex); err != nil {
		return nil, err
	}
	return &value, nil
}

type valueReader func(file *os.File) (ConstantPoolValue, error)

var valueReaders = map[uint8]valueReader{
	1:  readConstantUtf8Value,
	3:  readConstantIntegerValue,
	4:  readConstantFloatValue,
	5:  readConstantLongValue,
	6:  readConstantDoubleValue,
	7:  readConstantClassRefValue,
	8:  readConstantStringRefValue,
	9:  readConstantFieldRefValue,
	10: readConstantMethodRefValue,
	11: readConstantInterfaceMethodRefValue,
	12: readConstantNameAndTypeDescriptorValue,
}

func readConstantPool(file *os.File, class *Class) (err error) {
	if err := binary.Read(file, binary.BigEndian, &class.ConstantPoolCount); err != nil {
		return err
	}

	class.ConstantPool.entries = make([]ConstantPoolEntry, class.ConstantPoolCount)

	for i := uint16(1); i < class.ConstantPoolCount; i++ {
		var tag uint8
		if err := binary.Read(file, binary.BigEndian, &tag); err != nil {
			return err
		}
		fmt.Println("Tag :", tag)

		reader, exists := valueReaders[tag]
		if !exists {
			//continue
			return fmt.Errorf("unknown tag %d", tag)
		}

		value, err := reader(file)
		if err != nil {
			return err
		}

		class.ConstantPool.entries[i] = ConstantPoolEntry{Tag: tag, Value: value}

		if tag == 5 || tag == 6 {
			i++
			if i < class.ConstantPoolCount {
				class.ConstantPool.entries[i] = ConstantPoolEntry{}
			}
		}
	}
	return nil
}
