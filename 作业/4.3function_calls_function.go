package main

var a string //初始化全局变量a

func main() {
   a = "G" //将G赋值给了全局变量a
   print(a) //打印G
   f1() //打印OG
}
//初始化局部变量a并打印，执行函数f2
func f1() {
   a := "O"
   print(a)
   f2()
}
//打印全局变量a
func f2() {
   print(a)
}