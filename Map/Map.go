package main

import "fmt"

func main() {
	//Mapa con clave de tipo string y valores enteros
	m := make(map[string]int)
	m["key1"]=2
	fmt.Println(m["key1"])
	//Mapa con clave de tipo int y valores de tipo string
	mi := make(map[int]string)
	mi[1]= "Value"
	mi[2]= "Value2"
	fmt.Println(mi[1])
	//Eliminar item con la clave
	delete(mi,1)
	for _, it := range mi {
		fmt.Println(it)
	}
}