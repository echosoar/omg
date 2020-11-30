package static

import (
	"testing"
	"net/http"
	"io/ioutil"
	"github.com/echosoar/omg"
	"github.com/echosoar/omg/ioc"
)

func TestMWStatic(t *testing.T) {
	s := omg.New();
	ioc.SetConfig(ConfigKey, ConfigInfo{
		"/",
		"./testdata",
		"index.html",
	});
	s.Use(MW);
	go s.Start("12355");
	clt := http.Client{};
	resp, _ := clt.Get("http://127.0.0.1:12355/");
	content, _ := ioutil.ReadAll(resp.Body);
	defer func() {
		clt.CloseIdleConnections();
		s.Close();
	}();

	if string(content) != "xxxx" {
		t.Errorf("mw/static")
	}
}