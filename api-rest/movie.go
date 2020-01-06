package main

type Movie struct {
	Nombre   string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

type Movies []Movie
