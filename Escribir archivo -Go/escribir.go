package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	nuevoTexto := os.Args[1] + "\n"
	//escribir := ioutil.WriteFile("fichero.txt", nuevoTexto, 07777)
	//showError(escribir)

	fichero, err := os.OpenFile("fichero.txt", os.O_APPEND, 0777)
	showError(err)

	escribir, err := fichero.WriteString(nuevoTexto)
	fmt.Println(escribir)
	showError(err)

	fichero.Close()

	texto, errorDeFichero := ioutil.ReadFile("fichero.txt")
	showError(errorDeFichero)

	fmt.Println(string(texto))

}

func showError(e error) {
	if e != nil {
		panic(e)
	}

}
