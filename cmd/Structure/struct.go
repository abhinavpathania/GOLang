package main

import "fmt"

// Declaration of the Structure
type gasEngine struct{ 
	mpg uint16
	gallons uint16
}

type electricEngine struct{
	mpkwh uint16
	kwh uint16
}
//Interface

type engine interface{
	Gasleft() uint16
}

// Meathod

func (e gasEngine) Gasleft()(uint16){
	return e.gallons*e.mpg
}

func (e electricEngine) Gasleft()(uint16){
	return e.mpkwh*e.kwh
}


func Canmakeit(e engine, miles uint16){
	if e.Gasleft()>=miles{
		fmt.Println("Yes you can make it there !")
	}else{
		fmt.Println("Need to Fuel up !!")

	}
} 

func main(){
	var myGasEngine gasEngine = gasEngine{10,20}
	fmt.Println(myGasEngine)

	var myElectricEngine electricEngine = electricEngine{10,30}
	fmt.Println(myElectricEngine)


	fmt.Printf("Gas left in Gas engine is %v \n",myGasEngine.Gasleft())
	fmt.Printf("Gas left in Electric engine is %v \n",myElectricEngine.Gasleft())
	var miles uint16 = 210
	fmt.Println("Can u make it ?")
	Canmakeit(myGasEngine,miles)
	Canmakeit(myElectricEngine,miles)
}