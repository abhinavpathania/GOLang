package main

import "fmt"

func switch_statement(remainder int){
switch remainder{
case 0:
	fmt.Printf("Wao no remainder")
case 1,2:
	fmt.Printf("Division was so close")
default:
	fmt.Printf("Haagu division")
}
}