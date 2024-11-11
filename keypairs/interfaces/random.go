package interfaces

type Randomizer interface {
	GenerateBytes(n int) ([]byte, error)
}