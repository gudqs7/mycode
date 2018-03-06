package main

import (
	"fmt"
	"runtime"
)

func say (s string) {
	for i :=0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("Hello")
	a := []int{1,2,3,4,5,6,7,8,9}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c 
	fmt.Println(x,y,x+y)
}

func sum (a []int, c chan int) {
	total := 0
	for _, v:= range a {
		total += v
	}
	c <- total
}

