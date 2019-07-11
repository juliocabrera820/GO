package main

import "fmt"

func main() {
	//Los strings son inmutables e indexables
	lengu := "JavaScript"
	//Longitud de la cadena
	fmt.Println(len(lengu))
	//Sustraer cadena
	//el rango es desde 0 hasta un caracter antes del rango, este caso seria 3
	fmt.Println(lengu[0:4])
	//Concatenar con el operador +
	lengu=lengu[0:4]+" EE"
	fmt.Println(lengu)
	//Imprimir con comillas dobles
	fmt.Println("\"Mensaje con comillas\"")

}