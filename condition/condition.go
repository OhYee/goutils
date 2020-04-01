// Code generated by GCG. DO NOT EDIT.
// Go Code Generator 0.0.8 (https://github.com/OhYee/gcg)

package condition

import ()

// IfString return value according condition
func IfString(condition bool, t string, f string) string {
	if condition {
		return t
	}
	return f
}

// IfInt return value according condition
func IfInt(condition bool, t int, f int) int {
	if condition {
		return t
	}
	return f
}

// IfInt8 return value according condition
func IfInt8(condition bool, t int8, f int8) int8 {
	if condition {
		return t
	}
	return f
}

// IfInt16 return value according condition
func IfInt16(condition bool, t int16, f int16) int16 {
	if condition {
		return t
	}
	return f
}

// IfInt32 return value according condition
func IfInt32(condition bool, t int32, f int32) int32 {
	if condition {
		return t
	}
	return f
}

// IfInt64 return value according condition
func IfInt64(condition bool, t int64, f int64) int64 {
	if condition {
		return t
	}
	return f
}

// IfUint8 return value according condition
func IfUint8(condition bool, t uint8, f uint8) uint8 {
	if condition {
		return t
	}
	return f
}

// IfUint16 return value according condition
func IfUint16(condition bool, t uint16, f uint16) uint16 {
	if condition {
		return t
	}
	return f
}

// IfUint32 return value according condition
func IfUint32(condition bool, t uint32, f uint32) uint32 {
	if condition {
		return t
	}
	return f
}

// IfUint64 return value according condition
func IfUint64(condition bool, t uint64, f uint64) uint64 {
	if condition {
		return t
	}
	return f
}

// IfFloat32 return value according condition
func IfFloat32(condition bool, t float32, f float32) float32 {
	if condition {
		return t
	}
	return f
}

// IfFloat64 return value according condition
func IfFloat64(condition bool, t float64, f float64) float64 {
	if condition {
		return t
	}
	return f
}

// IfByte return value according condition
func IfByte(condition bool, t byte, f byte) byte {
	if condition {
		return t
	}
	return f
}

// IfBool return value according condition
func IfBool(condition bool, t bool, f bool) bool {
	if condition {
		return t
	}
	return f
}

// IfRune return value according condition
func IfRune(condition bool, t rune, f rune) rune {
	if condition {
		return t
	}
	return f
}

// If return value according condition
func If(condition bool, t any, f any) any {
	if condition {
		return t
	}
	return f
}
