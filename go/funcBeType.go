package main

import "fmt"

func main(){
	slice := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println("slice:",slice)
	odd := filter(slice,isOdd)	
	even:= filter(slice,isEven)
	fmt.Println(odd,even)
}

type funcType func(int) bool

func isOdd(i int) bool {
	if i%2 ==0 {
		return false
	}
	return true;
}

func isEven(i int) bool {
	if i%2 ==0 {
		return true
	}
	return false
}

func filter(slice []int,f funcType) []int {
	var result []int
	for _, value := range slice {
		if f(value){
			result = append(result, value)
		}
	}
	return result
}
