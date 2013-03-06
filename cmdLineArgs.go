package main

import (
	"fmt"
	"os"
)

//checking the number of command-line arguments passed

func main() {
	//fmt.Println("Please input 4 parameters...")
	fmt.Println("Command line arguments :", os.Args[1:])
	fmt.Println("Program name is:", os.Args[0])
	length := len(os.Args[1:])
	fmt.Println("Length of cmd-line arguments:", length)
	if length != 4 {
		fmt.Println("Usage: [  ] [ 0 1 2 3 ]")
		os.Exit(1)
	}
	fmt.Println("You entered parameters... :",os.Args[1:])

}
