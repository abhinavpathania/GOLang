package main

import "fmt"

func main(){
	// var mypointer *int32 = new(int32)
	// var i int32
	// mypointer=&i
	// i=10 // is similar to *mypointer = 10
	// fmt.Println(*mypointer)
	myarray:=[5]int32{1,2,3,4,5}
	fmt.Printf("Memory address of array in main block %p \n",&myarray)
	result := square(&myarray)
	fmt.Printf("%v",result)
}

func square(myarray *[5]int32)([5]int32){
	fmt.Printf("Memory address of array in func     block %p \n",myarray)
	for i:= range myarray{
		myarray[i]=myarray[i]*myarray[i]
	}
	return *myarray
}