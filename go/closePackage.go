package main

import f "fmt"

func main(){
	i:=2
	if 2>1 {
		i:= 0
		f.Println(i)
	}
	f.Println(i)
	n := Number()
	f.Println(n())
	f.Println(n())
	n2 := Number()
	f.Println(n2())
	f.Println(n2())
}

func Number() func() int {
	i:= 0
	return func() int {
		i++
		return i;
	}
}
