package checkit

import (
	"errors"
	"math"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// Validating ...
type Validating interface {
	Validate(value interface{}) (bool, error)
}

// CompoundValidating ...
type CompoundValidating []Validating

// Validate ...
func (c CompoundValidating) Validate(value interface{}) (bool, error) {
	for _, v := range c {
		_r, err := v.Validate(value)
		if err != nil {
			return false, err
		}
		if !_r {
			return false, nil
		}
	}
	return true, nil
}

// Accepted ...
func Accepted() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			switch v := value.(type) {
			case string:
				acceptedValues := []string{"yes", "on", "1"}
				return contains(acceptedValues, strings.ToLower(v)), nil
			case int8:
				return v == 1, nil
			case int16:
				return v == 1, nil
			case int32:
				return v == 1, nil
			case int64:
				return v == 1, nil
			case uint8:
				return v == 1, nil
			case uint16:
				return v == 1, nil
			case uint32:
				return v == 1, nil
			case uint64:
				return v == 1, nil
			default:
				return false, nil
			}
		},
		errorMessage: "The value must be yes, on, or 1. This is useful for validating \"Terms of Service\" acceptance.",
	}
}

// Aplha ...
func Aplha() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			return matchAnyWithRegex(regexAlpha, value)
		},
		errorMessage: "The value must be entirely alphabetic characters.",
	}
}

// AplhaDash ...
func AplhaDash() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			return matchAnyWithRegex(regexAlphaDash, value)
		},
		errorMessage: "The value may have alpha-numeric characters, as well as dashes and underscores.",
	}
}

// AplhaNumeric ...
func AplhaNumeric() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			return matchAnyWithRegex(regexAlphaNumeric, value)
		},
		errorMessage: "The value must be entirely alpha-numeric characters.",
	}
}

// AplhaUnderscore ...
func AplhaUnderscore() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			return matchAnyWithRegex(regexAlphaUnderscore, value)
		},
		errorMessage: "The value must be entirely alpha-numeric, with underscores but not dashes.",
	}
}

// Array ...
func Array() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			switch reflect.TypeOf(value).Kind() {
			case reflect.Array, reflect.Slice:
				return true, nil
			default:
				return false, nil
			}
		},
		errorMessage: "The value must be a valid array object.",
	}
}

// Base64 ...
func Base64() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			return matchAnyWithRegex(regexBase64, value)
		},
		errorMessage: "The value must be a base64 encoded value.",
	}
}

// Between ...
func Between(min interface{}, max interface{}) Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			lCompare, lErr := lessThanEqualTo(min, value)
			if lErr != nil {
				return false, lErr
			}
			rCompare, rErr := lessThanEqualTo(min, value)
			if rErr != nil {
				return false, rErr
			}
			return lCompare && rCompare, nil
		},
		errorMessage: "The value must have a size between the given min and max.",
	}
}

// Boolean ...
func Boolean() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			switch reflect.TypeOf(value).Kind() {
			case reflect.Bool:
				return true, nil
			default:
				return false, nil
			}
		},
		errorMessage: "The value must be a boolean.",
	}
}

// Contains ...
func Contains(v interface{}) Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			arr := reflect.ValueOf(value)
			if arr.Kind() != reflect.Array && arr.Kind() != reflect.Slice {
				return false, errors.New("The value must be an array or a slice")
			}
			for i := 0; i < arr.Len(); i++ {
				if arr.Index(i).Interface() == v {
					return true, nil
				}
			}
			return false, nil
		},
		errorMessage: "The value must contain the value.",
	}
}

// Date ...
func Date() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			switch value.(type) {
			case time.Time:
				return true, nil
			default:
				return false, nil
			}
		},
		errorMessage: "The value must be a valid date object.",
	}
}

// Email ...
func Email() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			return matchAnyWithRegex(regexEmail, value)
		},
		errorMessage: "The field must be a valid formatted e-mail address.",
	}
}

// Empty ...
func Empty() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			switch v := value.(type) {
			case []interface{}:
				return len(v) == 0, nil
			case string:
				return len(v) == 0, nil
			default:
				return false, errors.New("len() is not supported")
			}
		},
		errorMessage: "The value must be a empty collection.",
	}
}

