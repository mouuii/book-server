package core

import (
	"net/http"
	"reflect"
)

type ReturnHandler func(http.ResponseWriter, *http.Request, []reflect.Value)

func defaultReturnHandle() ReturnHandler  {
	return func(writer http.ResponseWriter, request *http.Request, values []reflect.Value) {

		var responseVal reflect.Value
		if len(values) > 1 && values[0].Kind() == reflect.Int {
			writer.WriteHeader(int(values[0].Int()))
			responseVal = values[1]
		}else if len(values) > 0 {
			responseVal = values[0]
		}

		if isByteSlice(responseVal) {
			writer.Write(responseVal.Bytes())
		}else {
			writer.Write([]byte(responseVal.String()))
		}
	}
}

func isByteSlice(val reflect.Value) bool {
	return val.Kind() == reflect.Slice && val.Type().Elem().Kind() == reflect.Uint8
}