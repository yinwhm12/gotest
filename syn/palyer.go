package syn

import (
	"fmt"
	"time"
)

type Player struct {
	PlayerId int
	RoomId int

	RoomSignal chan int
}

var Players map[int]*Player

func init()  {
	Players = make(map[int]*Player)
}

func AddPlayers(p1,p2,p3,p4 *Player)  {

}

func (p *Player)StartPlayer()  {
	fmt.Println("iiiiii")
	room := Rooms[p.RoomId]
	select {
	case <- p.RoomSignal:
		var input int
		fmt.Println("input please")
		fmt.Scanln(&input)
		fmt.Println("input something --",input)
		Rooms[p.RoomId].PlayerSignal <- (Rooms[p.RoomId].RoomTurn + 1) % 4
	case <- time.After(time.Second *2):
		fmt.Println("有一个玩家超时没接收到room的信号")
		fmt.Println("room=%v",Rooms[p.RoomId])
		Rooms[p.RoomId].PlayerSignal <- (Rooms[p.RoomId].RoomTurn + 1) % 4

	}
	Players[room.RoomPlayerMap[(room.RoomTurn + 1) % 4]].RoomSignal <- 1
	//Players[room.RoomPlayerMap[(room.RoomTurn + 1)% 4 ]].StartPlayer()
	fmt.Println("palyer end")
}

