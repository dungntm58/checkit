package checkit

// Validator ...
type Validator map[string]Validating

// ValidateSync ...
func ValidateSync(any interface{}, validator Validator) (bool, error) {
	return validator.ValidateSync(any)
}

// ValidateSync ...
func (v Validator) ValidateSync(any interface{}) (bool, error) {
	for keyPath, validating := range v {
		_test, _isResultAny, _isResultAll := getValueWithKeyPath(keyPath, any)
		if _isResultAll {
			_testArr, _ := _test.([]interface{})
			for _, _el := range _testArr {
				_r, _err := validating.Validate(_el)
				if _err != nil {
					return false, _err
				}
				if !_r {
					return false, nil
				}
			}
		} else if _isResultAny {
			_testArr, _ := _test.([]interface{})
			var _resultAny bool = true
			for _, _el := range _testArr {
				_r, _err := validating.Validate(_el)
				if _err != nil {
					return false, _err
				}
				_resultAny = _resultAny || _r
			}
			if !_resultAny {
				return false, nil
			}
		} else {
			_result, _err := validating.Validate(_test)
			if _err != nil {
				return false, _err
			}
			return _result, nil
		}
	}
	return true, nil
}
