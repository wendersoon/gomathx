// data/doc.go
// Package data provides core mathematical data structures for GoMathX.
//
// This package contains the fundamental Vector type and associated operations
// for mathematical computations. All operations are generic and support
// any numeric type through the Number interface constraint.
//
// The Vector type provides comprehensive mathematical operations including:
//   - Basic statistics (sum, mean, min, max, standard deviation)
//   - Vector transformations (normalize, scale, sort, reverse)
//   - Sequential operations (cumsum, diff)
//   - Functional operations (apply, unique)
//
// Example:
//
//	vec := &data.Vector[int]{Element: []int{1, 2, 3, 4, 5}}
//	sum := vec.Sum()                    // 15
//	mean, _ := vec.Mean()               // 3.0
//	normalized, _ := vec.Normalize()    // Unit vector
package data
