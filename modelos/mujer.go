package modelos

type Mujer struct {
	Hombre
}

func (mujer *Mujer) Respirar() string {
	return "La mujer respira"
}
func (mujer *Mujer) Comer() string {
	return "La mujer está comiendo"
}
func (mujer *Mujer) Pensar() string {
	return "La mujer está pensado"
}
func (mujer *Mujer) Saludar() string {
	return "Hola, soy una mujer"
}
func (mujer *Mujer) Sexo() string {
	return "Femenino"
}
