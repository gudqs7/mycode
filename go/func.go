package main

import "fmt"

var me map[string]string = map[string]string{"name":"wq","sex":"nan","habbits":"good"}

func main(){
	sum := jiechen(4);
	fmt.Println(jiechen)
	fmt.Println(sum)
	k,v := getkv("age")
	k1,v1 := getkv("sex")
	fmt.Println(k,v,k1,v1)
	printArgs(1,3,9)
	array := [3]int{10,20,30}
	fmt.Println("firstArray:",array)
	changeArray(array)
	fmt.Println("last:",array)
	slice := array[:]
	fmt.Println("firstSlice:",slice)
	changeSlice(slice)
	fmt.Println("last:",slice)
	tryDefer()
	fmt.Println(tryDefer)
}

func jiechen(n int) int {
	if n==1 {
		 return 1
	}
	return n*jiechen(n-1)
}

func getkv(key string) (string,string){
	val , ok := me[key]
	if ok {
		return key,val
	}else{
		return key,"No Exists"
	}
}

func printArgs(arg ...int){
	for i,val := range arg {
		fmt.Println(i,val)
	}
}

func changeArray(array [3]int){
	array[0]=0
	array[1]=0
	array[2]=0
}

func changeSlice(slice []int){
	slice[0]=0
	slice[1]=0
	slice[2]=0
}

func tryDefer(){
	for i:= 0; i<10; i++ {
		defer fmt.Println("Defer:",i)
	}
}
