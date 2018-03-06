package main

import "fmt"

func main(){
	var slice []int = make([]int,2)
	fmt.Println(slice,len(slice),cap(slice))
	arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	aslice := arr[2:]
	bslice := aslice[:8]
	bslice[7] = 99
	arr[0] = -99
	fmt.Println(slice,aslice,bslice)
}
