package main

import "fmt"

func main(){
	p := NewPerson(9,"gudqs")
	fmt.Println(p,p.name,p.age,*p,&p,*(&p),(*p).name)
	(*p).name = "pwq"
	p.name = "ppwq"
	fmt.Println(p.name,(*p).name)
}

type person struct {
	name string
	age int
}

func NewPerson(age int, name string) *person {
	if age <0 {
		return nil
	}
	return &person{name,age}
}
