package mxj

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAnyXmlHeader(t *testing.T) {
	fmt.Println("\n----------------  anyxml_test.go ...\n")
}

var anydata = []byte(`[
    {
        "somekey": "somevalue"
    },
    {
        "somekey": "somevalue"
    },
    {
        "somekey": "somevalue"
    }
]`)

func TestAnyXml(t *testing.T) {
	var i interface{}
	err := json.Unmarshal(anydata,&i)
	x, err := AnyXml(i)
	if err != nil {
		t.Fatal(err)
	}
   fmt.Println("[]->x:", string(x))

	x, err = AnyXml(3.14159625)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("f->x:", string(x))
}

func TestAnyXmlIndent(t *testing.T) {
	var i interface{}
	err := json.Unmarshal(anydata,&i)
	x, err := AnyXmlIndent(i, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
   fmt.Println("[]->x:\n", string(x))

	x, err = AnyXmlIndent(3.14159625, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("f->x:\n", string(x))
}