package omg
func Default404(ctx *Context) (string, error) {
	ctx.Status = 404;
	return "Page not found <br /> Powered by Gos", nil;
}