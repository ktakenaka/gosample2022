package main

import "fmt"

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Println("traditional")
	fmt.Println("sum ints", SumInts(ints))
	fmt.Println("sum floats", SumFloats(floats))

	fmt.Println("generics")
	fmt.Println("sum ints", SumIntsOrFloats[string, int64](ints))
	fmt.Println("sum floats", SumIntsOrFloats[string, float64](floats))

	fmt.Println("sum numbers")
	fmt.Println(SumNumbers(ints))
	fmt.Println(SumNumbers(floats))
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
