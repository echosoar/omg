package omg
func Default404(ctx *Context) (string, error) {
	ctx.Status = 404;
	ctx.Res.Type = "html";
	return "Page not found <br /> Powered by Omg", nil;
}