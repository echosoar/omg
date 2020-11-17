package main

import (
	"github.com/echosoar/omg"
	"github.com/echosoar/omg/mw/static"
)

func main() {
	s := omg.New();
	s.Handle("/test", func (ctx *omg.Context) (string, error) {
		return  "Api:" + ctx.Req.Url, nil
	}, omg.MethodGet);

	s.Use(static.MW);

	s.Start("12355");
}