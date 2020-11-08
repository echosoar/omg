package server

import (
	"net/http"
	"sort"
)
type Server struct {
	Port string
	Routers []RouteItem
	mw MiddleWareManager
	app *AppContext
}

func (s *Server) Start(port string) {
	s.Port = port;
	// 路径进行排序
	s.sortRegister();
	http.HandleFunc("/", s.handleRequest);
	http.ListenAndServe("0.0.0.0:" + s.Port, nil);
}

func (s *Server) Get(path string, handler Handler) *Server {
	s.putToRouterItem(path, MethodGet, handler);
	return s;
}

func (s *Server) Post(path string, handler Handler) *Server {
	s.putToRouterItem(path, MethodPost, handler);
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

func (s *Server) handleRequest(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path;
	method := GetMethod(req.Method);

	newReq := Request {
		path,
		method,
		req,
	};

	newRes := Response {
		make(map[string][]string),
	};

	ctx := &Context {newReq, newRes, 200, s.app};

	handler := s.findHandlerByPathAnd(path, method);
	
	result, err := s.mw.Exec(ctx, handler);

	respHead := w.Header();
	if len(ctx.Res.Headers) > 0 {
		for key, value := range ctx.Res.Headers {
			respHead[key] = value;
		}
	}

	if err != nil {
		ctx.Status = 503;
		result = err.Error();
	}

	w.WriteHeader(ctx.Status);
	w.Write([]byte(result));
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
	}
	return server;
}