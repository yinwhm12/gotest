package main

import (
	"fmt"
	"sort"
)

func main()  {
	a := []int{1,2,3,4,4,5,6,7}
	//b := []int{}
	 b :=make([]int,len(a))
	//b := []int{}
	copy(b, a)
	for i,v := range a{
		if v == 4{
			a[i] = 8
		}
	}
	fmt.Println("a1 ==",a)
	test11(a)
	fmt.Println("a2 ==",a)
	fmt.Println("b3 ==",b)
	sort.Ints(a)
	fmt.Println("a3 ==",a)
}
func test11(a []int)  {
	a[1] = 4
}
