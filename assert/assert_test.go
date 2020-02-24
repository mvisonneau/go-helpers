package test

import (
	"testing"
)

func TestEqualSucceed(t *testing.T) {
	a, b := 1, 1
	newTest := &testing.T{}
	Equal(newTest, a, b)
	if newTest.Failed() {
		t.Fatalf("should have succeeded")
	}
}

func TestEqualFailed(t *testing.T) {
	a, b := 1, 2
	newTest := &testing.T{}
	Equal(newTest, a, b)
	if !newTest.Failed() {
		t.Fatalf("should have failed")
	}
}

func TestTypeEqualSucceed(t *testing.T) {
	a, b := 1, 1
	newTest := &testing.T{}
	TypeEqual(newTest, a, b)
	if newTest.Failed() {
		t.Fatalf("should have succeeded")
	}
}

func TestTypeEqualFailed(t *testing.T) {
	a, b := 1, "1"
	newTest := &testing.T{}
	TypeEqual(newTest, a, b)
	if !newTest.Failed() {
		t.Fatalf("should have failed")
	}
}
