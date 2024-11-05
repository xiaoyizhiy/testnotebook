Go语言基础

1. 高性能、高并发
2. 语法简单、学习曲线平缓
3. 丰富的标准库完善的工具链
4. 静态链接
5. 快速编译
6. 跨平台
7. 垃圾回收

字节跳动为什么全面拥抱 Go 语言

1. 最初使用的 Python，由于性能问题换成了 Go
2. C++ 不太适合在线 Web 业务
3. 早期团队非 Java 背景
4. 性能比较好
5. 部署简单、学习成本低
6. 内部 RPC 和 HTTP 框架的推广

## 包和函数

```go
package main  #声明本代码所属的包
import(
"fmt"#导入fmt(是format的缩写)包，使其可用
)
func main(){ # func 关键字用于声明函数声明一个名为main 的函数，fmt 包提供了用于格式化输入和输出的函数。
    fmt.Printin("hello world” )在屏幕上打印出“Hello,playground”
}   #这是 Go 语言唯一允许的大括号放置风格    

```

package 关键字声明了代码所属的包，在本例中这个包的名字就是 main。所有用 Go编写的代码都会被组织成各式各样的包，并且每个包都对应一个单独的构想。

main 这一标识符(identifer)具有特殊意义。当我们运行一个 Go程序的时候，它总是从 main 包的 main 函数开始运行。如果 main 不存在，那么 Go编译器将报告一个错误,因为它无法得知程序应该从何处开始执行。

每次用到被导入包中的某个函数时，我们都需要在函数的名字前面加上包的名字以及一个点号作为前缀。Go 的这一特性可以让用户在阅读代码的时候立即弄清楚各个函数分别来自哪个包。

## 报错

undefined: fmt.printlncompilerUndeclaredImportedName

确保你没有错误拼写 `fmt.Println`。函数名应为 `Println`，而不是 `println` 或其他形式。

expected '(', found '{'syntax

