package checkit

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetValueOfKey_whenKeyIsAllAndValueIsASlice_shouldBeTheSame(t *testing.T) {
	var arr = []int{0, 1}
	var sl interface{} = arr[0:2]
	r := getValueForKey(keyAll, sl)
	q := getValueForKey(keyAny, sl)
	if !reflect.DeepEqual(r, sl) {
		t.Errorf("Result must be the given array")
	}
	if !reflect.DeepEqual(q, sl) {
		t.Errorf("Result must be the given array")
	}
}

func TestComplexNestedStruct(t *testing.T) {
	type str struct {
		a []int
	}

	keys := []string{keyAny, "a", keyAll}
	value := []str{
		str{a: []int{0, 1}},
		str{a: []int{1, 2}},
	}
	root := makeNormalWrappedKeyedValue(value, nil)
	buildWrappedKeyValueWithKeys(keys, 0, value, root)

	if len(root.children) != 2 {
		t.Errorf("Root children count must be %d", 2)
	}
	if len(root.children[0].children) != 1 {
		t.Errorf("Root first grand children count must be %d", 1)
	}
	fmt.Println(root.children[0].children[0].children)
	if len(root.children[0].children[0].children) != 2 {
		t.Errorf("Root first grand grand children count must be %d", 2)
	}
}
