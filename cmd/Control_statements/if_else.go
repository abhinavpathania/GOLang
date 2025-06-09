package main

import (
	"errors"
	"fmt"
)

var numerator int
var denominator int

func take_value_from_user()(int, int){
	fmt.Println("Enter the numerator: ")
	fmt.Scan(&numerator)
	fmt.Println("Enter the denominator: ")
	fmt.Scan(&denominator)
	return numerator,denominator
}


func main() {
	
	numerator,denominator=take_value_from_user()
	div_result, remainder, err := intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf("%s", err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is: %v", div_result)
	} else {
		fmt.Printf("The result of the integer division is: %v and remainder is: %v", div_result, remainder)
	}

	result:=function_to_add(numerator,denominator)
	fmt.Printf("\nThe sum of the numerator and denominator is: %v \n",result)

	switch_statement(remainder)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("error!! Division by 0 not possible")
		return 0, 0, err
	}
	div_result := numerator / denominator
	remainder := numerator % denominator
	return div_result, remainder, err
}
