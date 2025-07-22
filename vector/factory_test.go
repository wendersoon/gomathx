package vector_test

import (
	"math"
	"testing"

	"github.com/wendersoon/gomathx/data"
	"github.com/wendersoon/gomathx/vector"
)

// TestCreateVector tests the CreateVector function
func TestCreateVector_Success(t *testing.T) {
	vec, err := vector.CreateVector([]int{1, 2, 3})

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if vec == nil {
		t.Error("expected non-nil vector")
	}

	if vec.Len() != 3 {
		t.Errorf("expected length 3, got: %d", vec.Len())
	}
}

func TestCreateVector_EmptySlice(t *testing.T) {
	vec, err := vector.CreateVector([]int{})

	if err == nil {
		t.Fatal("expected error for empty slice, got nil")
	}

	if err != vector.ErrEmptyVector {
		t.Errorf("expected ErrEmptyVector, got: %v", err)
	}

	if vec != nil {
		t.Errorf("expected nil vector, got: %+v", vec)
	}
}

func TestCreateVector_DifferentTypes(t *testing.T) {
	// Test with float64
	vecFloat, err := vector.CreateVector([]float64{1.5, 2.5, 3.5})
	if err != nil {
		t.Errorf("expected no error for float64, got: %v", err)
	}
	if vecFloat.Len() != 3 {
		t.Errorf("expected length 3 for float64, got: %d", vecFloat.Len())
	}

	// Test with int32
	vecInt32, err := vector.CreateVector([]int32{10, 20, 30})
	if err != nil {
		t.Errorf("expected no error for int32, got: %v", err)
	}
	if vecInt32.Len() != 3 {
		t.Errorf("expected length 3 for int32, got: %d", vecInt32.Len())
	}
}

// TestAddVectors tests the AddVectors function
func TestAddVectors_Success(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{4, 5, 6})

	result, err := vector.AddVectors(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expected := []int{5, 7, 9}
	for i, val := range result.Element {
		if val != expected[i] {
			t.Errorf("expected %d at index %d, got %d", expected[i], i, val)
		}
	}
}

func TestAddVectors_MultipleVectors(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{4, 5, 6})
	vec3, _ := vector.CreateVector([]int{7, 8, 9})

	result, err := vector.AddVectors(vec1, vec2, vec3)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expected := []int{12, 15, 18}
	for i, val := range result.Element {
		if val != expected[i] {
			t.Errorf("expected %d at index %d, got %d", expected[i], i, val)
		}
	}
}

func TestAddVectors_SingleVector(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})

	_, err := vector.AddVectors(vec1)

	if err == nil {
		t.Error("expected error for single vector")
	}

	expectedMsg := "need at least two vectors to add"
	if err.Error() != expectedMsg {
		t.Errorf("expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestAddVectors_MismatchedLengths(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{4, 5})

	_, err := vector.AddVectors(vec1, vec2)

	if err == nil {
		t.Error("expected error for mismatched lengths")
	}

	if err != vector.ErrMismatchedLengths {
		t.Errorf("expected ErrMismatchedLengths, got: %v", err)
	}
}

func TestAddVectors_Float(t *testing.T) {
	vec1, _ := vector.CreateVector([]float64{1.5, 2.5, 3.5})
	vec2, _ := vector.CreateVector([]float64{0.5, 1.5, 2.5})

	result, err := vector.AddVectors(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expected := []float64{2.0, 4.0, 6.0}
	for i, val := range result.Element {
		if val != expected[i] {
			t.Errorf("expected %f at index %d, got %f", expected[i], i, val)
		}
	}
}

// TestSubVectors tests the SubVectors function
func TestSubVectors_Success(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{5, 7, 9})
	vec2, _ := vector.CreateVector([]int{1, 2, 3})

	result, err := vector.SubVectors(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expected := []int{4, 5, 6}
	for i, val := range result.Element {
		if val != expected[i] {
			t.Errorf("expected %d at index %d, got %d", expected[i], i, val)
		}
	}
}

func TestSubVectors_MismatchedLengths(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{4, 5})

	_, err := vector.SubVectors(vec1, vec2)

	if err != vector.ErrMismatchedLengths {
		t.Errorf("expected ErrMismatchedLengths, got: %v", err)
	}
}

func TestSubVectors_NegativeResults(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{5, 7, 9})

	result, err := vector.SubVectors(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expected := []int{-4, -5, -6}
	for i, val := range result.Element {
		if val != expected[i] {
			t.Errorf("expected %d at index %d, got %d", expected[i], i, val)
		}
	}
}

// TestMulVectors tests the MulVectors function
func TestMulVectors_Success(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{2, 3, 4})
	vec2, _ := vector.CreateVector([]int{5, 6, 7})

	result, err := vector.MulVectors(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expected := []int{10, 18, 28}
	for i, val := range result.Element {
		if val != expected[i] {
			t.Errorf("expected %d at index %d, got %d", expected[i], i, val)
		}
	}
}

func TestMulVectors_WithZeros(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{2, 0, 4})
	vec2, _ := vector.CreateVector([]int{5, 6, 0})

	result, err := vector.MulVectors(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expected := []int{10, 0, 0}
	for i, val := range result.Element {
		if val != expected[i] {
			t.Errorf("expected %d at index %d, got %d", expected[i], i, val)
		}
	}
}

func TestMulVectors_MismatchedLengths(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{4, 5})

	_, err := vector.MulVectors(vec1, vec2)

	if err != vector.ErrMismatchedLengths {
		t.Errorf("expected ErrMismatchedLengths, got: %v", err)
	}
}

