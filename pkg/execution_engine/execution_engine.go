package execution_engine

import (
	"errors"
	"fmt"
	"lava-vm/pkg/class"
)

type Class = class.Class
type Method = class.Method

type ExecutionEngine struct {
	class *Class
	heap  *Heap
	stack *OperandStack
}

type Heap struct {
	objects []Object
}

type Object struct {
	id uint32
}

type OperandStack struct {
	operands []interface{}
}

// Push adds an element to the top of the stack
func (os *OperandStack) Push(value interface{}) {
	os.operands = append(os.operands, value)
}

// Pop removes and returns the top element of the stack, or nil if the stack is empty
func (os *OperandStack) Pop() interface{} {
	if len(os.operands) == 0 {
		return nil
	}

	val := os.operands[len(os.operands)-1]
	os.operands = os.operands[:len(os.operands)-1]

	return val
}

func NewExectuionEngine(class *Class) *ExecutionEngine {
	return &ExecutionEngine{
		class: class,
		heap:  &Heap{},
		stack: &OperandStack{},
	}
}

func (e *ExecutionEngine) Execute() error {
	method, err := e.getMainMethod()
	if err != nil {
		return err
	}

	fmt.Printf("Found main method \n%s\n", method.String())
	code, err := method.GetCode()

	if err != nil {
		return err
	}

	fmt.Println("Found code")
	instructions, err := ParseInstructions(code)
	if err != nil {
		fmt.Println("Error", err)
		return err
	}
	fmt.Println("Parsed Instructions", instructions)
	for _, instruction := range instructions {
		switch instruction.Opcode {
		case 0xbb: // "new" opcode
			objectRef := e.allocateObject()
			e.stack.Push(objectRef)
			fmt.Printf("Created new object with reference: %v\n", objectRef)
		case 0x60: // "iadd" opcode
			err := e.iadd()
			if err != nil {
				return err
			}
		case 0x64: // "isub" opcode
			err := e.isub()
			if err != nil {
				return err
			}
		case 0x68: // "imul" opcode
			err := e.imul()
			if err != nil {
				return err
			}
		case 0x6c: // "idiv" opcode
			err := e.idiv()
			if err != nil {
				return err
			}
		default:
			fmt.Printf("Unhandled instruction: %02x\n", instruction.Opcode)
		}
	}

	return nil
}

func (e *ExecutionEngine) allocateObject() uint32 {
	// For now, just assign a unique ID as a reference
	objectRef := uint32(len(e.heap.objects) + 1)
	object := Object{id: objectRef}
	e.heap.objects = append(e.heap.objects, object)
	fmt.Println("Allocated object ", object)
	return objectRef
}

func (e *ExecutionEngine) getMainMethod() (Method, error) {
	for _, method := range e.class.Methods {
		name, err := e.class.GetConstantName(method.NameIndex)
		if err != nil {
			return Method{}, err
		}

		if name == "main" {
			fmt.Printf("Found main file %+v\n", method)
			return method, nil
		}
	}

	return Method{}, errors.New("main method not found in the class")
}

func (e *ExecutionEngine) iadd() error {
	op1Int, op2Int, err := e.popBinaryOpInt()
	if err != nil {
		return err
	}
	sum := op1Int + op2Int
	e.stack.Push(sum)
	return nil
}

func (e *ExecutionEngine) isub() error {
	op1Int, op2Int, err := e.popBinaryOpInt()
	if err != nil {
		return err
	}
	diff := op1Int - op2Int
	e.stack.Push(diff)
	return nil
}

func (e *ExecutionEngine) imul() error {
	op1Int, op2Int, err := e.popBinaryOpInt()
	if err != nil {
		return err
	}
	product := op1Int * op2Int
	e.stack.Push(product)
	return nil
}

func (e *ExecutionEngine) idiv() error {
	op1Int, op2Int, err := e.popBinaryOpInt()
	if err != nil {
		return err
	}

	if op2Int == 0 {
		return errors.New("divide by zero")
	}

	quotient := op1Int / op2Int
	e.stack.Push(quotient)
	return nil
}

func (e *ExecutionEngine) popBinaryOpInt() (int32, int32, error) {
	// Pop the top two elements from the stack
	op1 := e.stack.Pop()
	op2 := e.stack.Pop()

	// Check that they are both integers
	op1Int, ok1 := op1.(int32)
	op2Int, ok2 := op2.(int32)
	if !ok1 || !ok2 {
		return 0, 0, errors.New("operand not an integer")
	}
	return op1Int, op2Int, nil
}
