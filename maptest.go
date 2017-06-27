package main

import "fmt"

// 修改后 值是否立刻生效
func main()  {
	var m map[int]int
	m = make(map[int]int)
	m[101]=2
	m[103]=1
	m[104]=3
	m[105]=4
	copyM := new(map[int]int)
	copyM = &m
	//copyM[106] = 2
	fmt.Printf("m=%v\n", m)
	fmt.Printf("copyM=%v\n", copyM)
	//for i,v := range m{
	//	if v == 3{
	//		fmt.Println("i=",i)
	//		m[i]--
	//		m[i+1]--
	//		//fmt.Println("")
	//	}
	//	fmt.Printf("i==%d,v=%v\n", i,m[i])
	//}
	//for i := 101; i< len(m); i++{
	//	fmt.Println(m[i])
	//}
	//fmt.Println(m[1])

	//a := []int{1,2,4,5,6,7,8,9}
	//for _,v := range a{
	//	fmt.Println(v)
	//}
}
