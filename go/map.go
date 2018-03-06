package main

import "fmt"

func main(){
	var numbers map[string]int =map[string]int{}
	others := make(map[string]int)
	others["name"]=007
	numbers["one"] = 1
	numbers["ten"] = 10
	numbers["three"] = 3
	delete(numbers,"one")
	val,ok := numbers["one"]
	fmt.Println(numbers,numbers["one"],others,val,ok,len(numbers))
}
