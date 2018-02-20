// Code generated by "go run generate_getters.go"; DO NOT EDIT.
package venom

// GetBool attempts to cast the returned config value from the global
// Venom instance
func GetBool(key string) bool {
	return v.GetBool(key)
}

// GetBool attempts to cast the returned config value from the current
// Venom instance.
//
// If the key does not exist, or if the type contained in Venom can
// not be cast to the requested type, then the zero value of false
// will be returned.
func (v *Venom) GetBool(key string) bool {
	val, ok := v.Find(key)
	if !ok {
		return false
	}

	if value, ok := val.(bool); !ok {
		return false
	} else {
		return value
	}
}

// GetFloat32 attempts to cast the returned config value from the global
// Venom instance
func GetFloat32(key string) float32 {
	return v.GetFloat32(key)
}

// GetFloat32 attempts to cast the returned config value from the current
// Venom instance.
//
// If the key does not exist, or if the type contained in Venom can
// not be cast to the requested type, then the zero value of 0
// will be returned.
func (v *Venom) GetFloat32(key string) float32 {
	val, ok := v.Find(key)
	if !ok {
		return 0
	}

	if value, ok := val.(float32); !ok {
		return 0
	} else {
		return value
	}
}

// GetFloat64 attempts to cast the returned config value from the global
// Venom instance
func GetFloat64(key string) float64 {
	return v.GetFloat64(key)
}

// GetFloat64 attempts to cast the returned config value from the current
// Venom instance.
//
// If the key does not exist, or if the type contained in Venom can
// not be cast to the requested type, then the zero value of 0
// will be returned.
func (v *Venom) GetFloat64(key string) float64 {
	val, ok := v.Find(key)
	if !ok {
		return 0
	}

	if value, ok := val.(float64); !ok {
		return 0
	} else {
		return value
	}
}

// GetInt attempts to cast the returned config value from the global
// Venom instance
func GetInt(key string) int {
	return v.GetInt(key)
}

// GetInt attempts to cast the returned config value from the current
// Venom instance.
//
// If the key does not exist, or if the type contained in Venom can
// not be cast to the requested type, then the zero value of 0
// will be returned.
func (v *Venom) GetInt(key string) int {
	val, ok := v.Find(key)
	if !ok {
		return 0
	}

	if value, ok := val.(int); !ok {
		return 0
	} else {
		return value
	}
}

// GetInt8 attempts to cast the returned config value from the global
// Venom instance
func GetInt8(key string) int8 {
	return v.GetInt8(key)
}

// GetInt8 attempts to cast the returned config value from the current
// Venom instance.
//
// If the key does not exist, or if the type contained in Venom can
// not be cast to the requested type, then the zero value of 0
// will be returned.
func (v *Venom) GetInt8(key string) int8 {
	val, ok := v.Find(key)
	if !ok {
		return 0
	}

	if value, ok := val.(int8); !ok {
		return 0
	} else {
		return value
	}
}

// GetInt16 attempts to cast the returned config value from the global
// Venom instance
func GetInt16(key string) int16 {
	return v.GetInt16(key)
}

// GetInt16 attempts to cast the returned config value from the current
// Venom instance.
//
// If the key does not exist, or if the type contained in Venom can
// not be cast to the requested type, then the zero value of 0
// will be returned.
func (v *Venom) GetInt16(key string) int16 {
	val, ok := v.Find(key)
	if !ok {
		return 0
	}

	if value, ok := val.(int16); !ok {
		return 0
	} else {
		return value
	}
}

// GetInt32 attempts to cast the returned config value from the global
// Venom instance
func GetInt32(key string) int32 {
	return v.GetInt32(key)
}

// GetInt32 attempts to cast the returned config value from the current
// Venom instance.
//
// If the key does not exist, or if the type contained in Venom can
// not be cast to the requested type, then the zero value of 0
// will be returned.
func (v *Venom) GetInt32(key string) int32 {
	val, ok := v.Find(key)
	if !ok {
		return 0
	}

	if value, ok := val.(int32); !ok {
		return 0
	} else {
		return value
	}
}

