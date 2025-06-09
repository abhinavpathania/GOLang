package main

import "fmt"

func slice(){

	mySlice := []int32{1,2,3,4}

	fmt.Println(mySlice[0])
	fmt.Printf("The length is %v and capacity is %v\n",len(mySlice),cap(mySlice))
	mySlice = append(mySlice, 7 , 8)
	fmt.Printf("The length is %v and capacity is %v\n",len(mySlice),cap(mySlice))

	mySlice2 := []int32{5,6}
	mySlice2 =append(mySlice,mySlice2...)

	fmt.Println(mySlice2)
}