// ExactLength ...
func ExactLength(length int) Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			switch v := value.(type) {
			case []interface{}:
				return len(v) == length, nil
			case string:
				return len(v) == length, nil
			default:
				return false, errors.New("len() is not supported")
			}
		},
		errorMessage: "The field must have the exact length of \"val\".",
	}
}

// ExistsNonNil ...
func ExistsNonNil() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			return value != nil, nil
		},
		errorMessage: "The value under validation must not be undefined or nil.",
	}
}

// Finite ...
func Finite() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			switch v := value.(type) {
			case float64:
				if v == 0 {
					return true, nil
				}
				return !math.IsInf(v, 0), nil
			case int8, int16, int32, int64,
				uint8, uint16, uint32, uint64, float32:
				return true, nil
			default:
				return false, errors.New("The value must be a number")
			}
		},
		errorMessage: "The value under validation must be a finite number.",
	}
}

// Function ...
func Function() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			switch reflect.TypeOf(value).Kind() {
			case reflect.Func:
				return true, nil
			default:
				return false, nil
			}
		},
		errorMessage: "The value must be a function.",
	}
}

// GreaterThan ...
func GreaterThan(v interface{}) Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			lessThanEqualTo, err := lessThanEqualTo(value, v)
			if err != nil {
				return false, err
			}
			return !lessThanEqualTo, nil
		},
		errorMessage: "The value under validation must be \"greater than\" the given value.",
	}
}

// GreaterThanEqualTo ...
func GreaterThanEqualTo(v interface{}) Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			return greatThanEqualTo(value, v)
		},
		errorMessage: "The value under validation must be \"greater than\" or \"equal to\" the given value.",
	}
}

// Integer ...
func Integer() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			switch v := value.(type) {
			case string:
				return regexp.MustCompile(regexInteger).MatchString(v), nil
			case int8, int16, int32, int64,
				uint8, uint16, uint32, uint64:
				return true, nil
			}
			return false, nil
		},
		errorMessage: "The value must have an integer value.",
	}
}

// Ipv4 ...
func Ipv4() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			return matchAnyWithRegex(regexIpv4, value)
		},
		errorMessage: "The value must be formatted as an IPv4 address.",
	}
}

// Ipv6 ...
func Ipv6() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			return matchAnyWithRegex(regexIpv6, value)
		},
		errorMessage: "The value must be formatted as an IPv6 address.",
	}
}

// LessThan ...
func LessThan(v interface{}) Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			greatThanEqualTo, err := greatThanEqualTo(value, v)
			if err != nil {
				return false, err
			}
			return !greatThanEqualTo, nil
		},
		errorMessage: "The value under validation must be \"less than\" the given value.",
	}
}

// LessThanEqualTo ...
func LessThanEqualTo(v interface{}) Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			return lessThanEqualTo(value, v)
		},
		errorMessage: "The value under validation must be \"less than\" or \"equal to\" the given value.",
	}
}

// Luhn ...
func Luhn() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			return matchAnyWithRegex(regexLuhn, value)
		},
		errorMessage: "The given value must pass a basic luhn (credit card) check regular expression.",
	}
}

// Natural ...
func Natural() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			switch v := value.(type) {
			case string:
				return regexp.MustCompile(regexNatural).MatchString(v), nil
			case int8:
				return v >= 0, nil
			case int16:
				return v >= 0, nil
			case int32:
				return v >= 0, nil
			case int64:
				return v >= 0, nil
			case uint8, uint16, uint32, uint64:
				return true, nil
			default:
				return false, errors.New("Cannot compare to zero")
			}
		},
		errorMessage: "The value must be a natural number (a number greater than or equal to 0).",
	}
}

// NaN ...
func NaN() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			switch v := value.(type) {
			case float64:
				return math.IsNaN(v), nil
			case int8, int16, int32, int64,
				uint8, uint16, uint32, uint64, float32:
				return false, nil
			default:
				return false, errors.New("The value must be a number")
			}
		},
		errorMessage: "The value under validation must be a NaN.",
	}
}

// NaturalNonZero ...
func NaturalNonZero() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			switch v := value.(type) {
			case string:
				return regexp.MustCompile(regexNaturalNonZero).MatchString(v), nil
			case int8:
				return v > 0, nil
			case int16:
				return v > 0, nil
			case int32:
				return v > 0, nil
			case int64:
				return v > 0, nil
			case uint8:
				return v > 0, nil
			case uint16:
				return v > 0, nil
			case uint32:
				return v > 0, nil
			case uint64:
				return v > 0, nil
			default:
				return false, errors.New("Cannot compare to zero")
			}
		},
		errorMessage: "The value must be a natural number, greater than or equal to 1.",
	}
}

