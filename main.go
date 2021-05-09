package main

import (
	"fmt"

	"godxvincent.com/go-interfaces/models"
)

type Saludos2 struct {
	*string
	int
}

func main() {

	fmt.Println("2. Implementación patron builder")
	// Implementación patron builder
	// Se crea un nuevo builder
	userBuilder := models.NewUserBuilder()
	// Con el builder se crea un usuario por default y un celular.
	user, err := userBuilder.WithCellphone(3114467265).Build()
	// Se valida si hay error, si no lo hay se imprimen los valores con los que se creo el usuario
	if err == nil {
		fmt.Println(user)
		fmt.Println(user.GetUserName())
		fmt.Println(user.GetEmail())
		fmt.Println(user.GetCellphone())
	}

	// Uso de una estructura con campos anonimos
	fmt.Println("3. Estructuras con tipos")
	var nombre = "Ricardo"
	saludo2 := &Saludos2{&nombre, 33}
	fmt.Println(saludo2)
	fmt.Println(saludo2.int)

	// Uso de interfaces como parametros
	fmt.Println("4. Funciones e Interfaces (implementación 1)")
	saludo := &models.Saludos{}
	models.ProbarInterfaceSaludos(saludo)

	fmt.Println("4. Funciones e Interfaces (implementación 2)")
	saludoEmbebido := &models.SaludosInterfaceEmbebida{}
	models.ProbarInterfaceSaludos(saludoEmbebido)

	// Uso de la interface vacias interface{} para recibir cualquier type
	fmt.Println("5. Interface generica")
	models.ProbarInterfaceGenerica(user)

}