// GetInt64 attempts to cast the returned config value from the global
// Venom instance
func GetInt64(key string) int64 {
	return v.GetInt64(key)
}

// GetInt64 attempts to cast the returned config value from the current
// Venom instance.
//
// If the key does not exist, or if the type contained in Venom can
// not be cast to the requested type, then the zero value of 0
// will be returned.
func (v *Venom) GetInt64(key string) int64 {
	val, ok := v.Find(key)
	if !ok {
		return 0
	}

	if value, ok := val.(int64); !ok {
		return 0
	} else {
		return value
	}
}

// GetString attempts to cast the returned config value from the global
// Venom instance
func GetString(key string) string {
	return v.GetString(key)
}

// GetString attempts to cast the returned config value from the current
// Venom instance.
//
// If the key does not exist, or if the type contained in Venom can
// not be cast to the requested type, then the zero value of ""
// will be returned.
func (v *Venom) GetString(key string) string {
	val, ok := v.Find(key)
	if !ok {
		return ""
	}

	if value, ok := val.(string); !ok {
		return ""
	} else {
		return value
	}
}

// GetUint attempts to cast the returned config value from the global
// Venom instance
func GetUint(key string) uint {
	return v.GetUint(key)
}

// GetUint attempts to cast the returned config value from the current
// Venom instance.
//
// If the key does not exist, or if the type contained in Venom can
// not be cast to the requested type, then the zero value of 0
// will be returned.
func (v *Venom) GetUint(key string) uint {
	val, ok := v.Find(key)
	if !ok {
		return 0
	}

	if value, ok := val.(uint); !ok {
		return 0
	} else {
		return value
	}
}

// GetUint8 attempts to cast the returned config value from the global
// Venom instance
func GetUint8(key string) uint8 {
	return v.GetUint8(key)
}

// GetUint8 attempts to cast the returned config value from the current
// Venom instance.
//
// If the key does not exist, or if the type contained in Venom can
// not be cast to the requested type, then the zero value of 0
// will be returned.
func (v *Venom) GetUint8(key string) uint8 {
	val, ok := v.Find(key)
	if !ok {
		return 0
	}

	if value, ok := val.(uint8); !ok {
		return 0
	} else {
		return value
	}
}

// GetUint16 attempts to cast the returned config value from the global
// Venom instance
func GetUint16(key string) uint16 {
	return v.GetUint16(key)
}

// GetUint16 attempts to cast the returned config value from the current
// Venom instance.
//
// If the key does not exist, or if the type contained in Venom can
// not be cast to the requested type, then the zero value of 0
// will be returned.
func (v *Venom) GetUint16(key string) uint16 {
	val, ok := v.Find(key)
	if !ok {
		return 0
	}

	if value, ok := val.(uint16); !ok {
		return 0
	} else {
		return value
	}
}

// GetUint32 attempts to cast the returned config value from the global
// Venom instance
func GetUint32(key string) uint32 {
	return v.GetUint32(key)
}

// GetUint32 attempts to cast the returned config value from the current
// Venom instance.
//
// If the key does not exist, or if the type contained in Venom can
// not be cast to the requested type, then the zero value of 0
// will be returned.
func (v *Venom) GetUint32(key string) uint32 {
	val, ok := v.Find(key)
	if !ok {
		return 0
	}

	if value, ok := val.(uint32); !ok {
		return 0
	} else {
		return value
	}
}

// GetUint64 attempts to cast the returned config value from the global
// Venom instance
func GetUint64(key string) uint64 {
	return v.GetUint64(key)
}

// GetUint64 attempts to cast the returned config value from the current
// Venom instance.
//
// If the key does not exist, or if the type contained in Venom can
// not be cast to the requested type, then the zero value of 0
// will be returned.
func (v *Venom) GetUint64(key string) uint64 {
	val, ok := v.Find(key)
	if !ok {
		return 0
	}

	if value, ok := val.(uint64); !ok {
		return 0
	} else {
		return value
	}
}
