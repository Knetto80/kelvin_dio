package main //função principal//

import "fmt" //importação de função//

// K=C+273 formula temperatura de Graus Celsius em Kelvin

const C = 100
var K float64

func main()  { //função principal//
	K = C + 273
	fmt.Println("O ponto de ebulição da agua em Kelvin é de :",K,"ºK")
}