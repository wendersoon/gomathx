package vector

import "github.com/wendersoon/gomathx/data"

// CreateVector creates a new generic vector
func CreateVector[T data.Number](slice []T) (*data.Vector[T], error) {
	if len(slice) == 0 {
		return nil, ErrEmptyVector
	}
	return &data.Vector[T]{Element: slice}, nil
}
