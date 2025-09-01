package interfaces

type Humano interface {
	Saludar() string
	Pensar() string
	Respirar() string
	Comer() string
	Sexo() string
	EstaVivo() bool
}
