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

func getValueWithKeyPath(keyPath string, any interface{}) (value interface{}, isResultAny bool, isResultAll bool) {
	keys := strings.Split(keyPath, ".")
	isResultAny = contains(keys, keyAny)
	isResultAll = contains(keys, keyAll)

	var result interface{}
	result = any
	for _, k := range keys {
		shouldMap := k == keyAny || k == keyAll
		result = getValueWithKey(k, result, shouldMap)
	}
	value = result
	return
}

func getValueWithKey(key string, any interface{}, shouldMap bool) interface{} {
	switch v := any.(type) {
	case []interface{}:
		switch key {
		case keyAll:
		case keyAny:
			if !shouldMap {
				return v
			}
			var result []interface{}
			for _, el := range v {
				result = append(result, getValueWithKey(key, el, false))
			}
			return result
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
