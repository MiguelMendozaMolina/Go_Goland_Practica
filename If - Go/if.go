package main

import (
	"fmt"
	"os"
	"strconv"
)

// strconv = libreria para pasar una variable a entero
// os = leer parametros desde la linea de comandos
func main() {
	fmt.Println("***** MI PROGRAMA CON GO **********")

	// Convertir de un arreglo a dato accesible

	// acceder a los argumentos por consola
	fmt.Println("Hola " + os.Args[1] + " Bienvenido al programa en go")

	edad, _ := strconv.Atoi(os.Args[2])

	// and &&

	//	if edad >= 18 || edad == 17 && edad != 25 && edad != 99 {
	if edad >= 18 && edad != 25 && edad != 99 {
		fmt.Println("Eres mayor de edad")
	} else if edad == 25 {
		fmt.Println("Tienes 25 años")
	} else if edad == 99 {
		fmt.Println("Eres un abuelo")
	} else {
		fmt.Println("Eres menor de edad")
	}
}
