package main

var a = "G" //初始化全局变量a

func main() {
	n()
	m()
	n() //输出时在一行输出完
}

func n() {
	print(a)
} //a为全局变量域中的全局变量a

func m() {
	a := "O"
	print(a)
} //初始化局部变量a，只作用于m函数的局部变量域中，作用完就消失了
//输出结果为GOG
