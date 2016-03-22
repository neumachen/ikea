[Build Status]: https://travis-ci.org/magicalbanana/ikea
[Build Status Badge]: https://travis-ci.org/magicalbanana/ikea.svg?branch=master
[Coverage Report]: https://coveralls.io/github/magicalbanana/ikea?branch=master
[Coverage Report Badge]: https://coveralls.io/repos/github/magicalbanana/ikea/badge.svg?branch=master
[Go Report Card]: https://goreportcard.com/report/github.com/magicalbanana/ikea
[Go Report Card Badge]: https://goreportcard.com/badge/github.com/magicalbanana/ikea
[GoDoc]: http://godoc.org/github.com/magicalbanana/ikea
[GoDoc Badge]: https://godoc.org/github.com/valyala/quicktemplate?status.svg
[License]: https://raw.githubusercontent.com/magicalbanana/ikea/master/LICENSE
[License Badge]: https://img.shields.io/badge/license-MIT-orange.svg?style=flat-square

[golang]: https://golang.org/
[golang struct type]: https://gobyexample.com/structs
[fake]: https://github.com/icrowley/fake
[gofakeit]: https://github.com/brianvoe/gofakeit

[![Build Status][Build Status Badge]][Build Status]
[![Coverage Status][Coverage Report Badge]][Coverage Report]
[![Go Report Card][Go Report Card Badge]][Go Report Card]
[![GoDoc][GoDoc Badge]][GoDoc]
[![License][License Badge]][License]

[instructions]: https://godoc.org/github.com/magicalbanana/ikea#Instructions

# Abstract

Ikea is a library intended to fake a [golang struct type] with fake data
that you will be using through the [instructions] that you will be setting.

# Disclaimer

Package ikea is currently in prototype stage. The API may change or some
internal mechanisms maybe changed that changes the expected outcome.

# Author's POV

The motivation behind this is when the author was making an application in
[golang] and needed to extensively test the data stores (databases) with
inserts, updates etc, the author wanted an easy way of doing it without having
the need to use constructors. Hence, Ikea was made. The name was taken because
two weekends before making this package the author went to Ikea with the
author's significant other to buy some furniture there. He also realized that
how challenging assembly of a furniture is - on the author's relationship. So
the author dedicates this package to the all mighty IKEA.

# Usage

## Instructions

You first need to load some instructions so that when ikea assembles your
struct (that's right, ikea assembles your struct using instructions) it will
have a function to call to generate the necessary data that will be used to
set the struct field.

### Example:

You can use existing faker libraries such as [fake] or [gofakeit]

```go
// Custom instruction
func MakeMalm() interface{} {
    return &Foo
}

var testIns = map[string]func() interface{}{
	"character":  func() interface{} { return fake.Character() },
	"number":     func() interface{} { return gofakeit.Number(1, 1000) },
	"time":       func() interface{} { return gofakeit.Date() },
        "malm":       makeMalm,
}

i := ikea.NewInstructions()
for k, v := range testIns {
	if err := i.AddInstruction(k, v); err != nil {
		panic(err)
	}
}
```

## Assemble

To assemble a struct with the desired data or instruction you want to use, you
need to make sure that your struct field has the `ikea` tag and the
corresponding instruction you will use exists in the Instructions store.

### Example

```go
var testIns = map[string]func() interface{}{
	"skorva":  func() interface{} { return "skorva" },
}

type Malm struct {
    BedBase string `ikea:"skorva"`
}

i := ikea.NewInstructions()

for k, v := range testIns {
	if err := i.AddInstruction(k, v); err != nil {
		panic(err)
	}
}

err := ikea.Assemble(&malm, i)
```

## TODO

- Allow setting of struct types
  - e.g:
    ```go
      type Moo struct {
        Zoo string
      }

      type Foo struct {
        MooMoo Moo `ikea:"moo"`
      }
    ```
- Allow setting a value to skip fields/structs that are not nil
