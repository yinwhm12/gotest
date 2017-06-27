package main

import (
	"fmt"
	"gotest/algorithm"
)

func main()  {
	//a := []int{101,102,103,104,105,105,106,106,107,107,107,108,109}
	////flag := game.testHu(101,a,406)
	//flag := card.TestHu(101,a,406)
	//fmt.Println(flag)
	//game.
	//a := []int{105,105,105,106,107,108,201,202,203,401,401,401,403,403}
	//a := []int{105,105,105,106,107}
	//a := []int{106,107,108,108,108}
	//a := []int{106,107,108,108,108,201,202,203}
	//a := []int{106,106,106,107,107,107,108,108,108,109,109}
	//a := []int{105,105,105,106,107,108,401,401}
	//a := []int{101,101,101,102,103,103,103,104,104,105,105}
	//a := []int{101,102,103,104,105,106,107,108,109,109,109}
	//a := []int{101,101,101,101,102,103,104,104}//1111 2 3 44
	//a := []int{101,101,101,101,102,102,102,103}//1111 222 3
	//a := []int{101,101,101,101,102,102,103,103}//1111 22 33
	//a := []int{101,101,101,101,102,103,103,103} //1111 2 333
	//a := []int{101,101,101,101,102,102,103,103,104,105,105}//1111 22 33 4 55
	a := []int{101,102,103,204,205,206,303,304,305,306,306}//1111 22 33 4 55
	m := make(map[int]int)
	//m[2]=0
	//m[3]=0
	algorithm.CountSimilarCards(a,m)
	fmt.Printf("m=%v\n", m )
	count := algorithm.TestHu1(a,m)
	b := algorithm.EndSum(m)
	if b[1] != 0{
		fmt.Println("no hu")
	}else {
		left := count *3 + b[2] *2 + b[3] *3
		if len(a) - left == 0{
			fmt.Println("hu l")
		}else {
			fmt.Println("no hu")
		}
	}
	//fmt.Println("count=",count)
	////var two int
	//twoCount := 0
	//threeCount := 0
	//if v, ok := m[2]; ok&&v==1{
	//	twoCount = v
	//}
	//if v, ok := m[3]; ok&&v>0{
	//	threeCount = v
	//}
	//
	//fmt.Println("two=",twoCount,"three=",threeCount)
	//fmt.Printf("out_m=%v\n", m)
	//left := count * 3 + twoCount * 2 + threeCount * 3
	//if len(a) - left == 0{
	//	fmt.Println("hu l")
	//}else {
	//	fmt.Println("no hu ")
	//}
	//fmt.Println()

}
