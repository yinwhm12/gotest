package main

import (
	"math/rand"
	"time"
	"fmt"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}

type a []int

func main()  {
	//for i:=0; i< 10; i++{
	//	fmt.Println("---", rand.Intn(20))
	//}
	m := make(map[int]int)//利用map key唯一性 产生不重复的随机数
	//m[1]= 1
	//m[2]= 2
	//m[3]= 3
	//m[4]= 4
	//m[5]= 5
	//a := make([]int,13)
	b := make([]a,4)
	fmt.Println("bb--",b)
	fmt.Println("b-len--",len(b))

	for i := 0; i<len(b); i++{
		for len(m) < 13 * ( i + 1 ) {
			n := rand.Intn(136)
			//m[n] = n
			if _,ok := m[n]; !ok {
				m[n] = n
				b[i] = append(b[i],n)
			}
			//b[i] = append(b[i],m[n])
		}
		//b[i] = append(b[i],m)
		//b = append(b,m)
		fmt.Println("hhh--",b[i])

	}
	////fmt.Println(m)
	//for _,v := range m {
	//	fmt.Println("ll---",v)
	//}
}
