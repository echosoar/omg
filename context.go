package omg

import (
	"github.com/valyala/fasthttp"
)

type AppContext struct {
	pluginMap map[string]interface{}
}
func (app *AppContext) SetPlugin(pluginName string, pluginClient interface {}) {
	if app.pluginMap == nil {
		app.pluginMap = make(map[string]interface{});
	}
	app.pluginMap[pluginName] = pluginClient;
}

func (app *AppContext) GetPlugin(pluginName string) interface{} {
	return app.pluginMap[pluginName];
}

type Context struct {
	Req Request
	Res Response
	Status int // response status
	app *AppContext
	Body []byte
}

type Request struct {
	Url string
	Method Method
	OriginReq *fasthttp.Request
}

type Response struct {
	Headers map[string][]string
	Type string
}

func (ctx *Context) Get(key string) []string {
	headerValueBytes := ctx.Req.OriginReq.Header.Peek(key);
	var values []string;
	return append(values, string(headerValueBytes));
}

func (ctx *Context) Set(key string, value ...string) {
	ctx.Res.Headers[key] = value;
}

func (ctx *Context) Redirect(target string) {
	ctx.Status = 302;
	ctx.Set("Location", target);
}

func (ctx *Context) Plugin(pluginName string) interface{} {
	return ctx.app.GetPlugin(pluginName);
}