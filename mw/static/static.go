package static

import (
	"github.com/echosoar/omg"
	"github.com/echosoar/omg/ioc"
)

func MW(app *omg.AppContext) omg.MW {
	config := ioc.Config("mw-static").(Config);
	return func(ctx *omg.Context, next omg.Next) (string, error) {
		return  next();
	}
}