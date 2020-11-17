// Package ioc
package ioc

func SetConfig(key string, value interface{}) {
	Get("Config").ProvideByName(key, value);
}

func Config(key string) (interface{}, error) {
	return Get("Config").Inject(key);
}