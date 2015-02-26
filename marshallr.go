package marshallr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func MarshalJSONLowerFirst(myStruct interface{}) ([]byte, error) {
	s := reflect.ValueOf(myStruct)
	kindOfT := s.Kind()
	typeOfT := s.Type()
	println(fmt.Sprintf("This is a K: %v T: %v", kindOfT, typeOfT))
	if kindOfT != reflect.Struct {
		v2 := reflect.ValueOf(&myStruct)
		if kindOfT == reflect.Ptr {
			return nil, fmt.Errorf("This is not a struct. It's a Pointer. %v", v2)
		}
		if kindOfT == reflect.Interface {
			return nil, fmt.Errorf("This is not a struct. It's an Interface. %v", v2)
		}
		return nil, fmt.Errorf("This is not a struct. It's a '%d'", kindOfT)

	}
	var buf bytes.Buffer
	buf.WriteRune('{')
	shouldComma := false
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if !f.CanInterface() {
			continue
		}
		t := s.Type().Field(i)
		tag := t.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name := tag
		options := ""
		if idx := strings.Index(tag, ","); idx != -1 {
			name = tag[:idx]
			options = tag[idx+1:]
		}
		if strings.Contains(options, "omitempty") {
			if isEmptyValue(f) {
				continue
			}
		}
		//
		if shouldComma {
			buf.WriteRune(',')
		}
		shouldComma = true
		//marshal where appropriate
		if name == "" {
			myname := typeOfT.Field(i).Name
			p1 := myname[0:1]
			name = strings.ToLower(p1)
			if len(myname) > 1 {
				p2 := myname[1:]
				name += p2
			}
		}
		bn, err := json.Marshal(name)
		if err != nil {
			return nil, err
		}
		buf.Write(bn)
		buf.WriteRune(':')
		b, err := json.Marshal(f.Interface())
		if err != nil {
			return nil, err
		}
		buf.Write(b)

	}
	buf.WriteRune('}')
	//println(buf.String())
	return buf.Bytes(), nil
}


func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
