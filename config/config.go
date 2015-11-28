package config

import (
	"fmt"
	"reflect"
	"strconv"
)

type defaultError struct {
	err error
}

type defaults map[string]reflect.Value

func (de defaultError) Error() string {
	return fmt.Sprintf("Load default error: %s", de.Error())
}

func Create(c interface{}) error {
	ds, err := getDefaults(reflect.TypeOf(c))
	if err != nil {
		return fmt.Errorf("get Defaults error: %s", err.Error())
	}
	return ds.set(reflect.ValueOf(c))
}

func getDefaults(t reflect.Type) (defaults, error) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("Non struct type")
	}
	ds := defaults{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		dv, err := getDefaultValue(field)
		if err != nil {
			return ds, err
		}
		ds[field.Name] = dv
	}
	return ds, nil
}

func (ds defaults) set(v reflect.Value) error {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("Non struct value")
	}
	for key, value := range ds {
		v.FieldByName(key).Set(value)
	}
	return nil
}

func getDefaultValue(field reflect.StructField) (reflect.Value, error) {
	d := field.Tag.Get("default")
	v := reflect.New(field.Type).Elem()
	switch field.Type.Kind() {
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Int:
		i, err := strconv.ParseInt(d, 0, 64)
		if err != nil {
			return v, defaultError{err}
		}
		v.SetInt(i)
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.Uint:
		i, err := strconv.ParseUint(d, 0, 64)
		if err != nil {
			return v, defaultError{err}
		}
		v.SetUint(i)
	case reflect.Float64:
	case reflect.String:
		v.SetString(d)
	}
	return v, nil
}
