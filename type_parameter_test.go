package gogenerics

import (
	"fmt"
	"testing"
)

// func Length[T interface{}]() { lebih simple pakai any
func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var result string = Length("Mandor")
	fmt.Println(result)

	var resultNumber int = Length(10)
	fmt.Println(resultNumber)
}

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) (T1, T2) {
	return param1, param2
}

func TestMultipleParameter(t *testing.T) {
	var resultString, resultNumber = MultipleParameter("Mandor", 10)
	fmt.Println(resultString, resultNumber)
}
