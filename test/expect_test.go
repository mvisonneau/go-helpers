package test

import (
	"testing"
)

func TestExpectSucceed(t *testing.T) {
	a, b := 1, 1
	newTest := &testing.T{}
	Expect(newTest, a, b)
	if newTest.Failed() {
		t.Fatalf("should have succeeded")
	}
}

func TestExpectFailed(t *testing.T) {
	a, b := 1, 2
	newTest := &testing.T{}
	Expect(newTest, a, b)
	if !newTest.Failed() {
		t.Fatalf("should have failed")
	}
}
