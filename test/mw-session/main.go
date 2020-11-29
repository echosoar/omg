package main

import (
	"fmt"
	"strconv"
	"github.com/echosoar/omg"
	"github.com/echosoar/omg/mw/session"
)

func main() {
	s := omg.New();
	i := 1;
	s.Handle("/", func (ctx *omg.Context) (string, error) {
		sessionPlugin, _ := ctx.Plugin("session");
		sessionInstance := sessionPlugin.(*session.Session);
		sessionValue, _ := sessionInstance.Get();
		fmt.Println(sessionValue);
		if i > 5 {
			sessionInstance.Clear()
		} else if ( i < 2) {
			sessionInstance.Set("test session value")
		}
		i++;
		return  "Api:" + ctx.Req.Url + strconv.Itoa(i), nil
	}, omg.MethodGet);

	s.Use(session.MW);

	s.Start("12355");
}