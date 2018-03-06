package main

import "fmt"

func main(){
    isOk:=false
    isOk=true
	fmt.Println("isOk:",isOk)
	var array [10]int
	slice :=array[2:4]
	fmt.Println(slice,len(slice),cap(slice))
	slice =array[2:4:7] 
	fmt.Println(slice,len(slice),cap(slice))
}
