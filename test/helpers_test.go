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

func expect(t *testing.T, a interface{}, b interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	fn = strings.Replace(fn, wd+"/", "", -1)

	if !reflect.DeepEqual(a, b) {
		t.Errorf("(%s:%d) expected %v (type %v) - got %v (type %v)", fn, line, b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
