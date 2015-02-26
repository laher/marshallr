package marshallr

import (
	"encoding/json"
	"testing"
)

type T1 struct {
	A int
	B string
}

func (t1 T1) MarshalJSON() ([]byte, error) {
	return MarshalJSONLowerFirst(t1)
}

type T2 struct {
	Abra    int
	Babra   string
	Cadabra T1
}

func (t2 T2) MarshalJSON() ([]byte, error) {
	return MarshalJSONLowerFirst(t2)
}

type T3 struct {
	Abra  int
	babra string
}

func (t3 T3) MarshalJSON() ([]byte, error) {
	return MarshalJSONLowerFirst(t3)
}

func Test1(t *testing.T) {
	sitem := T1{23, "skidoo"}
	item := T2{134, "yes", sitem}

	//b, err := item.MarshalJSON()
	b, err := json.Marshal(item)
	if err != nil {
		t.Errorf("Error %v", err)
	}

	t.Logf("Result: %s\n", string(b))
}

func Test2(t *testing.T) {
	item := T3{134, "yes"}

	b, err := json.Marshal(item)
	if err != nil {
		t.Errorf("Error %v", err)
	}

	t.Logf("Result: %s\n", string(b))
}
