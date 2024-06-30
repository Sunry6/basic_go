package basic_type

import "unicode/utf8"

func String() {
	println("Hello Go")
	println("He said: \"Hello Go\"")
	println(`我可以换行
	这是新的行
	但是这里不能有反引号`)
	println("hello, " + "go!")
	println(len("你好"))
	println(utf8.RuneCountInString("你好"))
	println(utf8.RuneCountInString("你好cba"))
}
