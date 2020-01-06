package main

import (
	"fmt"
)

func main() {

	//numero := 26

	peliculas := []string{"El joker", "Civil war", "El ilusionista", "sin limites"}

	// Realizacion de un ciclo for normal

	for i := 0; i < len(peliculas); i++ {
		if i%2 == 0 {
			fmt.Println("La pelicula ->"+peliculas[i]+"es par", i)
		} else {
			fmt.Println("La pelicula -> "+peliculas[i]+"es impar", i)
		}
	}

	// Creacion de foreach , recorrer un array

	for _, pelicula := range peliculas {
		fmt.Println(pelicula)
	}

}
