package ikea

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

type malm struct {
	Luroy string `ikea:"characters"`
}

// MalmSet is the custom type we are using for testing setting a custom type.
type MalmSet struct {
	UnderBedStorage string `ikea:"malm"`
}

var testIns = map[string]func() interface{}{
	"character":  func() interface{} { return fake.Character() },
	"characters": func() interface{} { return fake.Characters() },
	"bool":       func() interface{} { return gofakeit.Bool() },
	"number":     func() interface{} { return gofakeit.Number(1, 1000) },
	"time":       func() interface{} { return gofakeit.Date() },
}

type testAssertion func(*testing.T, string)

func TestAssemble(t *testing.T) {
	i := NewInstructions()
	for k, v := range testIns {
		if err := i.AddInstruction(k, v); err != nil {
			panic(err)
		}
	}

	tests := []struct {
		desc      string
		assertion testAssertion
	}{
		{
			desc: "passes a non pointer",
			assertion: func(t *testing.T, m string) {
				assert.Error(t, Assemble(malm{}, i), m)
			},
		},
		{
			desc: "passes a non struct",
			assertion: func(t *testing.T, m string) {
				c := malm{}
				assert.Error(t, Assemble(&c.Luroy, i), m)
			},
		},
		{
			desc: "tag specified has no faker func",
			assertion: func(t *testing.T, m string) {
				malm := struct {
					Luroy string `ikea:"skorva"`
				}{}
				assert.Error(t, Assemble(&malm, i), m)
			},
		},
		{
			desc: "skips the value for the non tagged fields",
			assertion: func(t *testing.T, m string) {
				malm := struct {
					Luroy   string `ikea:"character"`
					Loenset string
				}{}
				assert.NoError(t, Assemble(&malm, i), m)
				assert.NotEqual(t, malm.Luroy, "")
			},
		},
		{
			desc: "value for string type is of string type",
			assertion: func(t *testing.T, m string) {
				malm := struct {
					Luroy string `ikea:"characters"`
				}{}
				assert.NoError(t, Assemble(&malm, i), m)
				assert.NotEqual(t, malm.Luroy, "")
			},
		},
		{
			desc: "value for string type is not of string type",
			assertion: func(t *testing.T, m string) {
				malm := struct {
					Luroy string `ikea:"number"`
				}{}
				assert.Error(t, Assemble(&malm, i), m)
			},
		},
		{
			desc: "value for bool type is of bool type",
			assertion: func(t *testing.T, m string) {
				malm := struct {
					Luroy bool `ikea:"bool"`
				}{}
				assert.NoError(t, Assemble(&malm, i), m)
			},
		},
		{
			desc: "value for bool type is not of bool type",
			assertion: func(t *testing.T, m string) {
				malm := struct {
					Luroy   bool `ikea:"character"`
					Loenset bool `ikea:"number"`
				}{}
				assert.Error(t, Assemble(&malm, i), m)
			},
		},
		{
			desc: "value for bool type is of bool type",
			assertion: func(t *testing.T, m string) {
				malm := struct {
					Luroy bool `ikea:"bool"`
				}{}
				assert.NoError(t, Assemble(&malm, i), m)
			},
		},
		{
			desc: "value for bool type is not of bool type",
			assertion: func(t *testing.T, m string) {
				malm := struct {
					Luroy bool `ikea:"character"`
				}{}
				assert.Error(t, Assemble(&malm, i), m)
			},
		},
		{
			desc: "value for int type is of int type",
			assertion: func(t *testing.T, m string) {
				malm := struct {
					Luroy int `ikea:"number"`
				}{}
				assert.NoError(t, Assemble(&malm, i), m)
			},
		},
		{
			desc: "value for int type is not of int type",
			assertion: func(t *testing.T, m string) {
				malm := struct {
					Luroy int `ikea:"character"`
				}{}
				assert.Error(t, Assemble(&malm, i), m)
			},
		},
		{
			desc: "value for time type is of time type",
			assertion: func(t *testing.T, m string) {
				malm := struct {
					Luroy time.Time `ikea:"time"`
				}{}
				assert.NoError(t, Assemble(&malm, i), m)
			},
		},
		{
			desc: "value for time type is not of time type",
			assertion: func(t *testing.T, m string) {
				malm := struct {
					Luroy   time.Time `ikea:"character"`
					Leonset string    `ikea:"time"`
				}{}
				assert.Error(t, Assemble(&malm, i), m)
			},
		},
		{
			desc: "unsupported type",
			assertion: func(t *testing.T, m string) {
				malm := struct {
					Luroy rune `ikea:"character"`
				}{}
				assert.Error(t, Assemble(&malm, i), m)
			},
		},
	}

	for _, test := range tests {
		test.assertion(t, test.desc)
	}
}
