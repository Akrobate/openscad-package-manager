package renderer

import "fmt"

type Renderer struct{}

func NewRenderer() (*Renderer, error) {

	return &Renderer{}, nil
}

func (r *Renderer) List() error {
	fmt.Println("Renderer List...")
	return nil
}
