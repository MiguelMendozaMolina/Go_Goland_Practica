package main

import "fmt"

func main(){
	var numero1 int = 10;
	var numero2 int = 6;
   
    //Suma
	fmt.Print("La suma es: ");
	fmt.Println(numero1 + numero2);

	// Resta
	fmt.Print("La resta es: ");
	fmt.Println(numero1 - numero2);

	// Multiplicacion
	fmt.Print("La mulplicacion es: ");
	fmt.Println(numero1 * numero2);

	//Division
	fmt.Print("La division es: ");
	fmt.Println(numero1 / numero2); 

	// Resto de la division
	fmt.Print("El resto de la division es: ");
	fmt.Println(numero1 % numero2);  
}