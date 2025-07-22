package vector

import "errors"

// ErrEmptyVector is returned when trying to create an empty vector
var ErrEmptyVector = errors.New("empty vector is not allowed")

var ErrMismatchedLengths = errors.New("vectors must have the same length")