// TestDivVectors tests the DivVectors function
func TestDivVectors_Success(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{10, 15, 20})
	vec2, _ := vector.CreateVector([]int{2, 3, 4})

	result, err := vector.DivVectors(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expected := []int{5, 5, 5}
	for i, val := range result.Element {
		if val != expected[i] {
			t.Errorf("expected %d at index %d, got %d", expected[i], i, val)
		}
	}
}

func TestDivVectors_DivisionByZero(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{10, 15, 20})
	vec2, _ := vector.CreateVector([]int{2, 0, 4})

	_, err := vector.DivVectors(vec1, vec2)

	if err == nil {
		t.Error("expected error for division by zero")
	}

	expectedMsg := "division by zero"
	if err.Error() != expectedMsg {
		t.Errorf("expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestDivVectors_MismatchedLengths(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{10, 15, 20})
	vec2, _ := vector.CreateVector([]int{2, 3})

	_, err := vector.DivVectors(vec1, vec2)

	if err != vector.ErrMismatchedLengths {
		t.Errorf("expected ErrMismatchedLengths, got: %v", err)
	}
}

func TestDivVectors_Float(t *testing.T) {
	vec1, _ := vector.CreateVector([]float64{10.0, 15.0, 20.0})
	vec2, _ := vector.CreateVector([]float64{2.0, 3.0, 4.0})

	result, err := vector.DivVectors(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expected := []float64{5.0, 5.0, 5.0}
	for i, val := range result.Element {
		if val != expected[i] {
			t.Errorf("expected %f at index %d, got %f", expected[i], i, val)
		}
	}
}

// TestDotProduct tests the DotProduct function
func TestDotProduct_Success(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{4, 5, 6})

	result, err := vector.DotProduct(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expected := 32 // 1*4 + 2*5 + 3*6 = 4 + 10 + 18 = 32
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestDotProduct_ZeroVectors(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{0, 0, 0})
	vec2, _ := vector.CreateVector([]int{1, 2, 3})

	result, err := vector.DotProduct(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result != 0 {
		t.Errorf("expected 0, got %d", result)
	}
}

func TestDotProduct_MismatchedLengths(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{4, 5})

	_, err := vector.DotProduct(vec1, vec2)

	if err != vector.ErrMismatchedLengths {
		t.Errorf("expected ErrMismatchedLengths, got: %v", err)
	}
}

func TestDotProduct_Float(t *testing.T) {
	vec1, _ := vector.CreateVector([]float64{1.5, 2.5, 3.5})
	vec2, _ := vector.CreateVector([]float64{2.0, 3.0, 4.0})

	result, err := vector.DotProduct(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expected := 24.5 // 1.5*2.0 + 2.5*3.0 + 3.5*4.0 = 3.0 + 7.5 + 14.0 = 24.5
	if result != expected {
		t.Errorf("expected %f, got %f", expected, result)
	}
}

// TestEuclideanDistance tests the EuclideanDistance function
func TestEuclideanDistance_Success(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{4, 6, 8})

	result, err := vector.EuclideanDistance(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Distance = sqrt((4-1)² + (6-2)² + (8-3)²) = sqrt(9 + 16 + 25) = sqrt(50) ≈ 7.071
	expected := math.Sqrt(50)
	if math.Abs(result-expected) > 1e-10 {
		t.Errorf("expected %f, got %f", expected, result)
	}
}

func TestEuclideanDistance_SameVectors(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{1, 2, 3})

	result, err := vector.EuclideanDistance(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result != 0 {
		t.Errorf("expected 0, got %f", result)
	}
}

func TestEuclideanDistance_MismatchedLengths(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{4, 5})

	_, err := vector.EuclideanDistance(vec1, vec2)

	if err != vector.ErrMismatchedLengths {
		t.Errorf("expected ErrMismatchedLengths, got: %v", err)
	}
}

// TestCosineSimilarity tests the CosineSimilarity function
func TestCosineSimilarity_Success(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{4, 5, 6})

	result, err := vector.CosineSimilarity(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Cosine similarity = dot(a,b) / (||a|| * ||b||)
	// dot(a,b) = 32, ||a|| = sqrt(14), ||b|| = sqrt(77)
	// result = 32 / (sqrt(14) * sqrt(77)) ≈ 0.9746
	expected := 32.0 / (math.Sqrt(14) * math.Sqrt(77))
	if math.Abs(result-expected) > 1e-10 {
		t.Errorf("expected %f, got %f", expected, result)
	}
}

func TestCosineSimilarity_IdenticalVectors(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{1, 2, 3})

	result, err := vector.CosineSimilarity(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if math.Abs(result-1.0) > 1e-10 {
		t.Errorf("expected 1.0, got %f", result)
	}
}

func TestCosineSimilarity_ZeroVector(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{0, 0, 0})
	vec2, _ := vector.CreateVector([]int{1, 2, 3})

	_, err := vector.CosineSimilarity(vec1, vec2)

	if err == nil {
		t.Error("expected error for zero vector")
	}

	expectedMsg := "cosine similarity undefined for zero-length vector"
	if err.Error() != expectedMsg {
		t.Errorf("expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestCosineSimilarity_MismatchedLengths(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{4, 5})

	_, err := vector.CosineSimilarity(vec1, vec2)

	if err != vector.ErrMismatchedLengths {
		t.Errorf("expected ErrMismatchedLengths, got: %v", err)
	}
}

// TestEqualVectors tests the EqualVectors function
func TestEqualVectors_Equal(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{1, 2, 3})

	result := vector.EqualVectors(vec1, vec2)

	if !result {
		t.Error("expected true for equal vectors")
	}
}

func TestEqualVectors_NotEqual(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{1, 2, 4})

	result := vector.EqualVectors(vec1, vec2)

	if result {
		t.Error("expected false for non-equal vectors")
	}
}

func TestEqualVectors_DifferentLengths(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{1, 2})

	result := vector.EqualVectors(vec1, vec2)

	if result {
		t.Error("expected false for different length vectors")
	}
}

