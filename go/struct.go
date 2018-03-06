package main

import "fmt"

func main(){
	var h humen 
	h.name="hwq"
	h.age=20
	fmt.Println("humen:",h)
	h2 := humen{"h2wq",22,179}
	fmt.Println("h2:",h2)
	h3 := humen{name:"h3wq",age:99}
	fmt.Println("h3:",h3)
	p := person{humen{"wq",20,180},99,[]string{"basketball","java","go"},"personwq"}
	fmt.Println("first:",p)
	changeStruct(&p)
	fmt.Println(p,"\n",p.name,p.age,p.height)
	p.Skill=append(p.Skill,"c","python")
	fmt.Println(p.Skill,p.int,p.Skill[2])
}

type humen struct {
	name string
	age int
	height int
}

type Skill []string

type person struct {
	humen
	int
	Skill
	name string
}

func changeStruct(p *person){
	p.name="jiade"
	fmt.Println((*p).name)
	(*p).name = "qq"
}
