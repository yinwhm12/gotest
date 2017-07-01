package main

import (
	"gotest/algorithm"
	"fmt"
)

func main()  {
	//a := []int{101,101,101,101,102,102,104,104,104,105,105}
	//a := []int{101,101,102,102,104,104,104,105,105}
	//a := []int{101,101,102,102,104,104,105,105}
	//a := []int{101,101,102,102,103,103,105,105}
	//a := []int{101,101,102,102,103,103,106,106}
	//a := []int{101,101,102,102,102,103,103,106,106}
	//a := []int{101,101,102,102,103,103,104,104,106,106}
	//a := []int{101,101,102,102,103,103,104,104}
	a := []int{101,101,102,102,103,103,104,104,105,105}
	var m map[int]int
	m = make(map[int]int)
	algorithm.CountCards(a,m,401)
	fmt.Println("1-",m)
	fmt.Println(a)
	count := algorithm.AboveDoubleArray(a,m)
	fmt.Println("count=",count)
	fmt.Println("m=",m)

}
