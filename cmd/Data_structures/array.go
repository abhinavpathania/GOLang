package main

import "fmt"

func main(){
	
	fmt.Println(`
	1. Arrays
	2. Slice
	3. Maps
	4. Measure time`)
	var valuex int
	fmt.Println("\t Enter the choice")
	fmt.Scan(&valuex)
	switch valuex{
	case 1:
		array()
	case 2:
		slice()
	case 3:
		maps()
	case 4:
		measure_time()
	default:
		fmt.Println("Invalid !!")
	}
}

func array(){
	var arr1[2] int32
	fmt.Println(arr1[0])

	arr2:= [...]int32{1,2,3,4}
	fmt.Println(arr2[3])

	arr2[0]=62
	fmt.Println(arr2[0])


	fmt.Printf("%v",&arr2[0])
}