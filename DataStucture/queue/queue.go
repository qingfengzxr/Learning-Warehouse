package main

import (
	"fmt"
)

type Queue struct{
	data	[10]int
	head	int
	tail	int
}

//这里要注意值接收者和指针接收者的差别
func (q *Queue) Put(d int){
	q.data[q.tail] = d
	q.tail += 1
}

func (q *Queue) Get() int{
	var tmp int
	tmp = q.data[q.head]
	q.head++
	return tmp 
}

func (q *Queue) IsEmpty() bool{
	if q.head ==  q.tail {
		return true
	}else{
		return false
	}
}	

func main() {
	//tq := Queue{head:0,
	//			tail:0}

	//or
	var tq Queue

	tq.head = 0
	tq.tail = 0

	if tq.IsEmpty() == true {
		fmt.Println("The queue is empty")
	}

	fmt.Println(tq.data)
	
	for i := 0; i < 9 ;i++ {
		tq.Put(i)
	}

	fmt.Println(tq.data)
	fmt.Println(tq.Get())
	fmt.Println(tq.Get())
	fmt.Println(tq.head,tq.tail)

}
