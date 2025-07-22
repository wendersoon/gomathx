// vector/doc.go
// Package vector provides factory functions and arithmetic operations for vectors.
//
// This package offers convenient factory functions for creating vectors
// and performing element-wise arithmetic operations between vectors.
// All operations maintain type safety through Go's generics system.
//
// Key functions include:
//   - CreateVector: Safe vector creation with validation
//   - AddVectors: Element-wise addition of multiple vectors
//   - SubVectors: Element-wise subtraction
//   - MulVectors: Element-wise multiplication
//   - DivVectors: Element-wise division with zero-check
//
// All arithmetic operations require vectors of equal length and will
// return ErrMismatchedLengths if dimensions don't match.
//
// Example:
//
//	vec1, _ := vector.CreateVector([]int{1, 2, 3})
//	vec2, _ := vector.CreateVector([]int{4, 5, 6})
//	sum, _ := vector.AddVectors(vec1, vec2)  // [5, 7, 9]
package vector
