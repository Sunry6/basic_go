package basic_type

func Byte() {
	var a byte = 'a'
	println(a) // ascii code

	var str string = "this is string"
	var bs []byte = []byte(str)
	println(bs)
}
