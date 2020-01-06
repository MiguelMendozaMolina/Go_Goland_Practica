package main

import "fmt"

func main() {
	fmt.Println("Pedido 1")
	fmt.Println(gorras(5000, "pesos"))
	fmt.Println("Pedido 2")
	fmt.Println(gorras(4500, "pesos"))

}

func gorras(pedido float32, moneda string) (string, float32, string) {

	// clousure una variable que tendra una funcion dentro

	precio := func() float32 {
		return pedido * 7
	}

	return "El precio del pedido:", precio(), moneda
}
