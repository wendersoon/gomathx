// Package gomathx provides high-performance mathematical operations for Go.
//
// GoMathX is a comprehensive mathematical library featuring generic vectors,
// matrices, and linear algebra operations. It's designed for scientific computing,
// data analysis, and machine learning applications.
//
// # Key Features
//
//   - Generic vectors supporting all numeric types
//   - Comprehensive statistical operations
//   - Element-wise vector arithmetic
//   - Performance-optimized algorithms
//   - Zero external dependencies
//
// # Quick Example
//
//	package main
//
//	import (
//		"fmt"
//		"log"
//
//		"github.com/wendersoon/gomathx/vector"
//	)
//
//	func main() {
//		vec, err := vector.CreateVector([]int{1, 2, 3, 4, 5})
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Printf("Sum: %d\n", vec.Sum())
//		mean, _ := vec.Mean()
//		fmt.Printf("Mean: %.2f\n", mean)
//	}
//
// # Subpackages
//
//   - data: Core data structures and vector operations
//   - vector: Vector creation and arithmetic operations
//   - matrix: Matrix operations (coming soon)
package gomathx
