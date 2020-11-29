package omg

import (
	"strings"
)
type Method string;

const (
	MethodGet Method = "get"
	MethodPost Method = "post"
	MethodAll Method = "all"
)

func GetMethod(method string) Method {
	lowerMethod := strings.ToLower(method);
	newMethod := MethodAll;
	switch lowerMethod {
		case "get":
			newMethod = MethodGet;
		case "post":
			newMethod = MethodPost;
	}

	return newMethod;
}