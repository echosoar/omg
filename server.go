package omg

import (
	"github.com/valyala/fasthttp"
	"sort"
	"bytes"
)
type Server struct {
	Port string
	Routers []RouteItem
	mw MiddleWareManager
	app *AppContext
	srv *fasthttp.Server
}

func (s *Server) Start(port string, protocol ...string) {
	s.Port = port;
	// 路径进行排序
	s.sortRegister();

	s.srv = &fasthttp.Server{
    Handler: s.handlerRequest,
    Name: "Omg",
	}

	protocolLen := len(protocol);
	if protocolLen == 0 {
		protocol = append(protocol, ProtocolHttp);
	}
	if UtilsExistsInSlice(protocol, "http") {
		s.srv.ListenAndServe(":" + s.Port);
	}

	// https
}

func (s *Server) Close() error {
	return s.srv.Shutdown();
}

func (s *Server) Get(path string, handler Handler) *Server {
	s.putToRouterItem(path, MethodGet, handler);
	return s;
}

func (s *Server) Post(path string, handler Handler) *Server {
	s.putToRouterItem(path, MethodPost, handler);
	return s;
}

func (s *Server) Handle(path string, handler Handler, methods ...Method) *Server {
	for _, method := range methods {
		s.putToRouterItem(path, method, handler);
	}
	return s;
}

func (s *Server) Use(mw MWWrapper) *Server {
	s.mw.Use(mw(s.app));
	return s;
}

func (s *Server) findRouter(path string) *RouteItem {
	for _, routerItem := range s.Routers {
		if routerItem.Path == path {
			return &routerItem;
		}
	}
	newRouter := GetRouteItem(path);
	s.Routers = append(s.Routers, newRouter);
	return &s.Routers[len(s.Routers) - 1]; 
}

func (s *Server) handlerRequest(fsh *fasthttp.RequestCtx) {
	path := string(fsh.Path());
	method := GetMethod(string(fsh.Method()));

	newReq := Request {
		path,
		method,
		&fsh.Request,
	};

	newRes := Response {
		make(map[string][]string),
		"",
		&fsh.Response,
	};

	ctx := &Context {newReq, newRes, 200, nil, s.app};

	handler := s.findHandlerByPathAnd(path, method);
	
	result, err := s.mw.Exec(ctx, handler);


	contentType := UtilsGetContentType(ctx.Res.Type, ctx.Req.Url);

	respHead := &fsh.Response.Header;
	respHead.Set("Content-Type", contentType);
	if len(ctx.Res.Headers) > 0 {
		for key, value := range ctx.Res.Headers {
			for index, valueLine := range value {
				if index == 0 {
					respHead.Set(key, valueLine);
				} else {
					respHead.Add(key, valueLine);
				}
			}
		}
	}

	if err != nil {
		ctx.Status = 503;
		result = err.Error();
		ctx.Body = nil;
	}

	fsh.Response.SetStatusCode(ctx.Status);

	if ctx.Body != nil {
		fsh.Response.SetBodyStream(bytes.NewReader(ctx.Body), len(ctx.Body));
	} else {
		fsh.Response.SetBodyString(result);
	}
}

func (s *Server) putToRouterItem(path string, method Method, handler Handler) {
	currentRouters := s.findRouter(path);
	currentRouters.Register = append(currentRouters.Register, RouteRegister{ method, handler });
}

func (s *Server) findHandlerByPathAnd(path string, method Method) Handler {
	var matchedRouter RouteItem;
	for _, router := range s.Routers {
		match := router.MatchReg.MatchString(path);
		if match {
			matchedRouter = router;
			break;
		}
	}
	if len(matchedRouter.Register) > 0 {
		for _, resgister := range matchedRouter.Register {
			if resgister.Method == method {
				return resgister.Handler;
			}
		}
	}
	// Todo return 404
	return Default404;
}

// 对路径注册器进行排序
func (s *Server) sortRegister() {
	sort.Slice(s.Routers, func(i, j int) bool {
		pre := s.Routers[i];
		next := s.Routers[j];
		if pre.Level == next.Level {
			if next.PurePath == pre.PurePath {
				return len(pre.Path) - len(next.Path) < 0;
			}
			return len(next.PurePath) - len(pre.PurePath) < 0;
		}
		return next.Level - pre.Level < 0;
	});
}

func New() *Server {
	server := &Server {
		"",
		nil,
		MiddleWareManager{},
		&AppContext{},
		nil,
	}
	return server;
}