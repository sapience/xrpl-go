package random

import (
	"crypto/rand"
	"io"
)

type Randomizer struct {
	io.Reader
}

func NewRandomizer() *Randomizer {
	return &Randomizer{
		Reader: rand.Reader,
	}
}

func (r *Randomizer) GenerateBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := r.Read(b) //nolint
	if err != nil {
		return nil, err
	}
	return b, nil
}
