package omg

import (
	"strings"
	"regexp"
)
type RouteItem struct {
	Path string
	Level int
	PurePath string
	MatchReg *regexp.Regexp
	Register []RouteRegister // 多个method
}

type Handler = func (ctx *Context) (string, error);

type RouteRegister struct {
	Method Method
	Handler Handler
}

func GetRouteItem(path string) RouteItem {
	pathList := strings.Split(path, "/");
	pathLevel := len(pathList) - 1;

	// 移除 * 的纯路径
	pureReg, _ := regexp.Compile(`\*+$`);
	purePath := pureReg.ReplaceAllString(path, "");

	// 用来匹配路径是否与当前注册器匹配
	pathMatchReplaceReg, _ := regexp.Compile(`\.+`);
	pathMatchStr := pathMatchReplaceReg.ReplaceAllString(path, `\\.`);
	pathMatchReplaceAllReg, _ := regexp.Compile(`\*+`);
	pathMatchStr = pathMatchReplaceAllReg.ReplaceAllString(pathMatchStr, `.*`);
	pathMatchRes, _ := regexp.Compile(pathMatchStr);

	newRouter := RouteItem{
		path,
		pathLevel,
		purePath,
		pathMatchRes,
		nil,
	}

	return newRouter;
}