具体来说，编译器或解释器期望在某个地方看到一个左括号 `(`，但实际看到的是一个左大括号 `{`。func main(){  将main后的（）忘记写了

## 计算

Go 跟其他编程语言一样，提供了+、-、*、/、%（取模：取余数）。

```go
//我的减重程序
package main
import "fmt"
//main 是所有程序的起始函数
func main(){
fmt.Print("My weight on the surface of Mars is ")
fmt.Print(149.0*0.3783)
fmt.Print(" lbs,and I would be ")
fmt.Print(41*365/687) 
fmt.Print("years old.")
//或者为：fmt,Println("My weight on the surface of Mars is", 149.0*0.3783, "lbs, and I would be",41*365.2425/687，"years old."),
```

打印出“My weight on the surface of Mars is 56.3667 lbs, and I wouldbe21.79758733624454 years old.”

## 变量

程序员有时候会把这种没有说明具体含义的字面数字称为魔数

var 变量名类型=表达式

var name string="zhangsan'

在函数内部，可以使用更简略的:方式声明并初始化变量。(短变量)只能用于声明局部变量，不能用于全局变量的声明。



```go
// 到达火星需要多长时间?
package main
import "fmt"
func main(){
const lightSpeed=299792//km/s  //两个新的关键字const和var，它们分别用于声明常量和变量。
var distance=56000000 // km    //变量必须先声明再赋值 distance=56000000会报错
fmt.Println(distance/lightSpeed，"seconds")  //打印出186 secends
distance =401000000
fmt.Println(distance/lightSpeed，"seconds")  //打印出1337 secends
}
```

注意 lightSpeed 常量是不能被修改的，尝试为其赋予新值将导致 Go 编译器报告错误:“无法对 lightSpeed 进行赋值”

注意 变量必须先声明后使用。如果尚未使用 var 关键字对变量进行声明，那么尝试向它赋值将导致 Go 报告错误，例如在前面的代码中执行 speed = 16 就会这样。这一限制有助于发现类似于“想要向 distance 赋值却键入了 distence”这样的问题。



#### 一次声明多个变量

```go
var distance= 56000000
var speed=100800
```

```go
var (
   distance=56000000
   speed =100800
)

```

```
var distance,speed=56000000，100800
```

#### 数字游戏

```go
//随机数字
package main
import (
"fmt "
"math/rand"  //rand 包的导入路径为 math/rand
)
func main(){
var num=rand.Intn(10)+1   //Intn生成0~9的伪随机数（它们并非真正随机，只是看上去或多或少像是随机的而已。）   注意 如果我们在写代码的时候忘记对伪随机数执行加一操作，那么程序将返回一个0~9的数字而不是我们想要的1~10 的数字。这是典型的“差一错误”(off-by-one error)的例子，这种错误是典型的计算机编程错误之一。
fmt.Println(num)
num =rand.Intn(10)+1
fmt.Println(num)
```

通过 Printf函数和格式化变量v,用户可以将值放置到被显示文本的任意位置上。

```go
fmt.Printf("%-15v $%4v\n"，"SpaceX"，94)
fmt.Printf("%-15v $%4v\n","Virgin Galactic"，100)
//SpaceX          $ 94
//Virgin Galactic $ 100
```

对 Go 来说，true 是唯一的真值,而false 则是唯一的假值。

## 比较

比较运算符：既可以比较文本，又可以比较数值。

| 符号 | 含义   | 符号 | 含义     |
| ---- | ------ | ---- | -------- |
| ==   | 相等   | <=   | 小于等于 |
| ！=  | 不相等 | >    | 大于     |
| <    | 小于   | >=   | 大于等于 |

```go
//比较数值
fmt,Println("There is a sign near the entrance that reads 'No Minors'.")
var age = 4l
var minor=age<18  //从右向左看，age<18 为false
fmt.Printf("At age %v,am I a minor? %v\n",age, minor)
//输出There is a sign near the entrance that reads 'No Minors'At age 4l,am I a minor? false
```

Go 只提供了一个相等运算符，并且它不允许直接将文本和数值进行比较。

## if else

```go
package main
import "fmt"
func main(){
var command="go east"
if command =="go east"{   //检查命令是否为“go east”
  fmt .Println("You head further up the mountain.")
}else if command =="go inside"{
//在第一次检查为假之后，检查命令是否为“go inside”
  fmt.Println("You enter the cave where you live out the rest of your life.")
}else {
//如果前两次检查都为假，
  fmt.Println("Didn't quite get that.")
}
//那么执行第三个分支
}
```

## 逻辑运算符

在 Go中，逻辑运算符||代表“逻辑或”，而逻辑运算符&&则代表“逻辑与”。这些逻辑运算符可以一次检查多个条件。

逻辑或:当a、b两个值中至少有一个为true 时，allb为true

逻辑与:当且仅当a、b两个值都为 true 时，a && b为 true

```go
//能够被4整除但是不能被100整除的年份是闰年或者可以被 400整除的年份是闰年。
fmt.Println("The year is 2100,should you leap?")
var year =2100
var leap=year%400==0||(year%4==0 && year%100 != 0)   //%取余
if leap {
  fmt.Println("Look before you leap!")}else {
  fmt.Println("Keep your feet on the ground." )
  //输出The year is 2100,should you leap?
  //输出Keep your feet on the ground.
```

跟大多数编程语言一样，Go也采用了短路逻辑:如果位于|1运算符之前的第一个条件为真，那么位于!运算符之后的条件就可以被忽略，没有必要再对其进行求值。具体到代码清单 3-4中的例子，当给定年份可以被400整除时，程序就不必再进行后续的判断了。&&运算符的行为与1运算符正好相反:只有在两个条件都为真的情况下，运算结果才为真。对于代码清单3-4中的例子，如果给定年份无法被4整除，那么程序就不会对后续条件进行求值。

逻辑非运算符!可以将一个布尔值从false 变为true，或者将true 变为 false。

```go
var haveTorch =true //有火把
var litTorch =false //火把没有点亮
if !haveTorch ||!litTorch { //如果 false ||true (如果有火把或者没点亮)
  fmt.Println("Nothing to see here.")
}
```

## 循环



## switch

switch语句，它可以将单个值和多个值进行比较。除文本以外，switch语句还可以接受数值作为条件。

```go
fmt.Println("There is a cavern entrance here and a path to the east.")
var command ="go inside"
switch command { //将命令和给定的多个分支进行比较
case "go east": 
fmt.Println("You head further up the mountain.")
case "enter cave","go inside": //使用逗号分隔多个可选值
fmt.Println("You find yourself in a dimly lit cavern.")
case "read sian":
fmt.Println("The sign reads No Minors'.")
default:
fmt.Println("Didn't quite get that.")
//输出There is a cavern entrance here and a path to the east.
//You find yourself in a dimly lit cavern.
```

switch的另一种用法

在 Go 语言中，fallthrough是 switch语句的一部分，用来控制程序从当前 case 继续执行到下一个 case。默认情况下，Go 中的 switch 语句会在匹配到一个 case 后立即终止，而不会自动跳转到下一个 case。但是，当你使用 fallthrough 时，它会强制程序继续执行下一个 case 语句块的代码，而不管下一个 case 的条件是否满足。

```go
var room ="lake"
switch{ //比较表达式将被放置到单独的分支里面
case room==cave":
fmt,Println("You find yourself in a dimly lit cavern.")
case room =="lake":
  fmt.Println("The ice seems solid enough.")
  fallthrough //下降至下一分支
case room =="underwater": 
  fmt.Println("The water is freezing cold.")

//The ice seems solid enough.
//The water is freezing cold.
```

## 数组



## 切片



## map



## range



## 函数

rand函数的声明：

rand 包中的 Intn 函数的声明如下:
func Intn (n int) int
下面是一个使用 Intn 函数的例子:
num:= rand.Intn(10)

![image-20241105131255164](https://github.com/xiaoyizhiy/testnotebook/blob/main/%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20241105131657.png?raw=true)



## 指针



## 结构体



## 结构体方法



## 错误处理



## 字符串操作



## 字符串格式化



## JSON处理



## 时间处理



## 数字解析



## 进程信息



## Print、Printin、Printf

| println  | printf(格式化输出)   | prin       |
| -------- | -------------------- | ---------- |
| 有空格   | 利用占位符（%d、%f） | 无空格     |
| 自动换行 |                      | 不自动换行 |
|          |                      |            |

## 注释

ctrl+/ 可以快速的注释

```go
//一次输入多个值的时候 Println 中间有空格 Print 没有
fmt.Println("go","python","php","javascript") // go python php javascript
fmt.Print("go","python","php","javascript")//gopythonphpjavascript
//Println 会自动换行，Print 不会
package main
import "fmt'
func main() {
fmt.Println("hello")
fmt.Println("world")
// hello
// world
fmt.Print("hello")
fmt.Print("world")
// helloworld
}
```

