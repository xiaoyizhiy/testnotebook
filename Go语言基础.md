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

## 变量

```go
package main
import(
  "fmt"
  "math"
)
func main(){
var a="initial"
var b，c int=1，2
var d= true
var e float64
f := float32(e)
g := a + "foo"
fmt.Println(a,b,c,d,e,f)//initial1 2 true 0 0
fmt.Println(g)          //initialapple
const s string ="constant"
const h= 500000000
const i=3e20 /h
 fmt.Println(s,h,i,math.Sin(h),math.Sin(i))
```

