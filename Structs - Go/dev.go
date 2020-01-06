package main

import "fmt"

type Gorra struct{
	marca string
	color string
	precio float32
	plana bool
}

func main(){

  /*var gorra_negra  = Gorra {
	  marca: "Nike",
	  color: "negro",
	  precio: 25.50,
	  plana: false}
  
	fmt.Println(gorra_negra)
	fmt.Println(gorra_negra.marca) */


    fmt.Println("Calculadora1");
	var numero1 float32 = 10;
	var numero2 float32 = 6;
	calculadora(numero1, numero2);

	fmt.Println("===================");
	
	fmt.Println("Calculadora2")
	var numero3 float32 = 4;
	var numero4 float32 = 12;
	calculadora(numero3, numero4);
	holaMundo();

	
}

func holaMundo (){
  fmt.Println("Hola Mundo")
}

func operacion(n1 float32,n2 float32,op string) float32{
 var resultado float32;
 if(op == "+"){
	 resultado = n1 + n2
 }

 if(op == "-"){
	resultado = n1 - n2
}

if(op == "*"){
	resultado = n1 * n2
}

if(op == "/"){
	resultado = n1 / n2
}
 return resultado
}

func calculadora(numero1 float32 , numero2 float32){
	//Suma
	fmt.Print("La suma es: ");
	fmt.Println(operacion(numero1, numero2, "+"));

	// Resta
	fmt.Print("La resta es: ");
	fmt.Println(operacion(numero1, numero2, "-"));

	// Multiplicacion
	fmt.Print("La mulplicacion es: ");
	fmt.Println(operacion(numero1, numero2, "*"));

	//Division
	fmt.Print("La division es: ");
	fmt.Println(operacion(numero1, numero2, "/")); 

	// Resto de la division
	//fmt.Print("El resto de la division es: ");
	//fmt.Println(numero1 % numero2);  
}