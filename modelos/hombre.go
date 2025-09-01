package modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	comiendo   bool
	pensando   bool
	vivo       bool
}

func (hombre *Hombre) Respirar() string {
	return "La hombre respira"
}

func (hombre *Hombre) Pensar() string {
	return "La hombre está pensando"
}
func (hombre *Hombre) Saludar() string {
	return "Hola, soy un hombre"
}
func (hombre *Hombre) Sexo() string {
	return "Masculino"
}
func (hombre *Hombre) Comer() string {
	return "El hombre está comiendo"
}
