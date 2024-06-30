package constx

func Const() {
	const a = "hello"
	println(a)
}

const (
	Status0 = iota
	Status1
	Status2
	Status3
)

const (
	One = iota << 1
	Two
	Three
)
