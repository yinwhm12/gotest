package main

import (
	"net"
	"encoding/binary"
)

//type Hello struct {
//	Name string
//}

func main()  {
	conn, err := net.Dial("tcp","127.0.0.1:3563")
	if err != nil{
		panic(err)
	}
	//hello :=` {"Hello":{"Name":"leaf"}}`
	data := []byte(`{
				"Hello" : {
					"Name" : "leaf"
					}}`)

	//fmt.Println("---hello---",data)
	//var data []byte
	//err = json.Unmarshal([]byte(hello),&data)
	//if err != nil{
	//	fmt.Println("Unmarshal JSON")
	//	return
	//}
	m := make([]byte, 2 + len(data))

	binary.BigEndian.PutUint16(m, uint16(len(data)))
	//fmt.Println("------mm--",m)

	copy(m[2:],data)
	//fmt.Println("------mmdddd--",m)

	conn.Write(m)
	
}
