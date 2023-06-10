// Packege main - Ejercicios con variables y tipos de declaraciones
package main

import "fmt"

func main() {
	// Declaración de variables independiente a la asignacion de las mismas
	var apple string
	apple = "🍎"
	var lemon string
	lemon = "🍋"
	var banana string
	banana = "🍌"
	fmt.Println("Declaración de variables 1 ", apple, lemon, banana)

	// Declaración de variables por bloques
	var (
		apple2  string = "🍎"
		lemon2  string = "🍋"
		banana2 string = "🍌"
	)
	fmt.Println()
	fmt.Println("Declaración de variables 2 ", apple2, lemon2, banana2)

	// Declaración de variables e inicializacion en una misma linea
	var apple3, lemon3, banana3 string = "🍎", "🍋", "🍌"

	fmt.Println()
	fmt.Println("Declaración de variables 3", apple3, lemon3, banana3)

	// Declaración de variables con el operador conrto de asignacion;
	// en este caso el tipo de dato es inferido por go y no puede cambiar
	// durante la ejecucion del programa
	apple4, lemon4, banana4 := "🍎", "🍋", "🍌"

	// Con el operador de asignacion corto es posible reasignar valores y
	// declarar nuevas variables en la misma linea
	apple4, orange := "🍎", "🍊"

	fmt.Println()
	fmt.Println("Declaración de variables 4", apple4, lemon4, banana4, orange)
}
