package main

import (
	"fmt"
	"gotest/card"
)

func main()  {
	a := []int{101,102,103,104,105,105,106,106,107,107,107,108,109}
	//flag := game.testHu(101,a,406)
	flag := card.TestHu(101,a,406)
	fmt.Println(flag)
	//game.

	//fmt.Println()
}
