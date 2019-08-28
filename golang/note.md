#### 函数

* 函数定义

```go
func funcName(param-list)(result-list){
	function-body
}
```

* 函数签名

函数类型又叫函数签名，一个函数的类型就是函数定义首行去掉函数名、参数名和{，可以使用fmt.Printf的“%T”格式化参数打印函数类型。

两个函数类型相同的条件是：拥有相同形参列表和返回值列表（列表元素的次序、个数和类型都相同），形参名可以不同。

可以用type定义函数类型，函数类型变量可以作为函数的参数或者返回值。

```go
func add(a, b int) int {
	return a + b
}
fmt.Printf("%T\n", add) // func(int,int) int
```

```go
package main

import "fmt"

func add(a,b int) int {
    return a+b
}

func sub(a,b int) int {
    return a-b
}

type Op func(int, int) int //定义一个函数类型

func do(f Op, a, b int) int { //定义一个函数，第一个参数是函数类型Op
    return f(a, b) //函数类型变量可以直接用来进行函数调用
}

func main(){
    a := do(add,1,2) //函数名add可以当作相同函数类型形参，不需要强制类型转换
    fmt.Println(a)
    s := do(sub,1,2)
    fmt.Println(s)
}

```

：不知道为什么，很有一种Python中装饰器的味道，暂时来说不知道在go语言的开发中会有怎样的用途，适用于怎样的场景，但感觉用途挺有意思的。

**注意：**
1. GO函数使用 caller-save 模式，即由调用者负责保存寄存器，由主调函数保存和恢复现场。
2. GO内嵌汇编和反汇编产生的代码并不是一一对应的，汇编编译器对内嵌汇编程序自动做了调整，主要差别就是增加了保护现场，以及函数调用前的保持PC、SP偏移地址重定位等逻辑。
3. GO函数调用前已经为返回值和参数分配了栈空间，分配顺序是从右向左的，先是返回值，然后是参数。
4. 函数的多值返回是主调函数预先分配好存放空间，然后被调函数执行时将返回值复制到该返回位置实现。



#### 结构体

* 初始化

```go
type Person struct {
	name string
	age int
}

//比较推荐指定字段名初始化
a := Person{name:"andes",age:"18"}
a := Person{
    name:"andes",
    age:18,
}
a := Person{
    name:"andes",
    age:18}
//初始化语句的末尾的‘}’独占一行时，最后一个字段的后面一定要带上逗号
```

* 匿名字段

在定义struct的过程中，如果字段只给出字段类型，没有给出字段名，则称这样的字段为匿名字段。

```go
type File struct{
	*file
}
```

1. 被匿名嵌入的字段必须是命名类型或命名类型的指针，类型字面量不能作为匿名字段使用。
2. 匿名字段的字段名默认就是类型名。
3. 如果匿名字段是指针类型，则默认的字段名就是指针指向的类型名。
4. 一个结构体里面不能同时存在某一类型及其指针类型的匿名字段，因为二者的字段名相等。
5. 如果嵌入的字段来自其他包，则需要加上包名，并且必须是其他包可以导出的类型。





