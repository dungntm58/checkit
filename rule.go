package checkit

import "regexp"

// Validating ...
type Validating interface {
	Validate(value interface{}) (bool, *string)
}

// Accepted ...
func Accepted() Validating {
	return &validator{
		validateFunc: func(value interface{}) bool {
			switch v := value.(type) {
			case string:
				acceptedValues := []string{"yes", "on", "1"}
				return contains(acceptedValues, v)
			case int8:
				return v == 1
			case int16:
				return v == 1
			case int32:
				return v == 1
			case int64:
				return v == 1
			case uint8:
				return v == 1
			case uint16:
				return v == 1
			case uint32:
				return v == 1
			case uint64:
				return v == 1
			default:
				return false
			}
		},
		errorMessage: "The value must be yes, on, or 1. This is useful for validating \"Terms of Service\" acceptance.",
	}
}

// Aplha ...
func Aplha() Validating {
	return &validator{
		validateFunc: func(value interface{}) bool {
			return matchAnyWithRegex(`^[a-zA-Z]+$`, value)
		},
		errorMessage: "The value must be entirely alphabetic characters.",
	}
}

// AplhaDash ...
func AplhaDash() Validating {
	return &validator{
		validateFunc: func(value interface{}) bool {
			return matchAnyWithRegex(`^[a-z0-9_\-]+$`, value)
		},
		errorMessage: "The value may have alpha-numeric characters, as well as dashes and underscores.",
	}
}

// AplhaNumeric ...
func AplhaNumeric() Validating {
	return &validator{
		validateFunc: func(value interface{}) bool {
			return matchAnyWithRegex(`^[a-zA-Z0-9]+$`, value)
		},
		errorMessage: "The value must be entirely alpha-numeric characters.",
	}
}

// AplhaUnderscore ...
func AplhaUnderscore() Validating {
	return &validator{
		validateFunc: func(value interface{}) bool {
			return matchAnyWithRegex(`^[a-zA-Z0-9_]+$`, value)
		},
		errorMessage: "The value must be entirely alpha-numeric, with underscores but not dashes.",
	}
}

// Base64 ...
func Base64() Validating {
	return &validator{
		validateFunc: func(value interface{}) bool {
			return matchAnyWithRegex(`^(?:[A-Za-z0-9+\/]{4})*(?:[A-Za-z0-9+\/]{2}==|[A-Za-z0-9+\/]{3}=)?$`, value)
		},
		errorMessage: "The value must be a base64 encoded value.",
	}
}

// Email ...
func Email() Validating {
	return &validator{
		validateFunc: func(value interface{}) bool {
			return matchAnyWithRegex(`^(.+)@(.+)\.(.+)$`, value)
		},
		errorMessage: "The field must be a valid formatted e-mail address.",
	}
}

// Integer ...
func Integer() Validating {
	return &validator{
		validateFunc: func(value interface{}) bool {
			switch v := value.(type) {
			case string:
				return regexp.MustCompile(`^\-?[0-9]+$`).MatchString(v)
			case int8, int16, int32, int64,
				uint8, uint16, uint32, uint64:
				return true
			}
			return false
		},
		errorMessage: "The value must have an integer value.",
	}
}

// Ipv4 ...
func Ipv4() Validating {
	return &validator{
		validateFunc: func(value interface{}) bool {
			return matchAnyWithRegex(`^((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\.){3}(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})$`, value)
		},
		errorMessage: "The value must be formatted as an IPv4 address.",
	}
}

// Ipv6 ...
func Ipv6() Validating {
	return &validator{
		validateFunc: func(value interface{}) bool {
			return matchAnyWithRegex(`^((?=.*::)(?!.*::.+::)(::)?([\dA-F]{1,4}:(:|\b)|){5}|([\dA-F]{1,4}:){6})((([\dA-F]{1,4}((?!\3)::|:\b|$))|(?!\2\3)){2}|(((2[0-4]|1\d|[1-9])?\d|25[0-5])\.?\b){4})$`, value)
		},
		errorMessage: "The value must be formatted as an IPv6 address.",
	}
}

// Luhn ...
func Luhn() Validating {
	return &validator{
		validateFunc: func(value interface{}) bool {
			return matchAnyWithRegex(`^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`, value)
		},
		errorMessage: "The given value must pass a basic luhn (credit card) check regular expression.",
	}
}

// Natural ...
func Natural() Validating {
	return &validator{
		validateFunc: func(value interface{}) bool {
			switch v := value.(type) {
			case string:
				return regexp.MustCompile(`^[0-9]+$`).MatchString(v)
			case int8:
				return v >= 0
			case int16:
				return v >= 0
			case int32:
				return v >= 0
			case int64:
				return v >= 0
			case uint8, uint16, uint32, uint64:
				return true
			default:
				return false
			}
		},
		errorMessage: "The value must be a natural number (a number greater than or equal to 0).",
	}
}

// NaturalNonZero ...
func NaturalNonZero() Validating {
	return &validator{
		validateFunc: func(value interface{}) bool {
			switch v := value.(type) {
			case string:
				return regexp.MustCompile(`^[1-9][0-9]*$`).MatchString(v)
			case int8:
				return v > 0
			case int16:
				return v > 0
			case int32:
				return v > 0
			case int64:
				return v > 0
			case uint8:
				return v > 0
			case uint16:
				return v > 0
			case uint32:
				return v > 0
			case uint64:
				return v > 0
			default:
				return false
			}
		},
		errorMessage: "The value must be a natural number, greater than or equal to 1.",
	}
}

// URL ...
func URL() Validating {
	return &validator{
		validateFunc: func(value interface{}) bool {
			return matchAnyWithRegex(`^((http|https):\/\/(\w+:{0,1}\w*@)?(\S+)|)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?$`, value)
		},
		errorMessage: "The value must be formatted as an URL.",
	}
}

// UUID ...
func UUID() Validating {
	return &validator{
		validateFunc: func(value interface{}) bool {
			return matchAnyWithRegex(`^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`, value)
		},
		errorMessage: "Passes for a validly formatted UUID.",
	}
}

func matchAnyWithRegex(regex string, any interface{}) bool {
	switch v := any.(type) {
	case string:
		return regexp.MustCompile(regex).MatchString(v)
	default:
		return false
	}
}

type validateFunc func(interface{}) bool

type validator struct {
	validateFunc validateFunc
	errorMessage string
}

func (v *validator) Validate(value interface{}) (bool, *string) {
	if v.validateFunc(value) {
		return true, nil
	}
	return false, &v.errorMessage
}
