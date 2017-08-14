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

	var s2 []int
	s2 = make([]int,len(s0)*2)
	//var s3 []int
	//var s4 []int

	s2 = s0
	fmt.Println("s2==",s2,"len",len(s2),cap(s2))
	s2 = s2[:len(s2)-1]
	fmt.Println("s2==",s2,"len",len(s2),cap(s2))
	s2 = append(s2,8)
	fmt.Println("s2==",s2,"len",len(s2),cap(s2))


}
