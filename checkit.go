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
	var result bool = false
	for keyPath, validating := range v {
		r, _ := validateKeyPathWithValidating(value, keyPath, validating)
		result = result || r
	}
	return result, nil
}

func validateKeyPathWithValidating(value interface{}, keyPath string, validating Validating) (bool, error) {
	var keys []string = []string{}
	for _, k := range strings.Split(keyPath, ".") {
		if len(k) > 0 {
			keys = append(keys, k)
		}
	}
	if len(keys) == 0 {
		return validating.Validate(value)
	}

	rootWrapped := makeNormalWrappedKeyedValue(value, nil)
	root := &rootWrapped
	buildWrappedKeyValueWithKeys(keys, 0, value, root)

	return root.validateWithValidating(validating)
}

func (w *wrappedKeyedValue) validateWithValidating(validating Validating) (bool, error) {
	if w.shouldValidateAll {
		for _, child := range w.children {
			r, err := child.validateWithValidating(validating)
			if err != nil {
				return false, err
			}
			if !r {
				return false, nil
			}
		}
		return true, nil
	} else if w.shouldValidateAny {
		var result bool = false
		for _, child := range w.children {
			r, _ := child.validateWithValidating(validating)
			result = result || r
		}
		return result, nil
	}
	return validating.Validate(w.value)
}
