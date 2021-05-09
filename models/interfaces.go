package models

import "fmt"

type SaludosI interface {
	Generico() string
	Personalizado(string) string
}

func ProbarInterfaceSaludos(interfaz SaludosI) {
	fmt.Println(interfaz.Generico())
}

func ProbarInterfaceGenerica(interfaz interface{}) {
	fmt.Println(interfaz)
}
