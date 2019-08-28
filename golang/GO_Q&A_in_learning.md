### GO语言疑问与解答

1. 为什么cap(b) 等于 8 ？

```go
a := [...]int{0,1,2,3,4,5,6}
b := make([]int,2,4)
c := a[0:3]

fmt.Println(len(b))  //2
fmt.Println(cap(b))  //4
b = append(b,1)
fmt.Println(b)  //[0,0,1]
fmt.Println(len(b))  //3
fmt.Println(cap(b))  //4

b = append(b,c...)
fmt.Println(b)  //[0 0 1 0 1 2]
fmt.Println(len(b))  //6
fmt.Println(cap(b))  //8

b = append(b,c...)
fmt.Println(b)  //[0 0 1 0 1 2 0 1 2]
fmt.Println(cap(b))  //16   
```

原因：切片在追加元素时如果容量cap不足，将按len的2倍扩容。

