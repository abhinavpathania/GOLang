package main

import (
	"fmt"
	"strings"
)

func processRune(){
	mystring := []rune("résumé")
	var indexed = mystring[1]
	fmt.Printf("%v \t %T \n",indexed,indexed)

	for i,v := range mystring{
		fmt.Printf("%v \t %v \n",i,v)
	}

	myrune := []rune{'a','b','c'}
	var strBuilder strings.Builder
	for i := range myrune{
		strBuilder.WriteString(string(myrune[i]))
	}
	var catstring = strBuilder.String()

	fmt.Printf("\n %v",catstring)
}