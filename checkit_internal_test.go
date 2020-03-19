package checkit

import (
	"testing"
)

func TestSomething1(t *testing.T) {
	str := "ab"
	_, err := Validator(map[string]Validating{
		"a": Between(0, 2),
		"b": MaxLength(2),
		"c": ExactLength(3),
	}).ValidateSync(struct {
		a int
		b *string
		c []interface{}
	}{
		a: 1,
		b: &str,
		c: []interface{}{0, 1, 2},
	})
	if err != nil {
		t.Error(err)
	}
}
