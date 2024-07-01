package main

import "fmt"

var i = 0

func Closure(name string) func() string {
	return func() string {
		i++
		world := "world"
		fmt.Println(i)
		return "hello, " + name + world
	}
}

func ClosureInvoke() {
	c := Closure("666")
	println(c())
}
