// Code generated by "go run generate_coercers.go"; DO NOT EDIT.
package venom

import "fmt"

// A CoerceErr is returned when incompatible types are attempted to be coerced
type CoerceErr struct {
	From interface{}
	To   string
}

func (e *CoerceErr) Error() string {
	return fmt.Sprintf("venom: can not coerce %T to %q", e.From, e.To)
}

func coerceString(val interface{}) (string, error) {
	if value, ok := val.(string); !ok {
		return "", &CoerceErr{val, "string"}
	} else {
		return value, nil
	}
}

func coerceStringSlice(val interface{}) ([]string, error) {
	if value, ok := val.([]string); !ok {
		return nil, &CoerceErr{val, "[]string"}
	} else {
		return value, nil
	}
}

func coerceBool(val interface{}) (bool, error) {
	if value, ok := val.(bool); !ok {
		return false, &CoerceErr{val, "bool"}
	} else {
		return value, nil
	}
}

func coerceBoolSlice(val interface{}) ([]bool, error) {
	if value, ok := val.([]bool); !ok {
		return nil, &CoerceErr{val, "[]bool"}
	} else {
		return value, nil
	}
}

func coerceInt(val interface{}) (int, error) {
	if value, ok := val.(int); !ok {
		return 0, &CoerceErr{val, "int"}
	} else {
		return value, nil
	}
}

func coerceIntSlice(val interface{}) ([]int, error) {
	if value, ok := val.([]int); !ok {
		return nil, &CoerceErr{val, "[]int"}
	} else {
		return value, nil
	}
}

func coerceInt8(val interface{}) (int8, error) {
	if value, ok := val.(int8); !ok {
		return 0, &CoerceErr{val, "int8"}
	} else {
		return value, nil
	}
}

func coerceInt8Slice(val interface{}) ([]int8, error) {
	if value, ok := val.([]int8); !ok {
		return nil, &CoerceErr{val, "[]int8"}
	} else {
		return value, nil
	}
}

func coerceInt16(val interface{}) (int16, error) {
	if value, ok := val.(int16); !ok {
		return 0, &CoerceErr{val, "int16"}
	} else {
		return value, nil
	}
}

func coerceInt16Slice(val interface{}) ([]int16, error) {
	if value, ok := val.([]int16); !ok {
		return nil, &CoerceErr{val, "[]int16"}
	} else {
		return value, nil
	}
}

func coerceInt32(val interface{}) (int32, error) {
	if value, ok := val.(int32); !ok {
		return 0, &CoerceErr{val, "int32"}
	} else {
		return value, nil
	}
}

func coerceInt32Slice(val interface{}) ([]int32, error) {
	if value, ok := val.([]int32); !ok {
		return nil, &CoerceErr{val, "[]int32"}
	} else {
		return value, nil
	}
}

func coerceInt64(val interface{}) (int64, error) {
	if value, ok := val.(int64); !ok {
		return 0, &CoerceErr{val, "int64"}
	} else {
		return value, nil
	}
}

func coerceInt64Slice(val interface{}) ([]int64, error) {
	if value, ok := val.([]int64); !ok {
		return nil, &CoerceErr{val, "[]int64"}
	} else {
		return value, nil
	}
}

func coerceUint(val interface{}) (uint, error) {
	if value, ok := val.(uint); !ok {
		return 0, &CoerceErr{val, "uint"}
	} else {
		return value, nil
	}
}

func coerceUintSlice(val interface{}) ([]uint, error) {
	if value, ok := val.([]uint); !ok {
		return nil, &CoerceErr{val, "[]uint"}
	} else {
		return value, nil
	}
}

func coerceUint8(val interface{}) (uint8, error) {
	if value, ok := val.(uint8); !ok {
		return 0, &CoerceErr{val, "uint8"}
	} else {
		return value, nil
	}
}

func coerceUint8Slice(val interface{}) ([]uint8, error) {
	if value, ok := val.([]uint8); !ok {
		return nil, &CoerceErr{val, "[]uint8"}
	} else {
		return value, nil
	}
}

func coerceUint16(val interface{}) (uint16, error) {
	if value, ok := val.(uint16); !ok {
		return 0, &CoerceErr{val, "uint16"}
	} else {
		return value, nil
	}
}

func coerceUint16Slice(val interface{}) ([]uint16, error) {
	if value, ok := val.([]uint16); !ok {
		return nil, &CoerceErr{val, "[]uint16"}
	} else {
		return value, nil
	}
}

func coerceUint32(val interface{}) (uint32, error) {
	if value, ok := val.(uint32); !ok {
		return 0, &CoerceErr{val, "uint32"}
	} else {
		return value, nil
	}
}

func coerceUint32Slice(val interface{}) ([]uint32, error) {
	if value, ok := val.([]uint32); !ok {
		return nil, &CoerceErr{val, "[]uint32"}
	} else {
		return value, nil
	}
}

func coerceUint64(val interface{}) (uint64, error) {
	if value, ok := val.(uint64); !ok {
		return 0, &CoerceErr{val, "uint64"}
	} else {
		return value, nil
	}
}

func coerceUint64Slice(val interface{}) ([]uint64, error) {
	if value, ok := val.([]uint64); !ok {
		return nil, &CoerceErr{val, "[]uint64"}
	} else {
		return value, nil
	}
}

func coerceFloat32(val interface{}) (float32, error) {
	if value, ok := val.(float32); !ok {
		return 0, &CoerceErr{val, "float32"}
	} else {
		return value, nil
	}
}

func coerceFloat32Slice(val interface{}) ([]float32, error) {
	if value, ok := val.([]float32); !ok {
		return nil, &CoerceErr{val, "[]float32"}
	} else {
		return value, nil
	}
}

func coerceFloat64(val interface{}) (float64, error) {
	if value, ok := val.(float64); !ok {
		return 0, &CoerceErr{val, "float64"}
	} else {
		return value, nil
	}
}

func coerceFloat64Slice(val interface{}) ([]float64, error) {
	if value, ok := val.([]float64); !ok {
		return nil, &CoerceErr{val, "[]float64"}
	} else {
		return value, nil
	}
}