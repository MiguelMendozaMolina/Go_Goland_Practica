package main

import (
	"fmt"
	"time"
)

func main() {

	momento := time.Now()

	hoy := momento.Weekday()

	switch hoy {
	case 0:
		fmt.Println("Hoy es domingo")

	case 1:
		fmt.Println("Hoy es lunes")

	case 2:
		fmt.Println("Hoy es martes")

	case 3:
		fmt.Println("Hoy es miercoles")

	case 4:
		fmt.Println("Hoy es jueves")

	case 5:
		fmt.Println("Hoy es viernes")

	case 6:
		fmt.Println("Hoy es sabado")

	default:
		fmt.Println("no se encuentra dia de la semana")

	}

}
