package session

import (
	"testing"
	"net/http"
	"math/rand"
	"io/ioutil"
	"strconv"
	"github.com/echosoar/omg"
)

func TestMWStatic(t *testing.T) {
	rankStr := strconv.Itoa(rand.Intn(100));
	s := omg.New();
	i := 1;
	s.Handle("/", func(ctx *omg.Context) (string, error) {
		sessionPlugin, _ := ctx.Plugin("session");
		sessionInstance := sessionPlugin.(*Session);
		sessionValue, _ := sessionInstance.Get();
		if i == 1 {
			sessionInstance.Set(rankStr);
		}
		i++;
		if sessionValue == nil {
			return "", nil;
		}
		return sessionValue.(string), nil;
	}, omg.MethodGet);
	s.Use(MW);
	go s.Start("12355");
	clt := http.Client{};
	url := "http://127.0.0.1:12355/";
	resp, _ := clt.Get(url);
	content, _ := ioutil.ReadAll(resp.Body);
	defer func() {
		clt.CloseIdleConnections();
		s.Close();
	}();

	if string(content) != "" {
		t.Errorf("mw/session first request value is not empty")
	}

	reqest, _ := http.NewRequest("GET", url, nil)
	originCookie := resp.Header["Set-Cookie"];
	reqest.Header.Add("Cookie", originCookie[0]);
	resp2, _ := clt.Do(reqest);
	content2, _ := ioutil.ReadAll(resp2.Body);
	if string(content2) != rankStr {
		t.Errorf("mw/session second request value is not match")
	}
}