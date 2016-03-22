package ikea

import "fmt"

// Instructions is an interface type that is returned when NewInstructions is
// called.
type Instructions interface {
	AddInstruction(string, func() interface{}) error
	GetInstruction(string) (func() interface{}, error)
}

type instructions struct {
	store map[string]func() interface{}
}

// NewInstructions initializes a new instruction type with a default value for
// the store and returns an Instructions interface.
func NewInstructions() Instructions {
	return &instructions{make(map[string]func() interface{})}
}

// AddInstruction takes an instruction string and the function that is used to
// store in map. If the instruction already exists it returns an error.
func (i *instructions) AddInstruction(instruction string, f func() interface{}) error {
	_, ok := i.store[instruction]
	if ok {
		return fmt.Errorf("The instruction: %v already exists.", instruction)
	}
	i.store[instruction] = f
	return nil
}

// GetInstruction takes an instruction string and looks up the function that
// is represents the value for the instruction string. If the instruction
// string does not exist in the map it returns an error.
func (i *instructions) GetInstruction(instruction string) (func() interface{}, error) {
	f, ok := i.store[instruction]
	if ok {
		return f, nil
	}
	return nil, fmt.Errorf("The instruction: %v does not exist!", instruction)
}
