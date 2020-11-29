package static

import (
	"github.com/echosoar/omg"
	"github.com/echosoar/omg/ioc"
	"io/ioutil"
	"regexp"
	"path"
	"os"
)

func MW(app *omg.AppContext) omg.MW {
	iconfig, _ := ioc.Config(ConfigKey);
	config := iconfig.(ConfigInfo);
	prefixReg, _ := regexp.Compile("^" + config.Prefix);
	cwd, _ := os.Getwd() 

	return func(ctx *omg.Context, next omg.Next) (string, error) {
		reqFilePath := string(prefixReg.ReplaceAll([]byte(ctx.Req.Url), []byte("")));
		staticFilePath := path.Join(cwd, config.Dir, reqFilePath);
		stat, statErr := os.Stat(staticFilePath);
		if statErr != nil {
			return next();
		}

		if stat.IsDir() {
			staticFilePath = path.Join(staticFilePath, config.Index);
			stat, statErr = os.Stat(staticFilePath);
			if statErr != nil {
				return next();
			}
		}

		data, err := ioutil.ReadFile(staticFilePath);
		if err != nil {
			return "", err;
		}

		ctx.Body = data;

		return "", nil;
	}
}