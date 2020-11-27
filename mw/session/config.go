package session

import (
	"github.com/echosoar/omg/ioc"
)

const ConfigKey string = "mw-session";

type ConfigInfo struct {
	Key string
	MaxAge int
}

func init() {
	ioc.SetConfig(ConfigKey, ConfigInfo{
		"omg:session",
		86400000, // ms(24h)
	});
}