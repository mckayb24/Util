package config

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

type defaultError struct {
	err error
}

type flags map[string]reflect.Value

func (de defaultError) Error() string {
	return fmt.Sprintf("Load default error: %s", de.Error())
}

func Parse(c interface{}) error {
	fs, err := getFlags(reflect.TypeOf(c))
	if err != nil {
		return fmt.Errorf("get Defaults error: %s", err.Error())
	}
	flag.Parse()
	return fs.set(reflect.ValueOf(c))
}

func getFlags(t reflect.Type) (flags, error) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("Non struct type")
	}
	fs := flags{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fv, err := getFlagValue(field)
		if err != nil {
			return fs, err
		}
		fs[field.Name] = fv
	}
	return fs, nil
}

func (fs flags) set(v reflect.Value) error {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("Non struct value")
	}
	for key, fv := range fs {
		f := v.FieldByName(key)
		if f.Kind() == reflect.Ptr {
			f.Set(fv)
			continue
		}
		f.Set(fv.Elem())
	}
	return nil
}

func getFlagValue(field reflect.StructField) (reflect.Value, error) {
	d := field.Tag.Get("default")
	usage := field.Tag.Get("usage")
	switch field.Type.Kind() {
	case reflect.Int64:
		i, err := strconv.ParseInt(d, 0, 64)
		if err != nil {
			return reflect.Value{}, defaultError{err}
		}
		return reflect.ValueOf(flag.Int64(field.Name, i, usage)), nil
	case reflect.Int:
		i, err := strconv.Atoi(d)
		if err != nil {
			return reflect.Value{}, defaultError{err}
		}
		return reflect.ValueOf(flag.Int(field.Name, i, usage)), nil
	case reflect.Uint64:
		i, err := strconv.ParseUint(d, 0, 64)
		if err != nil {
			return reflect.Value{}, defaultError{err}
		}
		return reflect.ValueOf(flag.Uint64(field.Name, i, usage)), nil
	case reflect.Uint:
		i, err := strconv.ParseUint(d, 0, 64)
		if err != nil {
			return reflect.Value{}, defaultError{err}
		}
		return reflect.ValueOf(flag.Uint(field.Name, uint(i), usage)), nil
	case reflect.Float64:
		f, err := strconv.ParseFloat(d, 64)
		if err != nil {
			return reflect.Value{}, defaultError{err}
		}
		return reflect.ValueOf(flag.Float64(field.Name, f, usage)), nil
	case reflect.String:
		return reflect.ValueOf(flag.String(field.Name, d, usage)), nil
	case reflect.Bool:
		b, err := strconv.ParseBool(d)
		if err != nil {
			return reflect.Value{}, defaultError{err}
		}
		return reflect.ValueOf(flag.Bool(field.Name, b, usage)), nil
	}
	return reflect.Value{}, fmt.Errorf("missing kind %s", field.Type.Kind())
}
