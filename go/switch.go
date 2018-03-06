package main

import "fmt"

func main(){
	var i int
	fmt.Scanln(&i,&i)
	switch i {
		case 1:
			fmt.Println("i is equal to 1")
		case 2,3,4,5,6:
			fmt.Println("i is 2,3,4,5,6 ")
		default:
			fmt.Println("i is not 1-6")
	}
}
