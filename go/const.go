package main 

import "fmt"

func main(){
  const FEMALE=1
  const MALE=0
  const (
    a = iota
    b = iota<<3
    c 
    d
    e
  )
  fmt.Printf("FEMALE:%d , MALE:%d\n",FEMALE,MALE)
  fmt.Println(a,b,c,d,e)
}
