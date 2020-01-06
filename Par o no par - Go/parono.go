package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("******* Mi programa ***********")

	numero, _ := strconv.Atoi(os.Args[1])

	if numero%2 == 0 {
		fmt.Println("Numero es PAR")
	} else {
		fmt.Println("Numero es IMPAR")
	}

}
