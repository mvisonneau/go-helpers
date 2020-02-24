package test

import (
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

var (
	wd, _ = os.Getwd()
)

// Equal compares 2 values and throw an error if they differ
func Equal(t *testing.T, a interface{}, b interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	fn = strings.Replace(fn, wd+"/", "", -1)

	if !reflect.DeepEqual(a, b) {
		t.Errorf("(%s:%d) expected %v (type %v) - got %v (type %v)", fn, line, b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

// TypeEqual compares 2 values types and throw an error if they differ
func TypeEqual(t *testing.T, a interface{}, b interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	fn = strings.Replace(fn, wd+"/", "", -1)

	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		t.Errorf("(%s:%d) expected type %v - got %v", fn, line, reflect.TypeOf(b), reflect.TypeOf(a))
	}
}
