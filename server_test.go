package omg

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	s := New();
	s.Handle("/test", func (ctx *Context) (string, error) {
		return  "Api:" + ctx.Req.Url, nil
	}, MethodGet);

	go s.Start("12355");
	clt := http.Client{};
	resp, _ := clt.Get("http://127.0.0.1:12355/test");
	content, _ := ioutil.ReadAll(resp.Body);

	defer func() {
		clt.CloseIdleConnections();
		s.Close();
	}();

	if string(content) != "Api:/test" {
		t.Errorf("request /test")
	}
}