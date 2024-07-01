package main

func Defer() {
	defer func() {
		println("first defer")
	}()

	defer func() {
		println("second defer")
	}()
}

func DeferLoop(max int) {
	for i := 0; i < max; i++ {
		defer func() {
			println("hello", i)
		}()
	}
}

func DeferClosure() {
	j := 0
	defer func() {
		println(j)
	}()
	j = 1
}

func DeferClosureV1() {
	j := 0
	defer func(j int) {
		println(j)
	}(j)
	j = 1
}

func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturnV2() *MyStruct {
	res := &MyStruct{
		name: "Tom",
	}
	defer func() {
		res.name = "Jerry"
	}()

	return res
}

type MyStruct struct {
	name string
}
