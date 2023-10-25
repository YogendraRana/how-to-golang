package main

import (
	"fmt"
	"math"
)

func main() {
	// Boolean
	var b bool = true
	fmt.Printf("Boolean: %v\n", b)

	// Integer
	var i8 int8 = math.MaxInt8
	fmt.Printf("Int8: %v\n", i8)
	var i16 int16 = math.MaxInt16
	fmt.Printf("Int16: %v\n", i16)
	var i32 int32 = math.MaxInt32
	fmt.Printf("Int32: %v\n", i32)
	var i64 int64 = math.MaxInt64
	fmt.Printf("Int64: %v\n", i64)

	var ui8 uint8 = math.MaxUint8
	fmt.Printf("Uint8: %v\n", ui8)
	var ui16 uint16 = math.MaxUint16
	fmt.Printf("Uint16: %v\n", ui16)
	var ui32 uint32 = math.MaxUint32
	fmt.Printf("Uint32: %v\n", ui32)
	var ui64 uint64 = math.MaxUint64
	fmt.Printf("Uint64: %v\n", ui64)

	// Floating Point
	var f32 float32 = math.MaxFloat32
	fmt.Printf("Float32: %v\n", f32)
	var f64 float64 = math.MaxFloat64
	fmt.Printf("Float64: %v\n", f64)

	// Complex
	var c64 complex64 = complex(math.MaxFloat32, math.MaxFloat32)
	fmt.Printf("Complex64: %v\n", c64)
	var c128 complex128 = complex(math.MaxFloat64, math.MaxFloat64)
	fmt.Printf("Complex128: %v\n", c128)

	// String
	// with double quotes i cannot print the string in multiple lines
	// with backticks i can print the string in multiple lines
	var s string = "Hello, World!"
	fmt.Printf("String: %v\n", s)
}
