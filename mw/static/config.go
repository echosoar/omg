package static

import (
	"github.com/echosoar/omg/ioc"
)

type Config struct {
	prefix string
}

ioc.SetConfig("mw-static", MWStaticConfig{"/"});