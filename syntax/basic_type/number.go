package basic_type

import "math"

func Number() {
	var a int = 456
	var b int = 123
	println(a + b)
	println(a - b)
	println(a * b)
	println(a / b)
}

func Extremum() {
	println("float64 max", math.MaxFloat64)
	println("float64 min", math.SmallestNonzeroFloat64)
	println("float32 max", math.MaxFloat32)
	println("float32 min", math.SmallestNonzeroFloat32)
}
