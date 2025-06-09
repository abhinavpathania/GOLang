package main
import "fmt"
import "unicode/utf8"

func data_types(){
	fmt.Println("Hello world")

	var myNum int = 1
	fmt.Println(myNum)

	myNum2 := 2
	fmt.Println(myNum2)

	mystring:="Hello dost"
	mystring2:=`Hello
dost how are you ?`
	fmt.Println(mystring)
	fmt.Println(mystring2)

	fmt.Println(len(mystring)) // This gives the number of bytes in the  string
	fmt.Println(utf8.RuneCountInString(mystring)) // this function gives the number of characters in the string.
	
	const myConstant int = 101
	fmt.Println(myConstant)
	// myConstant =101 Gives error
}
func main(){
	data_types()
}
