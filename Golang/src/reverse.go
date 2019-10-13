package main

import (
	"fmt"
)

/*@func: 反转slice
**@params: s []int：需要反转的slice
**@return: 无
*/
func reverse(s []int){
	for i ,j := 0,len(s)-1;i<j;i,j = i+1,j-1{
		s[i], s[j] = s[j], s[i]
	}
}

//将一个slice左移n个元素的简单办法是连续调用reverse()三次
//1.反转前n个元素
//2.反转剩下的元素
//3.整体反转一次
func main(){
	s := []int{0,1,2,3,4,5}
	reverse(s[:2])
	fmt.Println(s)
	reverse(s[2:])
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)
}
