package main

import f "fmt"

func main(){
	p := person{"wq",19}
	p.sayhi()
	s := student{person{"xiao",29},"001","110"}
	s.sayhi()
	s.setNo("007")
	s.sayhi()
	var i Int = 90
	i.parseInt("hh")
}

type person struct {
	name string
	age int
}

type student struct {
	person
	no string
	phone string
}

func(p person) sayhi() {
	f.Println(p.name,p.age)
}

func (s student) sayhi() {
	f.Println(s.no,s.name,s.age,s)
}

func (s *student) setNo(no string) {//如果不传指针, 则修改no时,失败
	s.no = no
}	

type Int int

func (this Int) parseInt(str string) {
	f.Println("Not exists!")
}

//func (this Int) parseInt() {  //not support !
//	println("no,no,no")
//}
