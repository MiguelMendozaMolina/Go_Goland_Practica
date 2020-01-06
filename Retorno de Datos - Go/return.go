package main

import "fmt"

func main() {
	fmt.Println(devolverTexto())
}

// forma de escribir funciones
/*func devolverTexto() (string, string) {

	dato1 := "Miguel"
	dato2 := "Mendoza"
	return dato1, dato2
} */

// forma de escribir funciones mas eficaz
func devolverTexto() (dato1 string, dato2 string) {

	dato1 = "Miguel"
	dato2 = "Mendoza"
	return
}
