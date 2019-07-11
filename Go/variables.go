package main

import "fmt"

func main() {
	//Maneras de asignar
	//Con 'var', seguido del nombre de la variable y el tipo, ademas del operador = y el valor
	var num int = 1
	fmt.Println(num)
	//Asignar con el operador :=,se define el tipo de la variable
	nom := "RUBY"
	//Si con el operador := asignamos un valor de tipo cadena, no se puede agregar posteriormente un dato de diferente tipo
	nom = "Golang"
	fmt.Println(nom)
	//Tambien se puede asignar 2 valores al mismo tiempo
	num,nom = 10,"Python"
	fmt.Println("Numero =>",num," Nombre",nom)
	java,swift := "Java","Swift"
	fmt.Println(java,swift)
	//Intercambiar valores de las variables de mismo tipo
	swift,java = "java","swift"
	fmt.Println("Valores cambiados\n",java,swift)
	//Asignar multiples variables con var y =
	var (
		len = "C"
		poo = "PHP"
		fun = "racket"
	)
	fmt.Println(len,poo,fun)
	//Variables locales y globales, las variables que se declaran fuera del main, son globlales y se deben declar con var nombredelatributo = valor
}

