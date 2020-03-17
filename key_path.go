package checkit

import (
	"reflect"
	"strconv"
)

const (
	keyAll   = "all"
	keyAny   = "any"
	keyFirst = "first"
	keyLast  = "last"
)

type wrappedKeyedValue struct {
	value              interface{}
	shouldValidateNorm bool
	shouldValidateAny  bool
	shouldValidateAll  bool

	next *wrappedKeyedValue
}

func makeWrappedValueWithKey(key string, parent *wrappedKeyedValue) *wrappedKeyedValue {
	switch v := parent.value.(type) {
	case []interface{}:
		switch key {
		case keyAll:
			return makeWrappedKeyedValue(v, false, true, parent)
		case keyAny:
			return makeWrappedKeyedValue(v, true, false, parent)
		case keyFirst:
			return makeNormalWrappedKeyedValue(v[0], parent)
		case keyLast:
			return makeNormalWrappedKeyedValue(v[len(v)-1], parent)
		default:
			if intValue, err := strconv.Atoi(key); err == nil {
				return makeNormalWrappedKeyedValue(v[intValue], parent)
			}
		}
	case map[int8]interface{}:
		if intValue, err := strconv.ParseInt(key, 10, 8); err == nil {
			return makeNormalWrappedKeyedValue(v[int8(intValue)], parent)
		}
	case map[int16]interface{}:
		if intValue, err := strconv.ParseInt(key, 10, 16); err == nil {
			return makeNormalWrappedKeyedValue(v[int16(intValue)], parent)
		}
	case map[int32]interface{}:
		if intValue, err := strconv.ParseInt(key, 10, 32); err == nil {
			return makeNormalWrappedKeyedValue(v[int32(intValue)], parent)
		}
	case map[int64]interface{}:
		if intValue, err := strconv.ParseInt(key, 10, 64); err == nil {
			return makeNormalWrappedKeyedValue(v[int64(intValue)], parent)
		}
	case map[uint8]interface{}:
		if uintValue, err := strconv.ParseUint(key, 10, 8); err == nil {
			return makeNormalWrappedKeyedValue(v[uint8(uintValue)], parent)
		}
	case map[uint16]interface{}:
		if uintValue, err := strconv.ParseUint(key, 10, 16); err == nil {
			return makeNormalWrappedKeyedValue(v[uint16(uintValue)], parent)
		}
	case map[uint32]interface{}:
		if uintValue, err := strconv.ParseUint(key, 10, 32); err == nil {
			return makeNormalWrappedKeyedValue(v[uint32(uintValue)], parent)
		}
	case map[uint64]interface{}:
		if uintValue, err := strconv.ParseUint(key, 10, 64); err == nil {
			return makeNormalWrappedKeyedValue(v[uint64(uintValue)], parent)
		}
	case map[string]interface{}:
		return makeNormalWrappedKeyedValue(v[key], parent)
	default:
		structValue := reflect.ValueOf(parent.value).Elem()
		structType := structValue.Type()
		for i := 0; i < structValue.NumField(); i++ {
			if structType.Field(i).Name == key {
				return makeNormalWrappedKeyedValue(structValue.Field(i).Interface(), parent)
			}
		}
	}
	return nil
}

func makeNormalWrappedKeyedValue(value interface{}, parent *wrappedKeyedValue) *wrappedKeyedValue {
	return &wrappedKeyedValue{
		value:              value,
		shouldValidateNorm: true,
		shouldValidateAny:  false,
		shouldValidateAll:  false,
		next:               nil,
	}
}

func makeWrappedKeyedValue(value interface{}, shouldValidateAny bool, shouldValidateAll bool, parent *wrappedKeyedValue) *wrappedKeyedValue {
	newWrappedKeyedValue := &wrappedKeyedValue{
		value:              value,
		shouldValidateNorm: false,
		shouldValidateAny:  shouldValidateAny,
		shouldValidateAll:  shouldValidateAll,
		next:               nil,
	}
	parent.next = newWrappedKeyedValue
	return newWrappedKeyedValue
}

func contains(arr []string, check string) bool {
	for _, e := range arr {
		if e == check {
			return true
		}
	}
	return false
}
