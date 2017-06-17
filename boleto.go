package main

import (
	"fmt"
	"os"
)

func main() {
	importe := 45
	acum := 0

	for {
		moneda := 0
		fmt.Print("Ingresar moneda: ")
		fmt.Scan(&moneda)
		fmt.Println("tu moneda es de ", moneda)
		acum = acum + int(moneda)

		if acum <= importe {
			{
				break
			}
		} else if acum >= importe {
			vuelto := acum - importe
			fmt.Println("Te imprimo el boleto")
			fmt.Println("Tu vuelto es:", vuelto)
			os.Exit(1)
		} else {
			fmt.Println("Te imprimo el boleto")
			os.Exit(1)
		}
	}
}
