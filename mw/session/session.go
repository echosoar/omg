package session


import (
	"github.com/echosoar/omg"
	"github.com/echosoar/omg/ioc"
	"io/ioutil"
	"regexp"
	"path"
	"os"
)


type Session {
	store map[string]interface{}
	config ConfigInfo
}

func MW(app *omg.AppContext) omg.MW {
	iconfig, _ := ioc.Config(ConfigKey);
	config := iconfig.(ConfigInfo);

	store := make(map[string]interface{});

	app.SetPlugin("session", Session{&store, config});

	return func(ctx *omg.Context, next omg.Next) (string, error) {
		return next();
	}
}