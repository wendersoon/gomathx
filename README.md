# GoMathX

[![Go Reference](https://pkg.go.dev/badge/github.com/wendersoon/gomathx.svg)](https://pkg.go.dev/github.com/wendersoon/gomathx)
[![Go Report Card](https://goreportcard.com/badge/github.com/wendersoon/gomathx)](https://goreportcard.com/report/github.com/wendersoon/gomathx)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A high-performance mathematical library for Go featuring generic vectors, matrices, and linear algebra operations. GoMathX is designed to provide efficient and type-safe mathematical computations for scientific computing, data analysis, and machine learning applications.

## 🚀 Features

## Features

- **Generic Vectors**: Type-safe vector operations supporting all numeric types
- **Comprehensive Vector Operations**: Sum, mean, max, min, normalization, sorting, and more
- **Vector Arithmetic**: Element-wise addition, subtraction, multiplication, and division
- **Statistical Functions**: Standard deviation, cumulative sum, differences, and unique values
- **Performance Optimized**: Efficient algorithms with benchmark tests
- **Zero Dependencies**: Pure Go implementation with no external dependencies
- **Upcoming Features**: Matrix operations and linear algebra (coming soon)

## Installation

```bash
go get github.com/wendersoon/gomathx
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/wendersoon/gomathx/vector"
)

func main() {
    // Create a vector
    vec, err := vector.CreateVector([]int{1, 2, 3, 4, 5})
    if err != nil {
        log.Fatal(err)
    }
    
    // Basic operations
    fmt.Printf("Length: %d\n", vec.Len())           // Length: 5
    fmt.Printf("Sum: %d\n", vec.Sum())             // Sum: 15
    
    mean, _ := vec.Mean()
    fmt.Printf("Mean: %.2f\n", mean)               // Mean: 3.00
    
    max, _ := vec.Max()
    fmt.Printf("Max: %d\n", max)                   // Max: 5
}
```

## API Documentation

### Vector Package

#### Creating Vectors

```go
// CreateVector creates a new vector from a slice
vec, err := vector.CreateVector([]float64{1.5, 2.5, 3.5})
if err != nil {
    // Handle error (returns ErrEmptyVector for empty slices)
}
```

#### Basic Vector Operations

##### Length and Sum
```go
vec := &data.Vector[int]{Element: []int{1, 2, 3, 4, 5}}

length := vec.Len()     // Returns: 5
sum := vec.Sum()        // Returns: 15
```

##### Statistical Operations
```go
// Mean (returns float64 and error)
mean, err := vec.Mean()
if err != nil {
    // Handle error (empty vector)
}

// Standard deviation
stddev := vec.StdDev()

// Min and Max values
min, err := vec.Min()
max, err := vec.Max()

// Index of min and max values
minIdx := vec.ArgMin()  // Returns -1 for empty vector
maxIdx := vec.ArgMax()  // Returns -1 for empty vector
```

##### Vector Transformations
```go
// Clone vector (deep copy)
cloned := vec.Clone()

// Reverse vector in-place
vec.Reverse()

// Sort vector in-place (ascending order)
vec.Sort()

// Normalize vector (returns Vector[float64])
normalized, err := vec.Normalize()

// Absolute values (returns Vector[float64])
absVec := vec.Abs()

// Scale by scalar
scaled := vec.Scale(2)  // Multiply each element by 2

// Apply function to each element
squared := vec.Apply(func(x int) int { return x * x })

// Get unique elements
unique := vec.Unique()
```

##### Sequential Operations
```go
// Cumulative sum
cumsum := vec.Cumsum()  // [1,2,3,4] -> [1,3,6,10]

// Differences between consecutive elements
diff := vec.Diff()      // [1,3,6,10] -> [2,3,4]
```

#### Vector Arithmetic

GoMathX provides element-wise operations between vectors:

```go
// Create vectors
vec1, _ := vector.CreateVector([]int{1, 2, 3, 4})
vec2, _ := vector.CreateVector([]int{5, 6, 7, 8})
vec3, _ := vector.CreateVector([]int{2, 2, 2, 2})

// Addition (supports multiple vectors)
sum, err := vector.AddVectors(vec1, vec2, vec3)
// Result: [8, 10, 12, 14]

// Subtraction
diff, err := vector.SubVectors(vec2, vec1)
// Result: [4, 4, 4, 4]

// Element-wise multiplication
product, err := vector.MulVectors(vec1, vec2)
// Result: [5, 12, 21, 32]

// Element-wise division
quotient, err := vector.DivVectors(vec2, vec1)
// Result: [5, 3, 2, 2] (integer division)
```

**Important**: All vector arithmetic operations require vectors of the same length. Operations return `ErrMismatchedLengths` error if lengths don't match.

### Supported Numeric Types

GoMathX supports all Go numeric types through the `Number` interface:

- **Signed integers**: `int`, `int8`, `int16`, `int32`, `int64`
- **Unsigned integers**: `uint`, `uint8`, `uint16`, `uint32`, `uint64`  
- **Floating-point**: `float32`, `float64`

```go
// Examples with different types
intVec, _ := vector.CreateVector([]int{1, 2, 3})
floatVec, _ := vector.CreateVector([]float64{1.5, 2.5, 3.5})
float32Vec, _ := vector.CreateVector([]float32{1.1, 2.2, 3.3})
uint64Vec, _ := vector.CreateVector([]uint64{100, 200, 300})
```

## Examples

### Example 1: Basic Statistical Analysis

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/wendersoon/gomathx/vector"
)

func main() {
    // Sample data
    data := []float64{10.5, 12.3, 8.7, 15.2, 11.1, 9.8, 13.4, 14.6, 7.9, 16.3}
    
    vec, err := vector.CreateVector(data)
    if err != nil {
        log.Fatal(err)
    }
    
    // Statistical analysis
    fmt.Printf("Dataset: %v\n", vec.Element)
    fmt.Printf("Length: %d\n", vec.Len())
    
    mean, _ := vec.Mean()
    fmt.Printf("Mean: %.2f\n", mean)
    
    min, _ := vec.Min()
    max, _ := vec.Max()
    fmt.Printf("Range: %.1f - %.1f\n", min, max)
    
    stddev := vec.StdDev()
    fmt.Printf("Standard Deviation: %.2f\n", stddev)
    
    // Sort and find median
    sorted := vec.Clone()
    sorted.Sort()
    median := sorted.Element[sorted.Len()/2]
    fmt.Printf("Median: %.1f\n", median)
}
```

### Example 2: Data Transformation Pipeline

```go
package main

import (
    "fmt"
    "log"
    "math"
    
    "github.com/wendersoon/gomathx/vector"
)

func main() {
    // Raw data
    rawData := []float64{-2.5, 1.3, -0.8, 3.2, -1.1, 2.7, -3.4, 0.9}
    
    vec, err := vector.CreateVector(rawData)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Original: %v\n", vec.Element)
    
    // Step 1: Get absolute values
    absVec := vec.Abs()
    fmt.Printf("Absolute: %v\n", absVec.Element)
    
    // Step 2: Apply transformation (square root)
    sqrtVec := absVec.Apply(func(x float64) float64 {
        return math.Sqrt(x)
    })
    fmt.Printf("Square Root: %v\n", sqrtVec.Element)
    
    // Step 3: Normalize
    normalized, err := sqrtVec.Normalize()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Normalized: %v\n", normalized.Element)
    
    // Verify normalization (should be ~1.0)
    normSum := 0.0
    for _, val := range normalized.Element {
        normSum += val * val
    }
    fmt.Printf("Norm verification: %.10f\n", math.Sqrt(normSum))
}
```

### Example 3: Vector Arithmetic Operations

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/wendersoon/gomathx/vector"
)

func main() {
    // Create vectors
    prices := []float64{100.0, 150.0, 200.0, 75.0}
    quantities := []float64{2.0, 1.0, 3.0, 4.0}
    discounts := []float64{0.1, 0.15, 0.2, 0.05} // 10%, 15%, 20%, 5%
    
    priceVec, _ := vector.CreateVector(prices)
    qtyVec, _ := vector.CreateVector(quantities)
    discountVec, _ := vector.CreateVector(discounts)
    
    // Calculate subtotal (price * quantity)
    subtotal, err := vector.MulVectors(priceVec, qtyVec)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Subtotals: %v\n", subtotal.Element)
    
    // Calculate discount amounts
    discountAmounts, err := vector.MulVectors(subtotal, discountVec)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Discount amounts: %v\n", discountAmounts.Element)
    
    // Calculate final totals
    finalTotals, err := vector.SubVectors(subtotal, discountAmounts)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Final totals: %v\n", finalTotals.Element)
    
    // Grand total
    grandTotal := finalTotals.Sum()
    fmt.Printf("Grand total: %.2f\n", grandTotal)
}
```

### Example 4: Time Series Analysis

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/wendersoon/gomathx/vector"
)

func main() {
    // Daily stock prices
    prices := []float64{100, 102, 98, 105, 107, 103, 109, 112, 108, 115}
    
    priceVec, err := vector.CreateVector(prices)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Prices: %v\n", priceVec.Element)
    
    // Calculate daily changes
    changes := priceVec.Diff()
    fmt.Printf("Daily changes: %v\n", changes.Element)
    
    // Calculate cumulative returns (starting from 0)
    cumulativeReturns := changes.Cumsum()
    fmt.Printf("Cumulative changes: %v\n", cumulativeReturns.Element)
    
    // Find best and worst days
    maxIdx := changes.ArgMax()
    minIdx := changes.ArgMin()
    
    maxChange, _ := changes.Max()
    minChange, _ := changes.Min()
    
    fmt.Printf("Best day: Day %d with change +%.1f\n", maxIdx+2, maxChange)
    fmt.Printf("Worst day: Day %d with change %.1f\n", minIdx+2, minChange)
    
    // Calculate volatility (standard deviation of changes)
    volatility := changes.StdDev()
    fmt.Printf("Volatility (std dev): %.2f\n", volatility)
}
```

## Error Handling

GoMathX provides specific error types for different scenarios:

```go
import "github.com/wendersoon/gomathx/vector"

// ErrEmptyVector - returned when trying to create empty vector
_, err := vector.CreateVector([]int{})
if err == vector.ErrEmptyVector {
    fmt.Println("Cannot create empty vector")
}

// ErrMismatchedLengths - returned for vector operations with different lengths
vec1, _ := vector.CreateVector([]int{1, 2, 3})
vec2, _ := vector.CreateVector([]int{4, 5})
_, err = vector.AddVectors(vec1, vec2)
if err == vector.ErrMismatchedLengths {
    fmt.Println("Vectors must have same length")
}

// Methods that can return errors
mean, err := vec.Mean()           // Empty vector error
normalized, err := vec.Normalize() // Zero vector error
max, err := vec.Max()             // Empty vector error
min, err := vec.Min()             // Empty vector error
```

## Performance

GoMathX is designed for performance with:

- **Efficient memory usage**: Minimal allocations and memory copying
- **Optimized algorithms**: Built-in Go sorting and mathematical functions
- **Benchmark tests**: Comprehensive performance testing included

### Running Benchmarks

```bash
cd data
go test -bench=.
```

Sample benchmark results:
```
BenchmarkVectorSum-8     3000000    500 ns/op
BenchmarkVectorSort-8     100000   15000 ns/op
```

## Project Structure

```
gomathx/
├── data/                    # Core data structures
│   ├── number.go           # Number interface constraint
│   ├── vector.go           # Vector implementation
│   └── vector_test.go      # Comprehensive tests
├── vector/                  # Vector factory and operations
│   ├── error.go            # Error definitions
│   ├── factory.go          # Vector creation and arithmetic
│   └── factory_test.go     # Factory tests
├── matrix/                  # Matrix operations (coming soon)
├── go.mod                   # Module definition
├── LICENSE                  # License file
└── README.md               # This file
```

## Contributing

We welcome contributions! Please feel free to submit issues, feature requests, or pull requests.

### Development Guidelines

1. **Testing**: All new features must include comprehensive tests
2. **Benchmarks**: Performance-critical code should include benchmarks
3. **Documentation**: Public APIs must be well-documented
4. **Code Style**: Follow standard Go conventions and formatting

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
cd data && go test -v
cd vector && go test -v
```

## Roadmap

- ✅ **Generic Vector Operations** - Complete
- ✅ **Vector Arithmetic** - Complete  
- ✅ **Statistical Functions** - Complete
- 🚧 **Matrix Operations** - In Development
- 📋 **Linear Algebra** - Planned
- 📋 **Complex Number Support** - Planned
- 📋 **Sparse Vectors/Matrices** - Planned
- 📋 **BLAS Integration** - Under Consideration

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with Go's powerful generics system
- Inspired by NumPy and other mathematical libraries
- Designed for the Go community's scientific computing needs

---

**GoMathX** - Empowering Go developers with efficient mathematical computations.