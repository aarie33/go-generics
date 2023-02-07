package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func TestMin(t *testing.T) {
	assert.Equal(t, Min[int](2, 1), 1)
	assert.Equal(t, Min(2, 1), 1)
	assert.Equal(t, Min(1.0, 2.0), 1.0)
	assert.Equal(t, Min(2.0, 1.0), 1.0)

	// Age is not a Number even though it is an int

	// Use type aproximation to make it work with ~int
	assert.Equal(t, Min(Age(1), Age(2)), Age(1))
}

func TestMinTypeInference(t *testing.T) {
	assert.Equal(t, Min(100, 200), 100)
	assert.Equal(t, Min(int64(100), int64(200)), int64(100))
	assert.Equal(t, Min(float32(100), float32(200)), float32(100))
}
