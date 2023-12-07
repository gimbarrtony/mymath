package mymath

import (
	"testing"
)

func TestAbs(t *testing.T) {
	result := Abs(-4)
	expected := 4.0
	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

func TestMax(t *testing.T) {
	result := Max(4.0, 2.0)
	expected := 4.0
	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}


func TestSqrt(t *testing.T) {
	result := Sqrt(4.0)
	expected := 2.0
	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

func TestYn(t *testing.T) {
	result := Yn(2, 2.5)
	expected := -0.38133584924180314
	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}
