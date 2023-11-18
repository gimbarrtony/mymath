package mymath

import (
	"testing"
)

func TestCeil(t *testing.T) {
	result := Ceil(4.0)
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

func TestFloor(t *testing.T) {
	result := Floor(4.0)
	expected := 4.0
	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

func TestPow(t *testing.T) {
	result := Pow(4.0, 2.0)
	expected := 16.0
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

func TestMin(t *testing.T) {
	result := Min(4.0, 2.0)
	expected := 2.0
	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}
