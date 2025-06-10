package main

import "fmt"

func maps(){

	var myMap map[string]uint8 = make(map[string]uint8)
	myMap["Abhinav"]=21
	myMap["Chandu"]=62

	fmt.Printf("Value of 1st element is %v \n",myMap["Abhinav"])

	fmt.Println(myMap["Josh"]) // if does not exist return the default value of key(uint8)

	myMap["Palin"]=20
	myMap["Parth"]=20
	myMap["Anmol"]=20

	delete(myMap,"Abhinav")

	for key,value := range myMap{
		fmt.Printf("%s %d \n",key,value)
	}

	val,exist := myMap["Abhinav"]
	if exist{
		fmt.Printf("It exists and value is %v",val)
	}else{
		fmt.Println("Does not exist")
	}
	

}