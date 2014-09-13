package deepcopy

import (
	"reflect"
	"testing"
)

type A struct {
	String  string
	Int     int
	Strings []string
	Ints    map[string]int
}

func TestCopy(t *testing.T) {
	src := &A{
		String:  "Hello World",
		Int:     5,
		Strings: []string{"A", "B"},
		Ints:    map[string]int{"A": 1, "B": 2},
	}
	dst := &A{Strings: []string{"C"}, Ints: map[string]int{"B": 3, "C": 4}}
	expected := &A{
		String:  "Hello World",
		Int:     5,
		Strings: []string{"A", "B"},
		Ints:    map[string]int{"A": 1, "B": 2, "C": 4},
	}
	err := Copy(dst, src)
	t.Log(dst)
	if err != nil {
		t.Errorf("Unable to copy!")
	}
	if !reflect.DeepEqual(expected, dst) {
		t.Errorf("expected and dst differed")
	}
}
