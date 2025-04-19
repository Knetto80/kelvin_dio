package main //função principal//

import "fmt" //importação de função//

// K=C+273 formula temperatura de Graus Celsius para Kelvin//

const C = 100 //temperatura em ºC //
var K float64 // K kelvin //

func main()  { //função principal//
	K = C + 273
	fmt.Println("O ponto de ebulição da agua em Kelvin é de :",K,"ºK")
}