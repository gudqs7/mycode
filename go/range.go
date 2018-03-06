package main

import "fmt"

func main(){
	var array = [10]int{1,2,3,4,5,6,7,8,9}
	slice := array[:]
	for i , val := range slice {
		fmt.Println(i,val)
	}
	var mapa map[string]string = map[string]string{"name":"gudqs7","age":"haha"}
	for k,v := range mapa {
		fmt.Println(k,v)
	}
}
