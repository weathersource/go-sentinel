package sentinel

import (
	"math"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestNullFloat64(t *testing.T) {
	assert.True(t, math.IsNaN(NullFloat64()))
}

func TestNullFloat32(t *testing.T) {
	assert.True(t, math.IsNaN(float64(NullFloat32())))
}

func TestNullInt(t *testing.T) {
	assert.Equal(t, int(math.MaxInt), NullInt())
}

func TestNullInt64(t *testing.T) {
	assert.Equal(t, int64(math.MaxInt64), NullInt64())
}

func TestNullInt32(t *testing.T) {
	assert.Equal(t, int32(math.MaxInt32), NullInt32())
}

func TestNullInt16(t *testing.T) {
	assert.Equal(t, int16(math.MaxInt16), NullInt16())
}

func TestNullInt8(t *testing.T) {
	assert.Equal(t, int8(math.MaxInt8), NullInt8())
}

func TestNullUint64(t *testing.T) {
	assert.Equal(t, uint64(math.MaxUint64), NullUint64())
}

func TestNullUint(t *testing.T) {
	assert.Equal(t, uint(math.MaxUint), NullUint())
}

func TestNullUint32(t *testing.T) {
	assert.Equal(t, uint32(math.MaxUint32), NullUint32())
}

func TestNullUint16(t *testing.T) {
	assert.Equal(t, uint16(math.MaxUint16), NullUint16())
}

func TestNullUint8(t *testing.T) {
	assert.Equal(t, uint8(math.MaxUint8), NullUint8())
}

func TestNullString(t *testing.T) {
	assert.Equal(t, "", NullString())
}

func TestIsNull(t *testing.T) {

	var (
		f64t   float64 = math.NaN()
		f32t   float32 = float32(math.NaN())
		it     int64   = math.MaxInt
		i64t   int64   = math.MaxInt64
		i32t   int32   = math.MaxInt32
		i16t   int16   = math.MaxInt16
		i8t    int8    = math.MaxInt8
		uit    uint64  = math.MaxUint
		ui64t  uint64  = math.MaxUint64
		ui32t  uint32  = math.MaxUint32
		ui16t  uint16  = math.MaxUint16
		ui8t   uint8   = math.MaxUint8
		st     string
		pf64t  *float64
		pf32t  *float32
		pit    *int
		pi64t  *int64
		pi32t  *int32
		pi16t  *int16
		pi8t   *int8
		puit   *uint
		pui64t *uint64
		pui32t *uint32
		pui16t *uint16
		pui8t  *uint8
		pst    *string
	)
	testTrues := []interface{}{
		f64t,
		f32t,
		it,
		i64t,
		i32t,
		i16t,
		i8t,
		uit,
		ui64t,
		ui32t,
		ui16t,
		ui8t,
		st,
		&f64t,
		&f32t,
		&it,
		&i64t,
		&i32t,
		&i16t,
		&i8t,
		&uit,
		&ui64t,
		&ui32t,
		&ui16t,
		&ui8t,
		&st,
		nil,
		pf64t,
		pf32t,
		pit,
		pi64t,
		pi32t,
		pi16t,
		pi8t,
		puit,
		pui64t,
		pui32t,
		pui16t,
		pui8t,
		pst,
	}
	for _, test := range testTrues {
		assert.True(t, IsNull(test))
	}

	var (
		f64f  float64
		f32f  float32
		iif   int
		i64f  int64
		i32f  int32
		i16f  int16
		i8f   int8
		uif   uint
		ui64f uint64
		ui32f uint32
		ui16f uint16
		ui8f  uint8
		sf    string = "Not Null"
	)
	testFalses := []interface{}{
		f64f,
		f32f,
		iif,
		i64f,
		i32f,
		i16f,
		i8f,
		uif,
		ui64f,
		ui32f,
		ui16f,
		ui8f,
		sf,
		&f64f,
		&f32f,
		&iif,
		&i64f,
		&i32f,
		&i16f,
		&i8f,
		&uif,
		&ui64f,
		&ui32f,
		&ui16f,
		&ui8f,
		&sf,
		true,
	}
	for _, test := range testFalses {
		assert.False(t, IsNull(test))
	}
}
