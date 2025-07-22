package vector

import (
	"errors"
	"math"

	"github.com/wendersoon/gomathx/data"
)

// CreateVector creates a new generic vector
func CreateVector[T data.Number](slice []T) (*data.Vector[T], error) {
	if len(slice) == 0 {
		return nil, ErrEmptyVector
	}
	return &data.Vector[T]{Element: slice}, nil
}

// AddVectors performs element-wise addition on two or more vectors.
// All vectors must have the same length. Returns an error if fewer than two vectors
// are provided or if their lengths do not match.func AddVectors[T data.Number](vectors ...*data.Vector[T]) (*data.Vector[T], error) {
func AddVectors[T data.Number](vectors ...*data.Vector[T]) (*data.Vector[T], error) {

	if len(vectors) < 2 {
		return nil, errors.New("need at least two vectors to add")
	}
	length := vectors[0].Len()
	for _, v := range vectors {
		if v.Len() != length {
			return nil, ErrMismatchedLengths
		}
	}

	result := make([]T, length)
	for _, v := range vectors {
		for i, val := range v.Element {
			result[i] += val
		}
	}
	return &data.Vector[T]{Element: result}, nil
}

// SubVectors performs element-wise subtraction between two vectors (a - b).
// Returns an error if the vectors have different lengths.
func SubVectors[T data.Number](a, b *data.Vector[T]) (*data.Vector[T], error) {
	if a.Len() != b.Len() {
		return nil, ErrMismatchedLengths
	}
	result := make([]T, a.Len())
	for i := 0; i < a.Len(); i++ {
		result[i] = a.Element[i] - b.Element[i]
	}
	return &data.Vector[T]{Element: result}, nil
}

// MulVectors performs element-wise multiplication between two vectors.
// Returns an error if the vectors have different lengths.
func MulVectors[T data.Number](a, b *data.Vector[T]) (*data.Vector[T], error) {
	if a.Len() != b.Len() {
		return nil, ErrMismatchedLengths
	}
	result := make([]T, a.Len())
	for i := 0; i < a.Len(); i++ {
		result[i] = a.Element[i] * b.Element[i]
	}
	return &data.Vector[T]{Element: result}, nil
}

// DivVectors performs element-wise division between two vectors (a / b).
// Returns an error if the vectors have different lengths or if any element in b is zero.
func DivVectors[T data.Number](a, b *data.Vector[T]) (*data.Vector[T], error) {
	if a.Len() != b.Len() {
		return nil, ErrMismatchedLengths
	}
	result := make([]T, a.Len())
	for i := 0; i < a.Len(); i++ {
		if b.Element[i] == 0 {
			return nil, errors.New("division by zero")
		}
		result[i] = a.Element[i] / b.Element[i]
	}
	return &data.Vector[T]{Element: result}, nil
}

// DotProduct computes the dot product (scalar product) of two vectors.
// Returns an error if vectors have different lengths.
func DotProduct[T data.Number](a, b *data.Vector[T]) (T, error) {
	if a.Len() != b.Len() {
		return 0, ErrMismatchedLengths
	}
	var result T
	for i := 0; i < a.Len(); i++ {
		result += a.Element[i] * b.Element[i]
	}
	return result, nil
}

// EuclideanDistance calculates the Euclidean distance between two vectors.
// Returns an error if vectors have different lengths.
func EuclideanDistance[T data.Number](a, b *data.Vector[T]) (float64, error) {
	if a.Len() != b.Len() {
		return 0, ErrMismatchedLengths
	}
	var sum float64
	for i := 0; i < a.Len(); i++ {
		diff := float64(a.Element[i] - b.Element[i])
		sum += diff * diff
	}
	return math.Sqrt(sum), nil
}

// CosineSimilarity calculates the cosine similarity between two vectors.
// Returns a value between -1 and 1, or an error on invalid input.
func CosineSimilarity[T data.Number](a, b *data.Vector[T]) (float64, error) {
	if a.Len() != b.Len() {
		return 0, ErrMismatchedLengths
	}

	var dot, normA, normB float64
	for i := 0; i < a.Len(); i++ {
		ai := float64(a.Element[i])
		bi := float64(b.Element[i])
		dot += ai * bi
		normA += ai * ai
		normB += bi * bi
	}

	if normA == 0 || normB == 0 {
		return 0, errors.New("cosine similarity undefined for zero-length vector")
	}

	return dot / (math.Sqrt(normA) * math.Sqrt(normB)), nil
}

// EqualVectors checks if all elements in the vectors are equal.
// Returns false if lengths are different or any elements mismatch.
func EqualVectors[T data.Number](a, b *data.Vector[T]) bool {
	if a.Len() != b.Len() {
		return false
	}
	for i := 0; i < a.Len(); i++ {
		if a.Element[i] != b.Element[i] {
			return false
		}
	}
	return true
}

// ElementWiseMax returns a vector containing the element-wise maximum values of two vectors.
func ElementWiseMax[T data.Number](a, b *data.Vector[T]) (*data.Vector[T], error) {
	if a.Len() != b.Len() {
		return nil, ErrMismatchedLengths
	}
	result := make([]T, a.Len())
	for i := 0; i < a.Len(); i++ {
		if a.Element[i] > b.Element[i] {
			result[i] = a.Element[i]
		} else {
			result[i] = b.Element[i]
		}
	}
	return &data.Vector[T]{Element: result}, nil
}

// ElementWiseMin returns a vector containing the element-wise minimum values of two vectors.
func ElementWiseMin[T data.Number](a, b *data.Vector[T]) (*data.Vector[T], error) {
	if a.Len() != b.Len() {
		return nil, ErrMismatchedLengths
	}
	result := make([]T, a.Len())
	for i := 0; i < a.Len(); i++ {
		if a.Element[i] < b.Element[i] {
			result[i] = a.Element[i]
		} else {
			result[i] = b.Element[i]
		}
	}
	return &data.Vector[T]{Element: result}, nil
}
