package omg

import (
	"github.com/valyala/fasthttp"
	"time"
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
	Body []byte
	app *AppContext
}

type Request struct {
	Url string
	Method Method
	OriginReq *fasthttp.Request
}

type Response struct {
	Headers map[string][]string
	Type string
	OriginRes *fasthttp.Response
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

func (ctx *Context) GetCookie(cookieName string) string {
	cookieValueBytes := ctx.Req.OriginReq.Header.Cookie(cookieName);
	return string(cookieValueBytes);
}

func (ctx *Context) SetCookie(cookieName string, cookieValue string, options map[string]interface{}) {
	newCookie := &fasthttp.Cookie{};
	newCookie.SetKey(cookieName);
	newCookie.SetValue(cookieValue);
	if options["domain"] != nil {
		newCookie.SetDomain(UtilsToString(options["domain"]));
	}
	if options["expire"] != nil {
		newCookie.SetExpire(options["expire"].(time.Time));
	}
	if options["http-only"] != nil {
		options["httpOnly"] = options["http-only"];
	}
	if options["httpOnly"] != nil {
		newCookie.SetHTTPOnly(UtilsToBool(options["httpOnly"]));
	}
	
	if options["max-age"] != nil {
		options["maxAge"] = options["max-age"];
	}
	if options["maxAge"] != nil {
		newCookie.SetMaxAge(UtilsToInt(options["maxAge"]));
	}

	if options["path"] != nil {
		newCookie.SetPath(UtilsToString(options["path"]));
	}

	if options["secure"] != nil {
		newCookie.SetSecure(UtilsToBool(options["secure"]));
	}

	if options["same-site"] != nil {
		options["sameSite"] = options["same-site"];
	}

	if options["sameSite"] != nil {
		newCookie.SetSameSite(options["sameSite"].(fasthttp.CookieSameSite));
	}
 	ctx.Res.OriginRes.Header.SetCookie(newCookie);
}
