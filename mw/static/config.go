package static

import (
	"github.com/echosoar/omg/ioc"
)

const ConfigKey string = "mw-static";

type ConfigInfo struct {
	Prefix string
	Dir string
	Index string
}

func init() {
	ioc.SetConfig(ConfigKey, ConfigInfo{
		"/",
		"./",
		"index.html",
	});
}