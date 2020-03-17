package checkit

import (
	"testing"
)

func TestSomething1(t *testing.T) {
	Validator(map[string]Validating{
		"a": CompoundValidating{
			Integer(),
		},
	}).ValidateSync(struct {
		a int
	}{
		a: 1,
	})
}
