package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	texto, errorDeFichero := ioutil.ReadFile("fichero.txt")
	showError(errorDeFichero)

	fmt.Println(string(texto))

}

func showError(e error) {
	if e != nil {
		panic(e)
	}

}
