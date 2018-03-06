package main

import "errors"

import "fmt"

func main(){
	age,err := getwrong(-1)
	if err !=nil {
		fmt.Println(err)
	}
	println(age)
}

func getwrong(age int)(int,error) {
	if age <0 {
		return 0, errors.New("what? fuck u !")
	}
	return age,nil
}
