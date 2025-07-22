package vector_test

import (
	"testing"

	"github.com/wendersoon/gomathx/vector"
)

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
