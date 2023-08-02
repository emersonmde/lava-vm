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
	//return &ExecutionEngine{class: class}
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
		// Handle other instructions...
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
		// Assuming you have a method to get method's name
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
