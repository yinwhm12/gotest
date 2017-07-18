package syn

import (
	"fmt"
	"time"
)

type Player struct {
	PlayerId int
	RoomId int

	RoomSignal chan int
	OutCardSignal chan int
}

var Players map[int]*Player

func init()  {
	Players = make(map[int]*Player)
}

func AddPlayers(p1,p2,p3,p4 *Player)  {

}

func (p *Player)StartPlayer()  {
	fmt.Println("iiiiii")
	//room := Rooms[p.RoomId]
	id := p.PlayerId
	var input string
	select {
	case <- p.RoomSignal:
		fmt.Println("input please")
		//t := time.NewTimer(3 * time.Second)

		//go p.Input()
		//go p.TimeOver()
		//go func(id int) {
		//	fmt.Scanln(&input)
		//	fmt.Println("input something --",input)
		//}(id)
		fmt.Scanln(&input)
		fmt.Println("input something --",input)
		//p.OutCardSignal <- 1//没有超时 已打出
		Rooms[p.RoomId].PlayerSignal <- (Rooms[p.RoomId].RoomTurn + 1) % 4
	case <- time.After(time.Second *2):
		fmt.Println("有一个玩家超时没接收到room的信号")
		fmt.Println("room=%v",Rooms[p.RoomId])
		Rooms[p.RoomId].PlayerSignal <- (Rooms[p.RoomId].RoomTurn + 1) % 4

	}
	//id := room.RoomPlayerMap[(room.RoomTurn + 1) % 4]
	fmt.Println("playerid :-=",id)
	fmt.Println("player[id]=%v",Players[id])
	Players[id].RoomSignal <- 1
	//Players[room.RoomPlayerMap[(room.RoomTurn + 1)% 4 ]].StartPlayer()
	fmt.Println("palyer end")
}

func (p *Player)TimeOver()  {
	select {
	case <- time.After(time.Second * 3):
		fmt.Println("time over")

		//exec.CommandContext()
		Rooms[p.RoomId].PlayerSignal <- (Rooms[p.RoomId].RoomTurn + 1) % 4
	case <- p.OutCardSignal:
		fmt.Println("没有超时")
	}
}

func (p *Player)Input()  {
	fmt.Println("input please")
	var input string
	//select {
	//case <- fmt.Scanln(&input):
	//}
	fmt.Scanln(&input)
	fmt.Println("input something --",input)
	p.OutCardSignal <- 1
}
