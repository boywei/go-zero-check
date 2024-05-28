package cmd

import (
	"fmt"
	"github.com/traefik/yaegi/interp"
	"reflect"
	"testing"
)

func TestRunCode(t *testing.T) {
	result := f()
	fmt.Println(result)
	fmt.Println(*result)
	fmt.Println(&result)
	fmt.Println(fmt.Sprintf("%v", result))
	fmt.Println(fmt.Sprintf("%v", *result))
}

func f() *reflect.Value {
	i := interp.New(interp.Options{})
	result, _ := i.Eval("3")
	return &result
}
