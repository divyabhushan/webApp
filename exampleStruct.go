package main

import (
	"fmt"
)

type myStruct struct {
	name   string
	id     int
	salary float64
}
type mineStruct struct {
	myStruct //anonymous field (instance of structure myStruct)
	name     string
}

func main() {
	//	var MyVar = 300
	//st := myStruct{"greets", 20, 80000}
	//fmt.Println(st)
	//st.id = 100
	//fmt.Println(st.name, st.id, st.salary)
	my := mineStruct{myStruct{"saurabh", 23, 430000}, "Aashi"}
	fmt.Println("mineStrunct has the name:", my.name)
	fmt.Println("myStruct has the name:", my.myStruct.name)
}
