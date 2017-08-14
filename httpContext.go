package main

import (
	"net/http"
	"context"
	"time"
	"fmt"
)

func main()  {
	go http.ListenAndServe(":8080",nil)
	ctx, _ := context.WithTimeout(context.Background(),(25 * time.Second))
	go testA(ctx)
	select {
	case <- ctx.Done():
		fmt.Println("main")

	}
	
}

func testA(ctx context.Context)  {
	ctxA, _:= context.WithTimeout(ctx,(5 * time.Second) )
	ch := make(chan int)
	go testB(ctxA,ch)
	select {
	case <- ctx.Done():
		fmt.Println("testA Done")
		return
	case i := <- ch:
		fmt.Println(i)
	}
}

func testB(ctx context.Context, ch chan int)  {
	sumch := make(chan int)
	go func(sumch chan int) {
		sum := 10
		time.Sleep(10 * time.Second)
		sumch <- sum
	}(sumch)
	select {
	case <- ctx.Done():
		fmt.Println("testB Done")
		return
	case i := <- sumch:
		fmt.Println("send ", i)
		ch <- i
	}

}