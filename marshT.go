package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Register struct {
	ACTION int32
	SID    int32
}

func ExampleWrite() []byte {
	buf := new(bytes.Buffer)

	var info Register
	info.ACTION = 20004
	info.SID = 6

	err := binary.Write(buf, binary.LittleEndian, info)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x\n", buf.Bytes())
	return buf.Bytes()
}

func ExampleRead(b []byte) {
	var info Register
	buf := bytes.NewBuffer(b)

	err := binary.Read(buf, binary.LittleEndian, &info)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Print(info)
	// Output: 3.141592653589793
}

func main() {
	buf := ExampleWrite()
	ExampleRead(buf)
}