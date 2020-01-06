package main

import "fmt"

func main() {
	// slides son los arreglos que pueden crecer de forma dinamica

	peliculas := []string{
		"El joker",
		"Star Wars",
		"Lucha de poder",
		"Isla siniestra"}

	peliculas = append(peliculas, "Sin limites")
	peliculas = append(peliculas, "Camp Rock")

	fmt.Println(peliculas[5])

	// determinar el largo de un array
	fmt.Println(len(peliculas))

	// determinar cantidad de elementos a mostrar
	fmt.Println(peliculas[0:3])

}
