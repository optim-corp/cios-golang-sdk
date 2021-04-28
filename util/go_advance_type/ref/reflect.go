package ref

import (
	"fmt"
	"reflect"
)

type (
	FunctionName string
)

func ExecFunc(function interface{}, inputs ...interface{}) []reflect.Value {
	result, _ := ExecFuncWithName(function, inputs...)
	return result
}
func ExecFuncWithName(function interface{}, inputs ...interface{}) ([]reflect.Value, FunctionName) {
	if fv := Indirect(function); fv.Kind() == reflect.Func {
		var args []reflect.Value
		for _, input := range inputs {
			args = append(args, Indirect(input))
		}
		return fv.Call(args), FunctionName(fv.Type().Name())
	} else {
		panic("Not a function: " + fmt.Sprintf("%v", function))
	}
}

func Indirect(target interface{}) reflect.Value {
	return reflect.Indirect(reflect.ValueOf(target))
}
