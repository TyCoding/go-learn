package c02

type innerParam struct {
	A int
	B string
}

type OuterParam struct {
	A int
	B string
}

var (
	a int
	B string
)

const c int = 1

func innerM1() {
	println("innerFile.go 内部函数")
	outerM1()
	OuterM2()
}

func InnerM2() {
	println("innerFile.go 外部函数")
	outerM1()
	OuterM2()
}
