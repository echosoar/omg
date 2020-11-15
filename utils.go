package omg
import (
	"mime"
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