package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](g GetterSetter[T], value T) T {
	g.SetValue(value)
	return g.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (d *MyData[T]) GetValue() T {
	return d.Value
}

func (d *MyData[T]) SetValue(value T) {
	d.Value = value
}

func TestInterface(t *testing.T) {
	myData := MyData[string]{}
	result := ChangeValue[string](&myData, "Hello")

	assert.Equal(t, "Hello", result)
	assert.Equal(t, "Hello", myData.Value)
}
