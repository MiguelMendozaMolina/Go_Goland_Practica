package main

import "fmt"

func main() {

	var peliculas [3][2]string
	peliculas[0][0] = "El jajas"
	peliculas[0][1] = "El joker"

	peliculas[1][0] = "Civil War"
	peliculas[1][1] = "End Game"

	peliculas[2][0] = "El aviador"
	peliculas[2][1] = "transpoting"

	fmt.Println(peliculas)

}
