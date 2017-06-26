package main

import "fmt"

type Test struct {
	Name string
	ID	int
}

func main()  {
	t := Test{"a",11}
	m := map[int]*Test{}
	m[1] = &t
	m[4] = &Test{"b",123}
	m[6] = &Test{"g",555}


	for i, v := range m {
		//fmt.Println("")
		if v == &t{
			m[i] = nil
			break
		}
		//fmt.Printf("i=%d,v=%v\n",i,v )
	}
	fmt.Println(m)
	if v, ok := m[7]; ok {
		fmt.Println("存在")
	}else {
		fmt.Println("不存在",v)
		m[7] = &Test{"v",44}
	}
	fmt.Println(m)


	
}
