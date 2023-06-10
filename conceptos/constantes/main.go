// Packege main - Ejercicios con constantes y tipos de declaraciones
package main

import "fmt"

// Las constantes son comunmente declaradas a nivel de paquete
const (
	os    = "macOS"
	model = "Mac Book Pro"
)

// las constantes contrario a las variables no necesariamente deben ser utilizadas
// iota es un valor autoincrementable inicia en 0 y se puede utilizar para
// asignar valores numericos secuenciales
const (
	jun = iota + 1
	feb
	mar
	apr
	may
)

func main() {
	fmt.Println(os, model)

	fmt.Println(jun, feb)
}
