package random

import "io"

type Randomizer struct {
	io.Reader
}

func (r *Randomizer) GenerateBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := r.Read(b) //nolint
	if err != nil {
		return nil, err
	}
	return b, nil
}
