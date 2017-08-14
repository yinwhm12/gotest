package main

import (
	"time"
	"fmt"
)

func main()  {
	t := (int)(time.Now().Unix())
	fmt.Println(t)
}
