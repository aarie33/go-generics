package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface {
	int | int64 | float64 | float32
}](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin(100, 200))
	assert.Equal(t, 100.0, FindMin(100.0, 200.0))
	assert.Equal(t, int64(100), FindMin[int64](100, 200))
}

// type parameter dengan tipe generic
func GetFirst[T []E, E any](s T) E {
	return s[0]
}

func TestGetFirst(t *testing.T) {
	names := []string{"Joko", "Budi", "Susi"}

	first := GetFirst(names)
	assert.Equal(t, "Joko", first)
}
