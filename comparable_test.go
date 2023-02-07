package gogenerics

import (
	"fmt"
	"testing"
)

func IsSame[T comparable](param1 T, param2 T) bool {
	return param1 == param2
}

func TestIsSame(t *testing.T) {
	var result = IsSame("Mandor", "Mandora")
	fmt.Println(result)

	var resultNumber = IsSame(10, 10)
	fmt.Println(resultNumber)
}
