package main

import "fmt"

func main()  {
	s0 := []int{1,2,3,4,5,6,78,9}
	fmt.Println("sssss",s0)
	fmt.Println("sssssff",len(s0[:5]))
	fmt.Println("sssssfffffss",s0[:5])
	fmt.Println("sssssff",s0[5])
	//s1 := make([]int)
	s1 := make([]int,5)

	copy(s1,s0[:5])
	fmt.Println("so",s0)
	fmt.Println("s1",s1)
	s1 = append(s1,s0[5])
	fmt.Println("add",s1)

}
