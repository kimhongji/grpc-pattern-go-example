package main

import (
	"fmt"
	"time"
)

func channelTest1() {
	ch := make(chan int)
	go func() {
		fmt.Println("before ..")
		var i int
		for {
			i++
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		fmt.Println("i'm main...")
		fmt.Println(<- ch)
		time.Sleep(5 * time.Second)
	}
}

func channelTest2(name string, ch chan int){
	for {
		fmt.Println(name, " turn")
		i := <- ch
		fmt.Println(name, ": ", i)
		i++
		ch <- i

		time.Sleep(1 * time.Second)
	}
}

func main() {
	// case 1
	//channelTest1()

	// case 2
	//ch := make(chan int)
	//go channelTest2("player1", ch)
	//go channelTest2("player2", ch)
	//
	//ch <- 1
	//
	//for {
	//	time.Sleep(1 * time.Second)
	//}


}