package main

import (
	"time"
	"fmt"
)

//获取昨天的时间 0点开始
func main()  {
	t := int(time.Now().Unix())
	tt := 1499739252
	fmt.Println(time.Unix(int64(t),0))
	fmt.Println(time.Unix(int64(t),0))
	fmt.Println(time.Unix(int64(t-tt),0))
	fmt.Println(time.Unix(int64(tt-86400),0))
	fmt.Println(time.Now().Date())
	fmt.Println("t-tt",t-tt)

	nTime := time.Now()
	day := nTime.Day()
	fmt.Println("day=",day)

	yesTime := nTime.AddDate(0,0,-1)
	hour := yesTime.Hour() * 60*60 + yesTime.Minute() *60 + yesTime.Second()
	l := int(yesTime.Unix())-hour
	fmt.Println(l)
	fmt.Println(time.Unix(int64(l),0))

	timeStr := time.Now().Format("2006-01-02")
	fmt.Println("timeStr",timeStr)
	ta, _ := time.Parse("2006-01-02",timeStr)
	timeNumber := ta.Unix()
	fmt.Println("timeNumber:", timeNumber)
	fmt.Println(time.Unix(timeNumber-8*60*60,0))
	//fmt.Println("ttt",timeNumber)
}
