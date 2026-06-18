package main

import "fmt"

const C float64 = 0      //Ponto de fusão da água em Celsius
const F float64 = 32     //Ponto de fusão da água em Fahrenheit
const K float64 = 273.15 //Ponto de fusão da água em Kelvin
var Converte float64 = 0 //Variável para Conversão de escalas de temperatura

func main() {
	var tempCelsius float64 = 100 //Temperatura em Celsius a ser convertida

	//Conversão de Celsius para Fahrenheit
	Converte = (tempCelsius * 9 / 5) + F
	fmt.Printf("%.2f graus Celsius equivalem a %.2f graus Fahrenheit\n", tempCelsius, Converte)

	//Conversão de Celsius para Kelvin
	Converte = tempCelsius + K
	fmt.Printf("%.2f graus Celsius equivalem a %.2f graus Kelvin\n", tempCelsius, Converte)

	//Conversão de Fahrenheit para Celsius
	Converte = (tempCelsius - F) * 5 / 9
	fmt.Printf("%.2f graus Fahrenheit equivalem a %.2f graus Celsius\n", tempCelsius, Converte)

	//Conversão de Kelvin para Celsius
	Converte = tempCelsius - K
	fmt.Printf("%.2f graus Kelvin equivalem a %.2f graus Celsius\n", tempCelsius, Converte)

	//Conversão de Fahrenheit para Kelvin
	Converte = (tempCelsius-F)*5/9 + K
	fmt.Printf("%.2f graus Fahrenheit equivalem a %.2f graus Kelvin\n", tempCelsius, Converte)

	//Conversão de Kelvin para Fahrenheit
	Converte = (tempCelsius-K)*9/5 + F
	fmt.Printf("%.2f graus Kelvin equivalem a %.2f graus Fahrenheit\n", tempCelsius, Converte)
}
