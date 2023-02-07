package gogenerics

import "testing"

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, item := range bag {
		println(item)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"John", "Paul", "George", "Ringo"}
	PrintBag(names)
}

func TestBagInt(t *testing.T) {
	numbers := Bag[int]{1, 2, 3, 4}
	PrintBag(numbers)
}