// Object ...
func Object() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			switch reflect.TypeOf(value).Kind() {
			case reflect.Invalid, reflect.Func, reflect.UnsafePointer:
				return false, nil
			default:
				return true, nil
			}
		},
		errorMessage: "The value must be a object.",
	}
}

// PlainObject ...
func PlainObject() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			switch reflect.TypeOf(value).Kind() {
			case reflect.Map:
				return true, nil
			default:
				return false, nil
			}
		},
		errorMessage: "The value must be a plain object.",
	}
}

// Regex ...
func Regex() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			switch value.(type) {
			case regexp.Regexp:
				return true, nil
			default:
				return false, nil
			}
		},
		errorMessage: "The value must be a RegExp object.",
	}
}

// String ...
func String() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			switch reflect.TypeOf(value).Kind() {
			case reflect.String:
				return true, nil
			default:
				return false, nil
			}
		},
		errorMessage: "The value must be a function.",
	}
}

// URL ...
func URL() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			return matchAnyWithRegex(regexURL, value)
		},
		errorMessage: "The value must be formatted as an URL.",
	}
}

// UUID ...
func UUID() Validating {
	return &validator{
		validateFunc: func(value interface{}) (bool, error) {
			return matchAnyWithRegex(regexUUID, value)
		},
		errorMessage: "Passes for a validly formatted UUID.",
	}
}

const (
	regexAlpha           = `/^[A_Za-z]+$/i`
	regexAlphaDash       = `/^[A_Za-z0-9_\-]+$/i`
	regexAlphaNumeric    = `/^[A_Za-z0-9]+$/i`
	regexAlphaUnderscore = `/^[A_Za-z0-9_]+$/i`
	regexBase64          = `/^(?:[A-Za-z0-9+\/]{4})*(?:[A-Za-z0-9+\/]{2}==|[A-Za-z0-9+\/]{3}=)?$/`
	regexEmail           = `/^(.+)@(.+)\.(.+)$/i`
	regexInteger         = `/^\-?[0-9]+$/`
	regexIpv4            = `/^((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\.){3}(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})$/i`
	regexIpv6            = `/^((?=.*::)(?!.*::.+::)(::)?([\dA-F]{1,4}:(:|\b)|){5}|([\dA-F]{1,4}:){6})((([\dA-F]{1,4}((?!\3)::|:\b|$))|(?!\2\3)){2}|(((2[0-4]|1\d|[1-9])?\d|25[0-5])\.?\b){4})$/i`
	regexLuhn            = `/^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$/`
	regexNatural         = `/^[0-9]+$/i`
	regexNaturalNonZero  = `/^[1-9][0-9]*$/i`
	regexURL             = `/^((http|https):\/\/(\w+:{0,1}\w*@)?(\S+)|)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?$/`
	regexUUID            = `/^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i`
)

func matchAnyWithRegex(regex string, any interface{}) (bool, error) {
	switch v := any.(type) {
	case string:
		return regexp.MustCompile(regex).MatchString(v), nil
	default:
		return false, errors.New("The value must be a string")
	}
}

func greatThanEqualTo(lhs interface{}, rhs interface{}) (bool, error) {
	lhsString, lokString := lhs.(string)
	rhsString, rokString := rhs.(string)

	if lokString && rokString {
		return lhsString >= rhsString, nil
	} else if (lokString && !rokString) || (!lokString && rokString) {
		return false, errors.New("Cannot compare a string to an instance of other type than string")
	}

	lhsDate, lokDate := lhs.(time.Time)
	rhsDate, rokDate := rhs.(time.Time)

	if lokDate && rokDate {
		return lhsDate.After(rhsDate) || lhsDate.Equal(rhsDate), nil
	} else if (lokDate && !rokDate) || (!lokDate && rokDate) {
		return false, errors.New("Cannot compare a date to an instance of other type than date")
	}

	lhsGreaterThanZero, lhsErr := isGreaterThanZero(lhs)
	if lhsErr != nil {
		return false, lhsErr
	}
	rhsGreaterThanZero, rhsErr := isGreaterThanZero(rhs)
	if rhsErr != nil {
		return false, rhsErr
	}
	if lhsGreaterThanZero && !rhsGreaterThanZero {
		return true, nil
	}
	if !lhsGreaterThanZero && rhsGreaterThanZero {
		return false, nil
	}

	isLHSSignedNumber := isSignedNumber(lhs)
	isRHSSignedNumber := isSignedNumber(rhs)

	if !isLHSSignedNumber || !isRHSSignedNumber {
		lhsUint64, _ := tryConvertToUint64(lhs)
		rhsUint64, _ := tryConvertToUint64(rhs)

		return lhsUint64 >= rhsUint64, nil
	}

	lhsFloat64, _ := tryConvertToFloat64(lhs)
	rhsFloat64, _ := tryConvertToFloat64(rhs)

	return lhsFloat64 >= rhsFloat64, nil
}

