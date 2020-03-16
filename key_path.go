package checkit

import (
	"reflect"
	"strconv"
	"strings"
)

const (
	keyAll   = "all"
	keyAny   = "any"
	keyFirst = "first"
	keyLast  = "last"
)

const (
	invalidArrayIndex     = -1
	virtualLastArrayIndex = -2
	anyArrayIndex         = -3
	allArrayIndex         = -4
)

func getValueWithKeyPath(keyPath string, any interface{}) (value interface{}, isAnyResult bool) {
	keys := strings.Split(keyPath, ".")
	isAnyResult = contains(keys, keyAny)

	var result interface{}
	result = any
	for _, k := range keys {
		result = getValueWithKey(k, result)
	}
	value = result
	return
}

func getValueWithKey(key string, any interface{}) interface{} {
	switch v := any.(type) {
	case []interface{}:
		switch key {
		case keyAll:
		case keyAny:
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
		break
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
		if intValue, err := strconv.ParseUint(key, 10, 8); err == nil {
			return v[uint8(intValue)]
		}
	case map[uint16]interface{}:
		if intValue, err := strconv.ParseUint(key, 10, 16); err == nil {
			return v[uint16(intValue)]
		}
	case map[uint32]interface{}:
		if intValue, err := strconv.ParseUint(key, 10, 32); err == nil {
			return v[uint32(intValue)]
		}
	case map[uint64]interface{}:
		if intValue, err := strconv.ParseUint(key, 10, 64); err == nil {
			return v[uint64(intValue)]
		}
	case map[string]interface{}:
		return v[key]
	default:
		structValue := reflect.ValueOf(any).Elem()
		structType := structValue.Type()
		for i := 0; i < structValue.NumField(); i++ {
			if structType.Field(i).Name == key {
				return structValue.Field(i).Interface()
			}
		}
	}
	return nil
}

func contains(arr []string, check string) bool {
	for _, e := range arr {
		if e == check {
			return true
		}
	}
	return false
}
