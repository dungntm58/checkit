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

	children []*wrappedKeyedValue
}

func getValueForKey(key string, obj interface{}) interface{} {
	objValue := reflect.ValueOf(obj)
	switch objValue.Kind() {
	case reflect.Array, reflect.Slice:
		switch key {
		case keyAll, keyAny:
			return obj
		case keyFirst:
			return getReferenceValue(objValue.Index(0))
		case keyLast:
			return getReferenceValue(objValue.Index(objValue.Len() - 1))
		default:
			if intValue, err := strconv.Atoi(key); err == nil {
				return getReferenceValue(objValue.Index(intValue))
			}
			return nil
		}
	case reflect.Map:
		mapKeys := objValue.MapKeys()
		if len(mapKeys) == 0 {
			return nil
		}
		reflectValueOfKey := getReflectKeyInMapKeys(mapKeys, key)
		reflectValueOfValue := objValue.MapIndex(reflectValueOfKey)
		if reflectValueOfValue.Kind() == reflect.Invalid {
			return nil
		}
		return reflectValueOfValue.Interface()
	case reflect.Interface, reflect.Ptr:
		objValue = flattenReflectValue(objValue)
		fallthrough
	case reflect.Struct:
		fieldByName := objValue.FieldByName(key)
		return getReferenceValue(fieldByName)
	default:
		break
	}
	return nil
}

func getReflectKeyInMapKeys(mapKeys []reflect.Value, key string) reflect.Value {
	for _, reflectKey := range mapKeys {
		itKey := reflectKey.Interface()
		switch itVal := itKey.(type) {
		case string:
			if key == itVal {
				return reflectKey
			}
		case int8:
			if intValue, err := strconv.ParseInt(key, 10, 8); err == nil && int8(intValue) == itVal {
				return reflectKey
			}
		case int16:
			if intValue, err := strconv.ParseInt(key, 10, 8); err == nil && int16(intValue) == itVal {
				return reflectKey
			}
		case int32:
			if intValue, err := strconv.ParseInt(key, 10, 8); err == nil && int32(intValue) == itVal {
				return reflectKey
			}
		case int64:
			if intValue, err := strconv.ParseInt(key, 10, 8); err == nil && int64(intValue) == itVal {
				return reflectKey
			}
		case uint8:
			if intValue, err := strconv.ParseInt(key, 10, 8); err == nil && uint8(intValue) == itVal {
				return reflectKey
			}
		case uint16:
			if intValue, err := strconv.ParseInt(key, 10, 8); err == nil && uint16(intValue) == itVal {
				return reflectKey
			}
		case uint32:
			if intValue, err := strconv.ParseInt(key, 10, 8); err == nil && uint32(intValue) == itVal {
				return reflectKey
			}
		case uint64:
			if intValue, err := strconv.ParseInt(key, 10, 8); err == nil && uint64(intValue) == itVal {
				return reflectKey
			}
		default:
			continue
		}
	}
	return reflect.Value{}
}

func flattenReflectValue(value reflect.Value) reflect.Value {
	val := value
	kind := val.Kind()
	for (kind == reflect.Interface || kind == reflect.Ptr) && !val.IsNil() {
		val = val.Elem()
		kind = val.Kind()
	}
	return val
}

func getReferenceValue(value reflect.Value) interface{} {
	v := flattenReflectValue(value)
	if v.Kind() == reflect.Invalid {
		return nil
	}
	if v.CanInterface() {
		return v.Interface()
	}
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint()
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Bool:
		return v.Bool()
	case reflect.String:
		return v.String()
	case reflect.Array, reflect.Slice:
		var arr = []interface{}{}
		for i := 0; i < v.Len(); i++ {
			el := getReferenceValue(v.Index(i))
			arr = append(arr, el)
		}
		return arr
	default:
		return nil
	}
}

func makeNormalWrappedKeyedValue(value interface{}, parent *wrappedKeyedValue) *wrappedKeyedValue {
	newWrappedKeyedValue := &wrappedKeyedValue{
		value:              value,
		shouldValidateNorm: true,
		shouldValidateAny:  false,
		shouldValidateAll:  false,
		children:           []*wrappedKeyedValue{},
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
	parent.value = nil
	key := keys[keyIndex]
	keyedValue := getValueForKey(key, value)
	if keyedValue == nil {
		return
	}
	switch key {
	case keyAny:
		parent.shouldValidateAny = true
		parent.shouldValidateAll = false
		parent.shouldValidateNorm = false

		arrValue := reflect.ValueOf(keyedValue)
		for i := 0; i < arrValue.Len(); i++ {
			el := getReferenceValue(arrValue.Index(i))
			newWrappedKeyedValue := makeNormalWrappedKeyedValue(el, parent)
			buildWrappedKeyValueWithKeys(keys, keyIndex+1, el, newWrappedKeyedValue)
		}
	case keyAll:
		parent.shouldValidateAny = false
		parent.shouldValidateAll = true
		parent.shouldValidateNorm = false

		arrValue := reflect.ValueOf(keyedValue)
		for i := 0; i < arrValue.Len(); i++ {
			el := getReferenceValue(arrValue.Index(i))
			newWrappedKeyedValue := makeNormalWrappedKeyedValue(el, parent)
			buildWrappedKeyValueWithKeys(keys, keyIndex+1, el, newWrappedKeyedValue)
		}
	default:
		newWrappedKeyedValue := makeNormalWrappedKeyedValue(keyedValue, parent)
		buildWrappedKeyValueWithKeys(keys, keyIndex+1, keyedValue, newWrappedKeyedValue)
	}
}