func lessThanEqualTo(lhs interface{}, rhs interface{}) (bool, error) {
	lhsString, lokString := lhs.(string)
	rhsString, rokString := rhs.(string)

	if lokString && rokString {
		return lhsString <= rhsString, nil
	} else if (lokString && !rokString) || (!lokString && rokString) {
		return false, errors.New("Cannot compare a string to an instance of other type than string")
	}

	lhsDate, lokDate := lhs.(time.Time)
	rhsDate, rokDate := rhs.(time.Time)

	if lokDate && rokDate {
		return lhsDate.Before(rhsDate) || lhsDate.Equal(rhsDate), nil
	} else if (lokDate && !rokDate) || (!lokDate && rokDate) {
		return false, errors.New("Cannot compare a date to an instance of other type than date")
	}

	lhsGreaterThanZero, lhsErr := isGreaterThanZero(lhs)
	if lhsErr != nil {
		return false, lhsErr
	}
	rhsGreaterThanZero, rhsErr := isGreaterThanZero(rhs)
	if rhsErr != nil {
		return false, rhsErr
	}
	if lhsGreaterThanZero && !rhsGreaterThanZero {
		return false, nil
	}
	if !lhsGreaterThanZero && rhsGreaterThanZero {
		return true, nil
	}

	isLHSSignedNumber := isSignedNumber(lhs)
	isRHSSignedNumber := isSignedNumber(rhs)

	if !isLHSSignedNumber || !isRHSSignedNumber {
		lhsUint64, _ := tryConvertToUint64(lhs)
		rhsUint64, _ := tryConvertToUint64(rhs)

		return lhsUint64 <= rhsUint64, nil
	}

	lhsFloat64, _ := tryConvertToFloat64(lhs)
	rhsFloat64, _ := tryConvertToFloat64(rhs)

	return lhsFloat64 <= rhsFloat64, nil
}

func isSignedNumber(any interface{}) bool {
	switch reflect.TypeOf(any).Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func isGreaterThanZero(any interface{}) (bool, error) {
	switch v := any.(type) {
	case int8:
		return v > 0, nil
	case int16:
		return v > 0, nil
	case int32:
		return v > 0, nil
	case int64:
		return v > 0, nil
	case uint8:
		return v > 0, nil
	case uint16:
		return v > 0, nil
	case uint32:
		return v > 0, nil
	case uint64:
		return v > 0, nil
	case float32:
		return v > 0, nil
	case float64:
		return v > 0, nil
	default:
		return false, errors.New("The value must be a number")
	}
}

func tryConvertToUint64(v interface{}) (uint64, bool) {
	switch _v := v.(type) {
	case uint64:
		return _v, true
	case uint32:
		return uint64(_v), true
	case uint16:
		return uint64(_v), true
	case uint8:
		return uint64(_v), true
	default:
		return 0, false
	}
}

func tryConvertToFloat64(v interface{}) (float64, bool) {
	switch _v := v.(type) {
	case float64:
		return _v, true
	case float32:
		return float64(_v), true
	case int64:
		return float64(_v), true
	case int32:
		return float64(_v), true
	case int16:
		return float64(_v), true
	case int8:
		return float64(_v), true
	default:
		return 0, false
	}
}

type validateFunc func(interface{}) (bool, error)

type validator struct {
	validateFunc validateFunc
	errorMessage string
}

func (v *validator) Validate(value interface{}) (bool, error) {
	result, err := v.validateFunc(value)
	if result {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, errors.New(v.errorMessage)
}
