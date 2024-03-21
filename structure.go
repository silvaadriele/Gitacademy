package main

import "fmt"

func main() {
	type Coordenada struct {
		latitude  float64
		longitude float64
	}
	coordenada := Coordenada{latitude: 37.42, longitude: 122.4}

	fmt.Printf("%+v", coordenada)
	fmt.Println(coordenada)
	fmt.Println(coordenada.latitude)
}
