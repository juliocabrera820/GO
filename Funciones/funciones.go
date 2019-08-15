package main
import "fmt"
//Funcion de tipo void con parametro string
func mostrarMensaje(msj string ) {
	fmt.Println(msj)
}
//Funcion con dos parametros que retorna un entero
func suma(num int,num2 int) int {
	return num+num2
}
//Funciones que devuelve 2 valores
func mess(num int) (string,int)  {
	return "mensaje",5
}
func main() {
	mostrarMensaje("Global scope")
	fmt.Println("Suma",suma(3,9))
	value,str := mess(5)
	fmt.Println("El",str,"es el",value)

}