func TestEqualVectors_EmptyVectors(t *testing.T) {
	// This test assumes we can create empty vectors through some other means
	// Since CreateVector doesn't allow empty vectors, we'll create them directly
	vec1 := &data.Vector[int]{Element: []int{}}
	vec2 := &data.Vector[int]{Element: []int{}}

	result := vector.EqualVectors(vec1, vec2)

	if !result {
		t.Error("expected true for equal empty vectors")
	}
}

// TestElementWiseMax tests the ElementWiseMax function
func TestElementWiseMax_Success(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 5, 3})
	vec2, _ := vector.CreateVector([]int{4, 2, 6})

	result, err := vector.ElementWiseMax(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expected := []int{4, 5, 6}
	for i, val := range result.Element {
		if val != expected[i] {
			t.Errorf("expected %d at index %d, got %d", expected[i], i, val)
		}
	}
}

func TestElementWiseMax_MismatchedLengths(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{4, 5})

	_, err := vector.ElementWiseMax(vec1, vec2)

	if err != vector.ErrMismatchedLengths {
		t.Errorf("expected ErrMismatchedLengths, got: %v", err)
	}
}

func TestElementWiseMax_NegativeNumbers(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{-1, -5, -3})
	vec2, _ := vector.CreateVector([]int{-4, -2, -6})

	result, err := vector.ElementWiseMax(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expected := []int{-1, -2, -3}
	for i, val := range result.Element {
		if val != expected[i] {
			t.Errorf("expected %d at index %d, got %d", expected[i], i, val)
		}
	}
}

// TestElementWiseMin tests the ElementWiseMin function
func TestElementWiseMin_Success(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 5, 3})
	vec2, _ := vector.CreateVector([]int{4, 2, 6})

	result, err := vector.ElementWiseMin(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expected := []int{1, 2, 3}
	for i, val := range result.Element {
		if val != expected[i] {
			t.Errorf("expected %d at index %d, got %d", expected[i], i, val)
		}
	}
}

func TestElementWiseMin_MismatchedLengths(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{1, 2, 3})
	vec2, _ := vector.CreateVector([]int{4, 5})

	_, err := vector.ElementWiseMin(vec1, vec2)

	if err != vector.ErrMismatchedLengths {
		t.Errorf("expected ErrMismatchedLengths, got: %v", err)
	}
}

func TestElementWiseMin_NegativeNumbers(t *testing.T) {
	vec1, _ := vector.CreateVector([]int{-1, -5, -3})
	vec2, _ := vector.CreateVector([]int{-4, -2, -6})

	result, err := vector.ElementWiseMin(vec1, vec2)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expected := []int{-4, -5, -6}
	for i, val := range result.Element {
		if val != expected[i] {
			t.Errorf("expected %d at index %d, got %d", expected[i], i, val)
		}
	}
}

// Benchmark tests
func BenchmarkAddVectors(b *testing.B) {
	vec1, _ := vector.CreateVector(make([]int, 1000))
	vec2, _ := vector.CreateVector(make([]int, 1000))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vector.AddVectors(vec1, vec2)
	}
}

func BenchmarkDotProduct(b *testing.B) {
	vec1, _ := vector.CreateVector(make([]int, 1000))
	vec2, _ := vector.CreateVector(make([]int, 1000))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vector.DotProduct(vec1, vec2)
	}
}

func BenchmarkEuclideanDistance(b *testing.B) {
	vec1, _ := vector.CreateVector(make([]int, 1000))
	vec2, _ := vector.CreateVector(make([]int, 1000))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vector.EuclideanDistance(vec1, vec2)
	}
}
