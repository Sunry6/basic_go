package main

import "fmt"

func main() {
	// ClosureInvoke()
	// Defer()
	// DeferLoop(5)
	// DeferClosure()
	// DeferClosureV1()
	// fmt.Println(DeferReturn())
	// fmt.Println(DeferReturnV1())
	fmt.Println(DeferReturnV2().name)
}

func Invoke() {
	str := Func0("大明")
	println(str)
	str1, err := Func2(12, 13)
	println(str1, err)
}

func Func0(name string) string {
	return "hello, " + name
}

func Func1(a, b, c int, d string) (string, error) {
	return "hello, world", nil
}

func Func2(a int, b int) (str string, err error) {
	str = "hello, world"
	return
}

func Func3(a int, b int) (str string, err error) {
	//str = "def"
	return "abc", nil
}

func Recursive() {

}

func A() {
	B()
}

func B() {
	C()
}

func C() {
	A()
}
