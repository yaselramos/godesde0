package interfaces

type Animales interface {
	Comer() string
	Respirar() string
	EsCarnivoro() bool
	EstaVivo() bool
}
