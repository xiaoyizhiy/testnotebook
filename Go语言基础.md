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
fmt .Println("You head further up the mountain.")}else if command =="go inside"{
在第一次检查为假之后，检查命令是否为“go inside”
fmt.Println("You enter the cave where you live out the rest of your life.")}else {
如果前两次检查都为假，
fmt.Println("Didn't quite get that.")
那么执行第三个分支
J
```



## 循环



## switch



## 数组



## 切片

