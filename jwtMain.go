package main

import (
	"gotest/jwt"
	"time"
	"fmt"
)

func main()  {
	tokenString, _ := jwt.CreateToken()
	valid, rtime := jwt.RequireTokenAuthentication(tokenString)

	time.Sleep(time.Second * 2)
	valid, rtime = jwt.RequireTokenAuthentication(tokenString)
	fmt.Printf("2:%v:%v\n", valid, rtime)

	time.Sleep(time.Second * 3)
	valid, rtime = jwt.RequireTokenAuthentication(tokenString)
	fmt.Printf("3:%v:%v\n", valid, rtime)

	time.Sleep(time.Second * 2)
	valid, rtime = jwt.RequireTokenAuthentication(tokenString)
	fmt.Printf("4:%v:%v\n", valid, rtime)

	time.Sleep(time.Second * 2)
	valid, rtime = jwt.RequireTokenAuthentication(tokenString)
	fmt.Printf("5:%v:%v\n", valid, rtime)
	
}