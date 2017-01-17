package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i) // imaginary part
	sqrt              = cmplx.Sqrt(2)        // is return and initialized
	f32    float32    = 0.04
	f64    float64    = 0.04
)

func main() {

	const f = "%T(%v) (%#v) \n" // All verbs: https://golang.org/pkg/fmt/
	fmt.Printf(f, ToBe, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt, MaxInt)
	fmt.Printf(f, z, z, z)
	fmt.Println(z)
	fmt.Println(f32)
	fmt.Println(f64)
	fmt.Println(sqrt)

	/*
	   uint8       the set of all unsigned  8-bit integers (0 to 255)
	   uint16      the set of all unsigned 16-bit integers (0 to 65535)
	   uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	   uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	   int8        the set of all signed  8-bit integers (-128 to 127)
	   int16       the set of all signed 16-bit integers (-32768 to 32767)
	   int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	   int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	   float32     the set of all IEEE-754 32-bit floating-point numbers
	   float64     the set of all IEEE-754 64-bit floating-point numbers

	   complex64   the set of all complex numbers with float32 real and imaginary parts
	   complex128  the set of all complex numbers with float64 real and imaginary parts

	   byte        alias for uint8
	   rune        alias for int32
	*/

}