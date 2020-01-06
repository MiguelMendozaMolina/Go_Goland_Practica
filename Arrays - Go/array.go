package main

import "fmt"

func main() {
	// primera forma de definir un array
	/*
		var peliculas [3]string

		peliculas[0] = "El joker"
		peliculas[1] = "El jajas"
		peliculas[2] = "Civil War"

		fmt.Println(peliculas)
	*/

	// segunda forma de definir un array

	peliculas := [3]string{"el jajas",
		"eljoker",
		"civil war"}
	fmt.Println(peliculas)

}
