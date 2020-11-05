package ioc

import (
	"reflect"
	"runtime"
	"strings"
)

func GetAnyThingNameAndType(anyThing interface{}) (string, string) {
	t := reflect.TypeOf(anyThing)
	kind := t.Kind();
	var typeStr = "unknow";

	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		typeStr = "int";
	case reflect.Float32, reflect.Float64:
		typeStr = "float";
	default:
		typeStr = strings.ToLower(kind.String());
	}

	if typeStr == "func" {
		return runtime.FuncForPC(reflect.ValueOf(anyThing).Pointer()).Name(), typeStr;
	}

	return t.Name(), typeStr;
}