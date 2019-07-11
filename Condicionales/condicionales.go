package main

import "fmt"

func main() {
	a,b := 5,6
	if a<b {
		fmt.Println("a es menor que b")
	}else{
		fmt.Println("b es menor que a")
	}
	//En las condiciones se puede realizar sentencias y despues verficar la condicion
	if a+=2; a > b{
		fmt.Println("a es mayor que b y tiene un valor de",a)
	}else{
		fmt.Println("a es menor que b")
	}
}