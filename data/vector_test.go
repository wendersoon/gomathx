package data

import (
	"math"
	"reflect"
	"testing"
)

func TestVectorLen(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector[int]
		expected int
	}{
		{"Empty vector", Vector[int]{Element: []int{}}, 0},
		{"Single element", Vector[int]{Element: []int{1}}, 1},
		{"Multiple elements", Vector[int]{Element: []int{1, 2, 3, 4, 5}}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vector.Len(); got != tt.expected {
				t.Errorf("Len() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestVectorSum(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector[int]
		expected int
	}{
		{"Empty vector", Vector[int]{Element: []int{}}, 0},
		{"Single element", Vector[int]{Element: []int{5}}, 5},
		{"Positive numbers", Vector[int]{Element: []int{1, 2, 3, 4, 5}}, 15},
		{"Mixed numbers", Vector[int]{Element: []int{-2, -1, 0, 1, 2}}, 0},
		{"Negative numbers", Vector[int]{Element: []int{-1, -2, -3}}, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vector.Sum(); got != tt.expected {
				t.Errorf("Sum() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestVectorSumFloat(t *testing.T) {
	v := Vector[float64]{Element: []float64{1.5, 2.5, 3.0}}
	expected := 7.0
	if got := v.Sum(); got != expected {
		t.Errorf("Sum() = %v, want %v", got, expected)
	}
}

func TestVectorMean(t *testing.T) {
	tests := []struct {
		name        string
		vector      Vector[int]
		expected    float64
		expectError bool
	}{
		{"Empty vector", Vector[int]{Element: []int{}}, 0, true},
		{"Single element", Vector[int]{Element: []int{10}}, 10.0, false},
		{"Multiple elements", Vector[int]{Element: []int{1, 2, 3, 4, 5}}, 3.0, false},
		{"Mixed numbers", Vector[int]{Element: []int{-2, -1, 0, 1, 2}}, 0.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.vector.Mean()
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			if got != tt.expected {
				t.Errorf("Mean() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestVectorMax(t *testing.T) {
	tests := []struct {
		name        string
		vector      Vector[int]
		expected    int
		expectError bool
	}{
		{"Empty vector", Vector[int]{Element: []int{}}, 0, true},
		{"Single element", Vector[int]{Element: []int{42}}, 42, false},
		{"Multiple elements", Vector[int]{Element: []int{1, 5, 3, 9, 2}}, 9, false},
		{"All same", Vector[int]{Element: []int{7, 7, 7}}, 7, false},
		{"Negative numbers", Vector[int]{Element: []int{-5, -1, -10}}, -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.vector.Max()
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			if got != tt.expected {
				t.Errorf("Max() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestVectorMin(t *testing.T) {
	tests := []struct {
		name        string
		vector      Vector[int]
		expected    int
		expectError bool
	}{
		{"Empty vector", Vector[int]{Element: []int{}}, 0, true},
		{"Single element", Vector[int]{Element: []int{42}}, 42, false},
		{"Multiple elements", Vector[int]{Element: []int{1, 5, 3, 9, 2}}, 1, false},
		{"All same", Vector[int]{Element: []int{7, 7, 7}}, 7, false},
		{"Negative numbers", Vector[int]{Element: []int{-5, -1, -10}}, -10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.vector.Min()
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			if got != tt.expected {
				t.Errorf("Min() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestVectorNormalize(t *testing.T) {
	tests := []struct {
		name        string
		vector      Vector[int]
		expectError bool
	}{
		{"Zero vector", Vector[int]{Element: []int{0, 0, 0}}, true},
		{"Valid vector", Vector[int]{Element: []int{3, 4}}, false},
		{"Single element", Vector[int]{Element: []int{5}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.vector.Normalize()
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Check if normalized vector has unit length
			var sumSquares float64
			for _, val := range got.Element {
				sumSquares += val * val
			}
			norm := math.Sqrt(sumSquares)
			if math.Abs(norm-1.0) > 1e-10 {
				t.Errorf("Normalized vector norm = %v, want 1.0", norm)
			}
		})
	}
}

func TestVectorNormalizeSpecific(t *testing.T) {
	// Test specific case: [3, 4] should normalize to [0.6, 0.8]
	v := Vector[int]{Element: []int{3, 4}}
	got, err := v.Normalize()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	expected := []float64{0.6, 0.8}
	for i, val := range got.Element {
		if math.Abs(val-expected[i]) > 1e-10 {
			t.Errorf("Normalize()[%d] = %v, want %v", i, val, expected[i])
		}
	}
}

func TestVectorClone(t *testing.T) {
	original := Vector[int]{Element: []int{1, 2, 3, 4, 5}}
	cloned := original.Clone()

	// Check if elements are the same
	if !reflect.DeepEqual(original.Element, cloned.Element) {
		t.Errorf("Clone() elements don't match original")
	}

	// Check if they are different slices (deep copy)
	cloned.Element[0] = 999
	if original.Element[0] == 999 {
		t.Errorf("Clone() created shallow copy instead of deep copy")
	}
}

func TestVectorReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Empty vector", []int{}, []int{}},
		{"Single element", []int{1}, []int{1}},
		{"Two elements", []int{1, 2}, []int{2, 1}},
		{"Multiple elements", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector[int]{Element: make([]int, len(tt.input))}
			copy(v.Element, tt.input)
			v.Reverse()

			if !reflect.DeepEqual(v.Element, tt.expected) {
				t.Errorf("Reverse() = %v, want %v", v.Element, tt.expected)
			}
		})
	}
}

func TestVectorAbs(t *testing.T) {
	v := Vector[int]{Element: []int{-3, -1, 0, 2, -5}}
	got := v.Abs()
	expected := []float64{3, 1, 0, 2, 5}

	if !reflect.DeepEqual(got.Element, expected) {
		t.Errorf("Abs() = %v, want %v", got.Element, expected)
	}
}

func TestVectorScale(t *testing.T) {
	v := Vector[int]{Element: []int{1, 2, 3, 4}}
	scalar := 3
	got := v.Scale(scalar)
	expected := []int{3, 6, 9, 12}

	if !reflect.DeepEqual(got.Element, expected) {
		t.Errorf("Scale(%d) = %v, want %v", scalar, got.Element, expected)
	}
}

func TestVectorScaleFloat(t *testing.T) {
	v := Vector[float64]{Element: []float64{1.5, 2.0, 2.5}}
	scalar := 2.0
	got := v.Scale(scalar)
	expected := []float64{3.0, 4.0, 5.0}

	if !reflect.DeepEqual(got.Element, expected) {
		t.Errorf("Scale(%f) = %v, want %v", scalar, got.Element, expected)
	}
}

func TestVectorApply(t *testing.T) {
	v := Vector[int]{Element: []int{1, 2, 3, 4}}
	square := func(x int) int { return x * x }
	got := v.Apply(square)
	expected := []int{1, 4, 9, 16}

	if !reflect.DeepEqual(got.Element, expected) {
		t.Errorf("Apply(square) = %v, want %v", got.Element, expected)
	}
}

func TestVectorCumsum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Empty vector", []int{}, []int{}},
		{"Single element", []int{5}, []int{5}},
		{"Multiple elements", []int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{"With negatives", []int{1, -1, 2, -2}, []int{1, 0, 2, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector[int]{Element: tt.input}
			got := v.Cumsum()

			if !reflect.DeepEqual(got.Element, tt.expected) {
				t.Errorf("Cumsum() = %v, want %v", got.Element, tt.expected)
			}
		})
	}
}

func TestVectorDiff(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Empty vector", []int{}, []int{}},
		{"Single element", []int{5}, []int{}},
		{"Two elements", []int{5, 8}, []int{3}},
		{"Multiple elements", []int{1, 3, 6, 10}, []int{2, 3, 4}},
		{"Decreasing", []int{10, 8, 5, 1}, []int{-2, -3, -4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector[int]{Element: tt.input}
			got := v.Diff()

			if !reflect.DeepEqual(got.Element, tt.expected) {
				t.Errorf("Diff() = %v, want %v", got.Element, tt.expected)
			}
		})
	}
}

func TestVectorArgMax(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Empty vector", []int{}, -1},
		{"Single element", []int{5}, 0},
		{"Multiple elements", []int{1, 5, 3, 9, 2}, 3},
		{"First element max", []int{9, 1, 2, 3}, 0},
		{"Last element max", []int{1, 2, 3, 9}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector[int]{Element: tt.input}
			got := v.ArgMax()

			if got != tt.expected {
				t.Errorf("ArgMax() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestVectorArgMin(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Empty vector", []int{}, -1},
		{"Single element", []int{5}, 0},
		{"Multiple elements", []int{5, 1, 9, 3, 2}, 1},
		{"First element min", []int{1, 5, 9, 3}, 0},
		{"Last element min", []int{5, 9, 3, 1}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector[int]{Element: tt.input}
			got := v.ArgMin()

			if got != tt.expected {
				t.Errorf("ArgMin() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestVectorSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Empty vector", []int{}, []int{}},
		{"Single element", []int{5}, []int{5}},
		{"Already sorted", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"Reverse sorted", []int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		{"Random order", []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector[int]{Element: make([]int, len(tt.input))}
			copy(v.Element, tt.input)
			v.Sort()

			if !reflect.DeepEqual(v.Element, tt.expected) {
				t.Errorf("Sort() = %v, want %v", v.Element, tt.expected)
			}
		})
	}
}

func TestVectorStdDev(t *testing.T) {
	tests := []struct {
		name      string
		input     []float64
		expected  float64
		tolerance float64
	}{
		{"Empty vector", []float64{}, 0.0, 0.0},
		{"Single element", []float64{5.0}, 0.0, 0.0},
		{"All same elements", []float64{3.0, 3.0, 3.0}, 0.0, 1e-10},
		{"Simple case", []float64{1.0, 2.0, 3.0, 4.0, 5.0}, 1.4142135623730951, 1e-10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector[float64]{Element: tt.input}
			got := v.StdDev()

			if math.Abs(got-tt.expected) > tt.tolerance {
				t.Errorf("StdDev() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestVectorUnique(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Empty vector", []int{}, []int{}},
		{"No duplicates", []int{1, 2, 3}, []int{1, 2, 3}},
		{"With duplicates", []int{1, 2, 2, 3, 1, 4}, []int{1, 2, 3, 4}},
		{"All same", []int{5, 5, 5, 5}, []int{5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector[int]{Element: tt.input}
			got := v.Unique()

			// Since maps don't guarantee order, we need to check if all expected elements are present
			if len(got.Element) != len(tt.expected) {
				t.Errorf("Unique() length = %v, want %v", len(got.Element), len(tt.expected))
				return
			}

			// Check if all expected elements are present
			expectedMap := make(map[int]bool)
			for _, val := range tt.expected {
				expectedMap[val] = true
			}

			for _, val := range got.Element {
				if !expectedMap[val] {
					t.Errorf("Unique() contains unexpected element %v", val)
				}
			}
		})
	}
}

// Benchmark tests
func BenchmarkVectorSum(b *testing.B) {
	v := Vector[int]{Element: make([]int, 1000)}
	for i := range v.Element {
		v.Element[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.Sum()
	}
}

func BenchmarkVectorSort(b *testing.B) {
	original := make([]int, 1000)
	for i := range original {
		original[i] = 1000 - i // Reverse order
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := Vector[int]{Element: make([]int, len(original))}
		copy(v.Element, original)
		b.StartTimer()

		v.Sort()
	}
}

// Test different numeric types
func TestVectorWithDifferentTypes(t *testing.T) {
	// Test with float32
	vFloat32 := Vector[float32]{Element: []float32{1.1, 2.2, 3.3}}
	if sum := vFloat32.Sum(); math.Abs(float64(sum-6.6)) > 1e-6 {
		t.Errorf("float32 Sum() = %v, want ~6.6", sum)
	}

	// Test with uint
	vUint := Vector[uint]{Element: []uint{1, 2, 3, 4}}
	if sum := vUint.Sum(); sum != 10 {
		t.Errorf("uint Sum() = %v, want 10", sum)
	}

	// Test with int64
	vInt64 := Vector[int64]{Element: []int64{100, 200, 300}}
	if sum := vInt64.Sum(); sum != 600 {
		t.Errorf("int64 Sum() = %v, want 600", sum)
	}
}
