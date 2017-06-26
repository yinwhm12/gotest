package main

import (
	"fmt"
	"sort"
)

type CardData struct {
	CardType int
	CardValue	int
}

type Cards []*CardData

func (c Cards)Len() int {
	return len(c)
}

func (c Cards)Less(i, j int) bool {
	if c[i].CardType == c[j].CardType{
		return c[i].CardValue < c[j].CardValue
	}else if c[i].CardType < c[j].CardType{
		return true
	}else {
		return false
	}

	//return
}

func (c Cards)Swap(i, j int)  {
	c[i], c[j] = c[j], c[i]
}

func main()  {
	cs := make(Cards,5)
	cs[0] = &CardData{3,5}
	cs[1] = &CardData{2,3}
	cs[2] = &CardData{1,2}
	cs[3] = &CardData{1,1}
	cs[4] = &CardData{8,5}
	fmt.Println(cs)
	for i, v := range cs{
		fmt.Println(i, *v)
	}
	sort.Sort(cs)
	for i, v := range cs{
		fmt.Println("----",i, *v)
	}
}


