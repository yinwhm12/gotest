package main

import (
	"os"
	"fmt"
	"log"
	"net"
	"time"
)

func main()  {
	if len(os.Args) <= 1{
		fmt.Println("input addr and message ")
		return
	}
	log.Println("begin dial...")
	conn, err := net.Dial("tcp",":8888")
	if err != nil{
		log.Println("error , dial")
		return
	}
	defer conn.Close()
	log.Println("Ok,dialed")
	time.Sleep(time.Second * 2)
	data := os.Args[1]
	conn.Write([]byte(data))

	time.Sleep(time.Second * 1000)
	
}
