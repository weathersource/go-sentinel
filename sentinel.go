package sentinel

import "math"

// Constant sentinel values. Excludes float64 and float32 values, as NaN is not a constant.
const (
	Int    int    = math.MaxInt
	Int64  int64  = math.MaxInt64
	Int32  int32  = math.MaxInt32
	Int16  int16  = math.MaxInt16
	Int8   int8   = math.MaxInt8
	Uint   uint   = math.MaxUint
	Uint64 uint64 = math.MaxUint64
	Uint32 uint32 = math.MaxUint32
	Uint16 uint16 = math.MaxUint16
	Uint8  uint8  = math.MaxUint8
	String string = ""
)

// NullFloat64 returns the NULL sentinel value for type float64.
func NullFloat64() float64 {
	return math.NaN()
}

// NullFloat32 returns the NULL sentinel value for type float32.
func NullFloat32() float32 {
	return float32(math.NaN())
}

// NullInt64 returns the NULL sentinel value for type int64.
func NullInt64() int64 {
	return Int64
}
// NullInt returns the NULL sentinel value for type int.
func NullInt() int {
	return Int
}

// NullInt32 returns the NULL sentinel value for type int32.
func NullInt32() int32 {
	return Int32
}

// NullInt16 returns the NULL sentinel value for type int16.
func NullInt16() int16 {
	return Int16
}

// NullInt8 returns the NULL sentinel value for type int8.
func NullInt8() int8 {
	return Int8
}

// NullUint returns the NULL sentinel value for type uint.
func NullUint() uint {
	return Uint
}

// NullUint64 returns the NULL sentinel value for type uint64.
func NullUint64() uint64 {
	return Uint64
}

// NullUint32 returns the NULL sentinel value for type uint32.
func NullUint32() uint32 {
	return Uint32
}

// NullUint16 returns the NULL sentinel value for type uint16.
func NullUint16() uint16 {
	return Uint16
}

// NullUint8 returns the NULL sentinel value for type uint8.
func NullUint8() uint8 {
	return Uint8
}

// NullString returns the NULL sentinel value for type string.
func NullString() string {
	return String
}

// IsNull tests input i for equality with the NULL sentinel value for the type of i.
func IsNull(i interface{}) bool {

	switch t := i.(type) {
	case string:
		return t == String
	case float64:
		return math.IsNaN(t)
	case float32:
		return math.IsNaN(float64(t))
	case int:
		return t == Int
	case int64:
		return t == Int64
	case int32:
		return t == Int32
	case int16:
		return t == Int16
	case int8:
		return t == Int8
	case uint:
		return t == Uint
	case uint64:
		return t == Uint64
	case uint32:
		return t == Uint32
	case uint16:
		return t == Uint16
	case uint8:
		return t == Uint8
	case *string:
		return t == nil || *t == String
	case *float64:
		return t == nil || math.IsNaN(*t)
	case *float32:
		return t == nil || math.IsNaN(float64(*t))
	case *int:
		return t == nil || *t == Int
	case *int64:
		return t == nil || *t == Int64
	case *int32:
		return t == nil || *t == Int32
	case *int16:
		return t == nil || *t == Int16
	case *int8:
		return t == nil || *t == Int8
	case *uint:
		return t == nil || *t == Uint
	case *uint64:
		return t == nil || *t == Uint64
	case *uint32:
		return t == nil || *t == Uint32
	case *uint16:
		return t == nil || *t == Uint16
	case *uint8:
		return t == nil || *t == Uint8
	case nil:
		return true
	}

	return false
}
