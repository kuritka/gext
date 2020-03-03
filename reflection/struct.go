// Package reflection provides reflection helpers over structures.
package reflection

import (
	"errors"
	"fmt"
	"reflect"
)

// StructSliceAppend appends v to the structure's field by the given field name.
// Returns an error if is not a pointer to struct or v has a different type. It panics
// if v mismatches the type of a slice.
func StructSliceAppend(structure, v interface{}, field string) error {
	ptr := reflect.ValueOf(structure)
	if ptr.Kind() != reflect.Ptr {
		return errors.New("not settable")
	}
	val := ptr.Elem()
	if val.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}
	slice := val.FieldByName(field)
	empty := reflect.Value{}
	if slice == empty {
		return fmt.Errorf("%q filed is not found", field)
	}
	if slice.Kind() != reflect.Slice {
		return fmt.Errorf("%q filed is not a slice", field)
	}
	if !slice.CanSet() {
		return fmt.Errorf("%q filed is not settable", field)
	}
	slice.Set(reflect.Append(slice, reflect.ValueOf(v)))
	return nil
}

// StructSet docs.
func StructSet(structure, v interface{}, field string) error {
	ptr := reflect.ValueOf(structure)
	if ptr.Kind() != reflect.Ptr {
		return errors.New("not settable")
	}
	val := ptr.Elem()
	if val.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}
	f := val.FieldByName(field)
	empty := reflect.Value{}
	if f == empty {
		return fmt.Errorf("%q field is not found", field)
	}
	vv := reflect.ValueOf(v)
	if vv.Kind() != f.Kind() {
		return fmt.Errorf("invalid type %v for %q", vv.Kind(), field)
	}
	if !f.CanSet() {
		return fmt.Errorf("%q field is not settable", field)
	}
	f.Set(vv)
	return nil
}

// SetStructValues put into structure data from array
// Example: structure.Field = data[index]
func SetStructValues(structure interface{}, data []string) {
	s := reflect.ValueOf(structure).Elem()
	loopLen := s.NumField()
	if len(data) < loopLen {
		loopLen = len(data)
	}
	for i := 0; i < loopLen; i++ {
		s.Field(i).SetString(data[i])
	}
}

// GetStructNumField return structure NumField
func GetStructNumField(structure interface{}) int {
	return reflect.ValueOf(structure).Elem().NumField()
}
