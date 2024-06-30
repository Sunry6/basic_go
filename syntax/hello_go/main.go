package main

func main() {
	var a int = 123
	var b int = 456
	println(a + b)
	println(a - b)
	println(a * b)
	println(a / b)

	// 同类型才能加减乘除
	var c float64 = 12.3
	println(a + int(c))
}
