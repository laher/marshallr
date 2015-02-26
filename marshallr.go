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
		if f.CanInterface() {
			if shouldComma {
				buf.WriteRune(',')
			}
			shouldComma = true
			//marshal where appropriate
			name := typeOfT.Field(i).Name
			p1 := name[0:1]
			out := strings.ToLower(p1)
			if len(name) > 1 {
				p2 := name[1:]
				out += p2
			}
			bn, err := json.Marshal(out)
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
	}
	buf.WriteRune('}')
	//println(buf.String())
	return buf.Bytes(), nil
}
