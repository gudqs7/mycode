package main

import "fmt"

func main(){
	num, num2 := 1, 1
	for num<=9 {
		num2 = 1
		for num2<=num {
			fmt.Printf("%d * %d = %d\t",num2,num,num*num2)
			num2++
		}
		fmt.Println()
		num++
	}
}
