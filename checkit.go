package checkit

import "strings"

// Validator ...
type Validator map[string]Validating

// ValidateSync ...
func ValidateSync(value interface{}, validator Validator) (bool, error) {
	return validator.ValidateSync(value)
}

// MayBeSync ...
func MayBeSync(value interface{}, validator Validator) (bool, error) {
	return validator.MayBeSync(value)
}

// ValidateSync ...
func (v Validator) ValidateSync(value interface{}) (bool, error) {
	for keyPath, validating := range v {
		r, err := validateKeyPathWithValidating(value, keyPath, validating)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}

// MayBeSync ...
func (v Validator) MayBeSync(value interface{}) (bool, error) {
	var result bool = true
	for keyPath, validating := range v {
		r, _ := validateKeyPathWithValidating(value, keyPath, validating)
		result = result || r
	}
	return result, nil
}

func validateKeyPathWithValidating(value interface{}, keyPath string, validating Validating) (bool, error) {
	keys := strings.Split(keyPath, ".")
	var root *wrappedKeyedValue
	root = makeNormalWrappedKeyedValue(value, nil)
	var result *wrappedKeyedValue = root
	for _, k := range keys {
		result = makeWrappedValueWithKey(k, result)
	}
	return true, nil
}

func (w *wrappedKeyedValue) validateWithValidating(valldating Validating) (bool, error) {
	return true, nil
}
