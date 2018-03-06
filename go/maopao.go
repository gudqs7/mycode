package main

import "fmt"

func main(){
	slice := []int{3,4,2,9,1,0,88,44,20}
	maopao(slice)
	fmt.Println(slice)
}

func swap(slice []int, index1 int,index2 int) {
	temp := slice[index1]
	slice[index1]=slice[index2]
	slice[index2]=temp
}

func maopao(slice []int) {
	for i := 0; i< len(slice)-1;i++ {
		for j:=0; j< len(slice)-1-i; j++ {
			if slice[j]>slice[j+1] {
				swap(slice,j,j+1)
			}
		}
	}
}
