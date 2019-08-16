package main

import "fmt"

func main() {
	//switch case con un simple caso
	name := "suitable"
	switch name {
	case "sui":
		fmt.Println("sui")
	case "suitable":
		fmt.Println("suitable")
	default:
		fmt.Println("Ninguno de los casos anteriores")
	} 
	//mutiples casos que realizan una misma operacion
	num := 2
	switch num {
	case 1,2,3:
		fmt.Println("El numero esta entre 1 y 3")
    case 4,5,6:
		fmt.Println("El numero esta entre 4 y 6")
	}
	//para que continue en los siguientes casos se debe usar falltrough
}