package ikea

import (
	"errors"
	"fmt"
	"reflect"
)

// Assemble takes a val interface and an Instruction interface. If the
// val is not  a pointer and a struct it returns an error.
func Assemble(val interface{}, ins Instructions) error {
	ptrRef := reflect.ValueOf(val)
	if ptrRef.Kind() != reflect.Ptr {
		return errors.New("Expected a pointer to a Struct")
	}
	ref := ptrRef.Elem()
	if ref.Kind() != reflect.Struct {
		return errors.New("Expected a struct type to be passed")
	}
	return assembleStruct(ref, val, ins)
}

func assembleStruct(ref reflect.Value, val interface{}, inst Instructions) error {
	refType := ref.Type()
	for i := 0; i < refType.NumField(); i++ {
		tag := getTag(refType.Field(i))
		if tag == "" {
			continue
		}
		ins, err := inst.GetInstruction(tag)
		if err != nil {
			return err
		}
		// TODO: add logic that will skip fields that have values
		// already set
		// TODO: add logic that will determine to set a StructField or
		// a Struct
		if err := setField(ref.Field(i), refType.Field(i), ins()); err != nil {
			return err
		}
	}
	return nil
}

func getTag(field reflect.StructField) string {
	tag := field.Tag.Get("ikea")
	if tag != "" {
		return tag
	}
	return ""
}

func setField(field reflect.Value, refType reflect.StructField, value interface{}) error {
	vt := reflect.TypeOf(value)
	if field.Kind() != vt.Kind() {
		return setterError(field.Kind(), value)
	}

	field.Set(reflect.ValueOf(value))
	return nil
}

func setterError(k reflect.Kind, value interface{}) error {
	return fmt.Errorf(
		"Value not of %v type. Value: %v, Type: %T", k, value, value,
	)
}
