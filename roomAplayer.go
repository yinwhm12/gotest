package main

import (
	"gotest/syn"
	"fmt"
)

func main()  {
	m := make(chan int,1)
	m1 := make(chan int,1)
	m2 := make(chan int,1)
	m3 := make(chan int,1)
	player1 := syn.Player{PlayerId:0,RoomId: 1,}
	player1.RoomSignal = m
	player2 := syn.Player{PlayerId:1,RoomId: 1,}
	player2.RoomSignal = m1
	player3 := syn.Player{PlayerId:2,RoomId: 1,}
	player3.RoomSignal = m2
	player4 := syn.Player{PlayerId:3,RoomId: 1,}
	player4.RoomSignal = m3


	//Players[player1.PlayerId] = player1
	syn.Players[player1.PlayerId] = &player1
	syn.Players[player2.PlayerId] = &player2
	syn.Players[player3.PlayerId] = &player3
	syn.Players[player4.PlayerId] = &player4

	c := make(chan int,1)
	c <- 1
	room1 := syn.Room{RoomId:1,PlayerSignal:c,}
	room1.RoomPlayerMap = make(map[int]int)
	room1.RoomPlayerMap[0] = player1.PlayerId
	room1.RoomPlayerMap[1] = player2.PlayerId
	room1.RoomPlayerMap[2] = player3.PlayerId
	room1.RoomPlayerMap[3] = player4.PlayerId
	room2 := syn.Room{RoomId:2,}

	syn.Rooms[room1.RoomId] = &room1
	syn.Rooms[room2.RoomId] = &room2
		//go player1.StartPlayer()
	 room1.StartRoom()
	//go player1.StartPlayer()
	//go player2.StartPlayer()
	//go player3.StartPlayer()
	//go player4.StartPlayer()
	//time.Sleep(time.Second * 10)
	fmt.Println("end----------")

}
