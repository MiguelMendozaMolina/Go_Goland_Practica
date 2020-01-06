package main

import (
	"fmt"
	// "time"
)

func main() {
	// time.Sleep(time.Second * 5)

	// Declaracion de variables
	var suma int = 8 + 8
	var resta int = 6 - 4
	var nombre string = "Miguel"
    var apellidos string = "Mendoza Molina"
    
    // Asignacion de valor en tiempo de ejecucion 
    apellidos = "Minco"

   // Segunda deficion de variables
    pais := "Chile"

    // Declaracion de boleanos 
    var prueba bool = true
    fmt.Println(prueba)

    // Declaracion de numeros decimales
    var numero float32 = 3.12322323
    fmt.Println(numero)

    // declaracion de una constante
    // no se puede cambiar durante su tiempo de ejecucion 
    const year int = 2019

    
    fmt.Println(year)
	fmt.Println("Hola Mundo desde Go - Sunneo " + nombre + apellidos + pais)
	fmt.Println(suma)
	fmt.Println(resta)
}
