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
		i64t   int64   = math.MaxInt64
		i32t   int32   = math.MaxInt32
		i16t   int16   = math.MaxInt16
		i8t    int8    = math.MaxInt8
		ui64t  uint64  = math.MaxUint64
		ui32t  uint32  = math.MaxUint32
		ui16t  uint16  = math.MaxUint16
		ui8t   uint8   = math.MaxUint8
		st     string  = ""
		pf64t  *float64
		pf32t  *float32
		pi64t  *int64
		pi32t  *int32
		pi16t  *int16
		pi8t   *int8
		pui64t *uint64
		pui32t *uint32
		pui16t *uint16
		pui8t  *uint8
		pst    *string
	)
	testTrues := []interface{}{
		f64t,
		f32t,
		i64t,
		i32t,
		i16t,
		i8t,
		ui64t,
		ui32t,
		ui16t,
		ui8t,
		st,
		&f64t,
		&f32t,
		&i64t,
		&i32t,
		&i16t,
		&i8t,
		&ui64t,
		&ui32t,
		&ui16t,
		&ui8t,
		&st,
		nil,
		pf64t,
		pf32t,
		pi64t,
		pi32t,
		pi16t,
		pi8t,
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
		f64f  float64 = 0
		f32f  float32 = 0
		i64f  int64   = 0
		i32f  int32   = 0
		i16f  int16   = 0
		i8f   int8    = 0
		ui64f uint64  = 0
		ui32f uint32  = 0
		ui16f uint16  = 0
		ui8f  uint8   = 0
		sf    string  = "Not Null"
	)
	testFalses := []interface{}{
		f64f,
		f32f,
		i64f,
		i32f,
		i16f,
		i8f,
		ui64f,
		ui32f,
		ui16f,
		ui8f,
		sf,
		&f64f,
		&f32f,
		&i64f,
		&i32f,
		&i16f,
		&i8f,
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
