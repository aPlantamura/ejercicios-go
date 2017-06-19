package main

import (
	"fmt"
	"os"
)

func main() {
	// Declaramos el importe
	importe := 75
	// Declaramos el contador (acum) como un int que vale 0
	acum := 0

	// Hacemos un loop (algoritmo de repetición)
	for {
		// Ponemos el monto de la moneda ingresada en 0. Esto es necesario porque
		// cada vez que vuelve a empezar el loop no queremos que moneda sea el
		// numero ingresado anteriormente
		moneda := 0
		// Ingresamos la moneda
		fmt.Print("Ingresar moneda: ")
		fmt.Scan(&moneda)
		// Sumamos la moneda ingresada al contador
		acum = acum + int(moneda)

		// Si el acumulado es exactamente igual al importe imprimimos el boleto
		if acum == importe {
			fmt.Println("Te imprimo el boleto")
			os.Exit(1)
			// Si es mayor al importe, imprimimos el voleto y damos vuelto
		} else if acum > importe {
			vuelto := acum - importe
			fmt.Println("Te imprimo el boleto")
			fmt.Println("Tu vuelto es:", vuelto)
			os.Exit(1)
		}
		// Si no es ni igual ni mayor al importe dejamos que el loop continue
	}
}
