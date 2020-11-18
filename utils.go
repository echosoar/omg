package omg
import (
	"mime"
	"reflect"
	"path/filepath"
)

func UtilsGetContentType(contentType string, reqUrl string) string {
	ctype := contentType;
	if ctype == "" {
		extension := filepath.Ext(reqUrl);
		if extension == "" {
			ctype = "html";
		} else {
			ctype = extension;
		}
	}
	if ctype[0] != '.' {
		ctype = "." + ctype;
	}
	return mime.TypeByExtension(ctype);
}

func UtilsExistsInSlice(s []string, findValue string) bool {
	for _, value := range s {
		if value == findValue {
			return true;
		}
	}
	return false;
}


func UtilsToBool(value interface{}) bool {
	t := reflect.TypeOf(value);
	kind := t.Kind();
	v := reflect.ValueOf(value);
	switch (kind) {
	case reflect.Bool:
		return v.Bool();
	case reflect.String:
		s := v.String();
		if s == "true" || s == "TRUE" || s == "True" {
			return true;
		}
		return false;
	}
	return false;
}

func UtilsToString(value interface{}) string {
	t := reflect.TypeOf(value);
	kind := t.Kind();
	v := reflect.ValueOf(value);
	switch (kind) {
	case reflect.String:
		return v.String();
	}
	return "";
}

func UtilsToInt(value interface{}) int {
	t := reflect.TypeOf(value);
	kind := t.Kind();
	v := reflect.ValueOf(value);

	switch (kind) {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int(v.Int());
	}
	return 0;
}