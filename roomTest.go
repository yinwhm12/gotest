package main

import "fmt"

var rooms []*Room

type Room struct {
	Name string
	Pwd string
}

func init()  {
	rooms = make([]*Room,0)
}
func main()  {
	fmt.Println(len(rooms),cap(rooms))
	//rooms[1]
	rooms = append(rooms,&Room{"y","332"})
	rooms = append(rooms,&Room{"y2","3321"})
	rooms = append(rooms,&Room{"y3","3322"})
	fmt.Println(len(rooms),cap(rooms))
	//fmt.Println()
	fmt.Printf("%v", rooms)
}

