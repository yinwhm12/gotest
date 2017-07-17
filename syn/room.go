package syn

import (
	"time"
	"fmt"
)

type Room struct {
	RoomId int
	RoomTurn int
	RoomPlayerMap map[int]int
	PlayerSignal chan  int

}

var Rooms map[int]*Room

func init()  {
	Rooms = make(map[int]*Room)
}

func (r *Room)StartRoom()  {
	for i:=0; i< 8 ; i++{
		for j := 0; j< 4; j++{
				//Players[r.RoomPlayerMap[r.RoomTurn]].StartPlayer()
			fmt.Println("444444--j=",j)
			select {
			case r.RoomTurn = <- r.PlayerSignal:
				fmt.Println("接收到一个玩家的信息 room")
				fmt.Println("%v",Players[r.RoomPlayerMap[r.RoomTurn]])
				Players[r.RoomPlayerMap[r.RoomTurn]].StartPlayer()
				//Players[r.RoomPlayerMap[r.RoomTurn]].RoomSignal <- 1
				fmt.Println("发送了一个信号量")
				//Players[r.RoomPlayerMap[r.RoomTurn]].StartPlayer()
				fmt.Println("%d=",Players[r.RoomPlayerMap[r.RoomTurn]].PlayerId)
				fmt.Println("give a info")

			case <- time.After(time.Second * 2):
				Players[r.RoomPlayerMap[r.RoomTurn]].RoomSignal <- 1

			}
		}
	}
}