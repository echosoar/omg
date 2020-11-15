package server;
type MiddleWareManager struct {
	list []*MW
}

type MWWrapper = func(ctx *AppContext) MW;
type MW = func(ctx *Context, next Next) (string, error);
type Next = func() (string, error);

func (mw *MiddleWareManager) Use(newMw MW) {
	mw.list = append(mw.list, &newMw)
}

func (mw *MiddleWareManager) Exec(ctx *Context, handler Handler) (string, error) {
	return mw.execIndex(ctx, 0, handler);
}

func (mw *MiddleWareManager) execIndex(ctx *Context, index int, handler Handler) (string, error) {
	mwLen := len(mw.list);
	if mwLen == 0 || index == mwLen {
		return handler(ctx);
	}
	var mwFun MW = *mw.list[index];
	data, err := mwFun(ctx, func() (string, error) {
		return mw.execIndex(ctx, index + 1, handler);
	});
	return data, err;
}