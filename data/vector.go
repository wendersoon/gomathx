package data

import (
	"errors"
	"math"
	"slices"
)

// Vector represents a generic vector
type Vector[T Number] struct {
	Element []T
}

// Len returns the number of elements in the vector
func (v *Vector[T]) Len() int {
	return len(v.Element)
}

// Sum returns the number of elements in the vector
func (v *Vector[T]) Sum() T {
	var total T
	for _, val := range v.Element {
		total += val
	}
	return total
}

// Mean returns the average value of the vector
func (v *Vector[T]) Mean() (float64, error) {
	if v.Len() == 0 {
		return 0, errors.New("cannot calculate mean of empty vector")
	}
	return float64(v.Sum()) / float64(v.Len()), nil
}

// Max returns the maximum value in the vector
func (v *Vector[T]) Max() (T, error) {
	if v.Len() == 0 {
		var zero T
		return zero, errors.New("cannot calculate max of empty vector")
	}
	max := v.Element[0]
	for _, val := range v.Element[1:] {
		if val > max {
			max = val
		}
	}
	return max, nil
}

// Min returns the minimum value in the vector
func (v *Vector[T]) Min() (T, error) {
	if v.Len() == 0 {
		var zero T
		return zero, errors.New("cannot calculate min of empty vector")
	}
	min := v.Element[0]
	for _, val := range v.Element[1:] {
		if val < min {
			min = val
		}
	}
	return min, nil
}

// Normalize returns a new vector scaled to unit length (Euclidean norm = 1)
func (v *Vector[T]) Normalize() (*Vector[float64], error) {
	var sumSquares float64
	for _, val := range v.Element {
		fVal := float64(val)
		sumSquares += fVal * fVal
	}

	if sumSquares == 0 {
		return nil, errors.New("cannot normalize zero vector")
	}

	norm := math.Sqrt(sumSquares)
	normalized := make([]float64, v.Len())
	for i, val := range v.Element {
		normalized[i] = float64(val) / norm
	}

	return &Vector[float64]{Element: normalized}, nil
}

// Clone returns a copy of the current vector
func (v *Vector[T]) Clone() *Vector[T] {
	cloned := make([]T, v.Len())
	copy(cloned, v.Element)
	return &Vector[T]{Element: cloned}
}

// Reverse reverses the elements in place
func (v *Vector[T]) Reverse() {
	for i, j := 0, v.Len()-1; i < j; i, j = i+1, j-1 {
		v.Element[i], v.Element[j] = v.Element[j], v.Element[i]
	}
}

// Abs returns a new a vector with absolute values (only meaningful for signed types)
func (v *Vector[T]) Abs() *Vector[float64] {
	absVec := make([]float64, v.Len())
	for i, val := range v.Element {
		absVec[i] = math.Abs(float64(val))
	}

	return &Vector[float64]{Element: absVec}
}

// Scale returns a new vector where each element is multiplied by a scalar
func (v *Vector[T]) Scale(scalar T) *Vector[T] {
	scaled := make([]T, v.Len())
	for i, val := range v.Element {
		scaled[i] = val * scalar
	}
	return &Vector[T]{Element: scaled}
}

// Apply applies a function to each element and returns a new vector
func (v *Vector[T]) Apply(f func(T) T) *Vector[T] {
	result := make([]T, v.Len())
	for i, val := range v.Element {
		result[i] = f(val)
	}
	return &Vector[T]{Element: result}
}

// Cumsum returns a new vector where each element is the cumulative sum
func (v *Vector[T]) Cumsum() *Vector[T] {
	result := make([]T, v.Len())
	var sum T
	for i, val := range v.Element {
		sum += val
		result[i] = sum
	}
	return &Vector[T]{Element: result}
}

// Diff returns a new vector with differences between consecutive elements
func (v *Vector[T]) Diff() *Vector[T] {
	if v.Len() < 2 {
		return &Vector[T]{Element: []T{}}
	}
	diff := make([]T, v.Len()-1)
	for i := 1; i < v.Len(); i++ {
		diff[i-1] = v.Element[i] - v.Element[i-1]
	}
	return &Vector[T]{Element: diff}
}

// ArgMax returns the index of the maximum value in the vector
func (v *Vector[T]) ArgMax() int {
	if v.Len() == 0 {
		return -1
	}
	maxIdx := 0
	for i := 1; i < v.Len(); i++ {
		if v.Element[i] > v.Element[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}

// ArgMin returns the index of the minimum value in the vector
func (v *Vector[T]) ArgMin() int {
	if v.Len() == 0 {
		return -1
	}
	minIdx := 0
	for i := 1; i < v.Len(); i++ {
		if v.Element[i] < v.Element[minIdx] {
			minIdx = i
		}
	}
	return minIdx
}

// Sort sorts the vector in ascending order
func (v *Vector[T]) Sort() {
	slices.Sort(v.Element)
}

// StdDev returns the standard deviation of the vector elements
func (v *Vector[T]) StdDev() float64 {
	n := float64(v.Len())
	if n == 0 {
		return 0
	}
	var sum, mean, sqDiff float64
	for _, val := range v.Element {
		sum += float64(val)
	}
	mean = sum / n
	for _, val := range v.Element {
		diff := float64(val) - mean
		sqDiff += diff * diff
	}
	return math.Sqrt(sqDiff / n)
}

// Unique returns a new vector with unique elements
func (v *Vector[T]) Unique() *Vector[T] {
	seen := make(map[T]struct{})
	unique := make([]T, 0, v.Len())
	for _, val := range v.Element {
		if _, ok := seen[val]; !ok {
			seen[val] = struct{}{}
			unique = append(unique, val)
		}
	}
	return &Vector[T]{Element: unique}
}
