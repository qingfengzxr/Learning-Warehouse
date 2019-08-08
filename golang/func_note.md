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









