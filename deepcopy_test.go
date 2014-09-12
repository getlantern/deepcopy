package deepcopy

import (
	"reflect"
	"testing"
)

type A struct {
	String string
	Int    int
}

func TestTypeMismatch(t *testing.T) {
	err := Copy(&A{}, "Hello World")
	if err == nil {
		t.Errorf("Expected error!")
	}
}

func TestCopy(t *testing.T) {
	src := &A{
		String: "Hello World",
		Int:    5,
	}
	dst := &A{}
	err := Copy(dst, src)
	if err != nil {
		t.Errorf("Unable to copy!")
	}
	if !reflect.DeepEqual(src, dst) {
		t.Errorf("src and dst differed")
	}
}
