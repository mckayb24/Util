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

type flags map[string]interface{}

func (de defaultError) Error() string {
	return fmt.Sprintf("Load default error: %s", de.Error())
}

func Parse(c interface{}) error {
	fs, err := getFlags(reflect.TypeOf(c), "")
	if err != nil {
		return fmt.Errorf("get Defaults error: %s", err.Error())
	}
	flag.Parse()
	return fs.set(reflect.ValueOf(c))
}

func getFlags(t reflect.Type, parent string) (flags, error) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("Non struct type")
	}
	fs := flags{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fv, err := getFlagValue(field, parent)
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
		if fs, ok := fv.(flags); ok {
			fs.set(f)
			continue
		}
		val := reflect.ValueOf(fv)
		if f.Kind() == reflect.Ptr {
			f.Set(val)
			continue
		}
		f.Set(val.Elem())
	}
	return nil
}

func getFlagValue(field reflect.StructField, parent string) (interface{}, error) {
	d := field.Tag.Get("default")
	usage := field.Tag.Get("usage")
	name := getName(field, parent)
	switch field.Type.Kind() {
	case reflect.Int64:
		i, err := strconv.ParseInt(d, 0, 64)
		if err != nil {
			return reflect.Value{}, defaultError{err}
		}
		return flag.Int64(name, i, usage), nil
	case reflect.Int:
		i, err := strconv.Atoi(d)
		if err != nil {
			return reflect.Value{}, defaultError{err}
		}
		return flag.Int(name, i, usage), nil
	case reflect.Uint64:
		i, err := strconv.ParseUint(d, 0, 64)
		if err != nil {
			return reflect.Value{}, defaultError{err}
		}
		return flag.Uint64(name, i, usage), nil
	case reflect.Uint:
		i, err := strconv.ParseUint(d, 0, 64)
		if err != nil {
			return reflect.Value{}, defaultError{err}
		}
		return flag.Uint(name, uint(i), usage), nil
	case reflect.Float64:
		f, err := strconv.ParseFloat(d, 64)
		if err != nil {
			return reflect.Value{}, defaultError{err}
		}
		return flag.Float64(name, f, usage), nil
	case reflect.String:
		return flag.String(name, d, usage), nil
	case reflect.Bool:
		b, err := strconv.ParseBool(d)
		if err != nil {
			return reflect.Value{}, defaultError{err}
		}
		return flag.Bool(name, b, usage), nil
	case reflect.Struct:
		return getFlags(field.Type, name)
	}
	return reflect.Value{}, fmt.Errorf("missing kind %s", field.Type.Kind())
}

func getName(field reflect.StructField, parent string) string {
	name := field.Tag.Get("name")
	if name == "" {
		name = field.Name
	}
	if parent != "" {
		name = parent + "." + name
	}
	return name
}
