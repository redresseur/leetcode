package nested

import (
	"encoding/json"
	"testing"
)

var str = `{ "A": 1, 
	"B.A": 2, 
	"B.B": 3, 
	"CC.D.E": 4, 
	"CC.D.F": 5
}`

func TestNestedDict_Transfer(t *testing.T) {
	dict := NewNestedDict()
	if err := dict.Transform([]byte(str)); err != nil {
		t.Fatalf("TestNestedDict_Transfer: %s", err.Error())
	} else {
		trans, _ := json.Marshal(&dict.root)
		t.Logf("passed : %s", trans)
	}

	if data, err := dict.Marshal(); err != nil {
		t.Fatalf("TestNestedDict_Transfer: %s", err.Error())
	} else {
		t.Logf("passed : %s", string(data))
	}

}
