package gogenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[T]) SayHello(name string) string {
	return "Hello" + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{First: "first", Second: "second"}
	fmt.Println(data.First, data.Second)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{First: "first", Second: "second"}

	assert.Equal(t, "Tono", data.ChangeFirst("Tono"))
	assert.Equal(t, "Hello Tono", data.SayHello(" Tono"))
}
