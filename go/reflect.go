package main

import "reflect"
import "fmt"

func main(){
	a := 3.3
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	fmt.Println(a,t,v)
	fmt.Println(t.Kind(),v.Kind())
	fmt.Println(v.Type())
	fmt.Println(v.Interface())
	fmt.Println()
	p := Person{3,"wq"}
	fmt.Println(p)
	pt := reflect.TypeOf(p)
	pv := reflect.ValueOf(&p).Elem()
	fmt.Println("person:",pt,pv)
	for i := 0; i < pv.NumField(); i++ {
		fmt.Println("field:",pv.Field(i))
	}
	pv.Method(0).Call(nil)   // private method can't call
	pv.Field(0).SetInt(88)
	//pv.Field(1).SetString("qq") //private Field ...
	fmt.Println(pv)
}

type Person struct {
	Id int
	name string
}

func (p Person) Sayhi() {
	fmt.Println(p.name,p.Id)
}
