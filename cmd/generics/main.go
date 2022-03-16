package main

import "fmt"

// Map turns a []T1 to a []T2 using a mapping function.
// This function has two type parameters, T1 and T2.
// This works with slices of any type.
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#Examples
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Reduce reduces a []T1 to a single value using a reduction function.
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#Examples
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filter filters values from a slice using a filter function.
// It returns a new slice with only the elements of s
// for which f returned true.
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#Examples
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

type Number interface {
	Integer | Float
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	~int |
		~int8 |
		~int16 |
		~int32 |
		~int64 |
		~uint |
		~uint8 |
		~uint16 |
		~uint32 |
		~uint64
}

func double[T Number](n T) T {
	return 2 * n
}

func add[T Number](m, n T) T {
	return m + n
}

func isOdd[T Integer](n T) bool {
	return n%2 != 0
}

type MyNumber int

func main() {
	fmt.Println(Map([]MyNumber{1, 2, 3}, double[MyNumber]))
	fmt.Println(Reduce([]int{1, 2, 3}, 0, add[int]))
	fmt.Println(Filter([]int{1, 2, 3}, isOdd[int]))
}
