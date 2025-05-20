package generics_test

import "testing"

type MyInt int
type Number interface {
	~int | ~float64
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SomeFunc[T Number](v T) {
	// ...
}
func TestGenerics(t *testing.T) {
	t.Logf("sum of ints : %v", SumIntsOrFloats(map[string]int64{"a": 1, "b": 2}))
	t.Logf("sum of float: %v", SumIntsOrFloats(map[string]float64{"a": 1.1, "b": 2.2}))

	SomeFunc(MyInt(1))
}
