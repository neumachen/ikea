package ikea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddInstruction(t *testing.T) {
	tests := []struct {
		desc      string
		assertion testAssertion
	}{
		{
			desc: "adding instruction successfully",
			assertion: func(t *testing.T, desc string) {
				ikeaIns := NewInstructions()
				for k, v := range testIns {
					assert.NoError(t, ikeaIns.AddInstruction(k, v))
				}
			},
		},
		{
			desc: "instruction already exists",
			assertion: func(t *testing.T, desc string) {
				ikeaIns := NewInstructions()
				for k, v := range testIns {
					assert.NoError(t, ikeaIns.AddInstruction(k, v))
				}
				for k, v := range testIns {
					assert.Error(t, ikeaIns.AddInstruction(k, v))
				}
			},
		},
	}
	for _, test := range tests {
		test.assertion(t, test.desc)
	}

}
