package models

type Saludos struct {
}

func (s *Saludos) Generico() string {
	return "Hola desde la implementación de la struct Saludos"
}

func (s *Saludos) Personalizado(nombre string) string {
	return "Hola " + nombre + " la implementación de la struct Saludos"
}

type SaludosInterfaceEmbebida struct {
	SaludosI
}

func (s *SaludosInterfaceEmbebida) Generico() string {
	return "Hola desde la implementación de la struct SaludosInterfaceEmbebida"
}
