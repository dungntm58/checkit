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

	children []wrappedKeyedValue
}

func getValueForKey(key string, obj interface{}) interface{} {
	switch v := obj.(type) {
	case []interface{}:
		switch key {
		case keyAll, keyAny:
			return v
		case keyFirst:
			return v[0]
		case keyLast:
			return v[len(v)-1]
		default:
			if intValue, err := strconv.Atoi(key); err == nil {
				return v[intValue]
			}
		}
	case map[int8]interface{}:
		if intValue, err := strconv.ParseInt(key, 10, 8); err == nil {
			return v[int8(intValue)]
		}
	case map[int16]interface{}:
		if intValue, err := strconv.ParseInt(key, 10, 16); err == nil {
			return v[int16(intValue)]
		}
	case map[int32]interface{}:
		if intValue, err := strconv.ParseInt(key, 10, 32); err == nil {
			return v[int32(intValue)]
		}
	case map[int64]interface{}:
		if intValue, err := strconv.ParseInt(key, 10, 64); err == nil {
			return v[int64(intValue)]
		}
	case map[uint8]interface{}:
		if uintValue, err := strconv.ParseUint(key, 10, 8); err == nil {
			return v[uint8(uintValue)]
		}
	case map[uint16]interface{}:
		if uintValue, err := strconv.ParseUint(key, 10, 16); err == nil {
			return v[uint16(uintValue)]
		}
	case map[uint32]interface{}:
		if uintValue, err := strconv.ParseUint(key, 10, 32); err == nil {
			return v[uint32(uintValue)]
		}
	case map[uint64]interface{}:
		if uintValue, err := strconv.ParseUint(key, 10, 64); err == nil {
			return v[uint64(uintValue)]
		}
	case map[string]interface{}:
		return v[key]
	default:
		structValue := reflect.ValueOf(obj).Elem()
		structType := structValue.Type()
		for i := 0; i < structValue.NumField(); i++ {
			if structType.Field(i).Name == key {
				return structValue.Field(i).Interface()
			}
		}
	}
	return nil
}

func makeNormalWrappedKeyedValue(value interface{}, parent *wrappedKeyedValue) wrappedKeyedValue {
	newWrappedKeyedValue := wrappedKeyedValue{
		value:              value,
		shouldValidateNorm: true,
		shouldValidateAny:  false,
		shouldValidateAll:  false,
		children:           []wrappedKeyedValue{},
	}
	if parent != nil {
		parent.children = append(parent.children, newWrappedKeyedValue)
	}
	return newWrappedKeyedValue
}

func buildWrappedKeyValueWithKeys(keys []string, keyIndex int, value interface{}, parent *wrappedKeyedValue) {
	if keyIndex == len(keys) {
		return
	}
	key := keys[keyIndex]
	keyedValue := getValueForKey(key, value)
	if keyedValue == nil {
		return
	}
	switch key {
	case keyAny:
		arr, _ := keyedValue.([]interface{})
		parent.shouldValidateAny = true
		parent.shouldValidateAll = false
		parent.shouldValidateNorm = false
		for _, el := range arr {
			buildWrappedKeyValueWithKeys(keys, keyIndex+1, el, parent)
		}
	case keyAll:
		arr, _ := keyedValue.([]interface{})
		parent.shouldValidateAny = false
		parent.shouldValidateAll = true
		parent.shouldValidateNorm = false
		for _, el := range arr {
			buildWrappedKeyValueWithKeys(keys, keyIndex+1, el, parent)
		}
	default:
		newWrappedKeyedValue := makeNormalWrappedKeyedValue(keyedValue, parent)
		buildWrappedKeyValueWithKeys(keys, keyIndex+1, keyedValue, &newWrappedKeyedValue)
	}
}
