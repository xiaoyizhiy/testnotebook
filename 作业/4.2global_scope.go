package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	print(a)
}

func m() {
	a = "O" //因为该局部作用域中没有a，所以将O赋值给了全局变量a
	print(a)
}

//输出结果为 GOO
