package main

import (
	"fmt"
)

func main(){
	var str string =  "résumé"
	var index = str[1]
	fmt.Printf("%v \t %T \n",index,index)

	for i,v := range str{
		fmt.Printf("%v \t %v \n",i,v)
	}
	processRune()
}
