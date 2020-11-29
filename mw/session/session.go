package session

import (
	"errors"
	"github.com/echosoar/omg"
	"github.com/echosoar/omg/ioc"
	"github.com/google/uuid"
)

/*
仅实现最基本的存储、获取和过期功能
存储：通过ctx的cookie中的 session key 字段来获取id（若没有则随机生成），记录存储的内容和插入时间
获取：通过ctx的cookie中的 session key 字段来获取id，从全局store中获取存储的内容，并校验过期时间
*/

type globalSession struct {
	store map[string]interface{}
	config ConfigInfo
}

type Session struct {
	ctx *omg.Context
	sess *globalSession
}


func (sess *globalSession) GetContextInstance(ctx *omg.Context) interface{} {
	return &Session{
		ctx: ctx,
		sess: sess,
	};
}

func (sess *Session) Get() (interface{}, error) {
	CookieKey := sess.sess.config.Key;
	sessionId := sess.ctx.GetCookie(CookieKey);
	if sessionId == "" {
		return nil, errors.New("session not exists");
	}
	sessionValue := sess.sess.store[sessionId];
	if sessionValue == nil {
		return nil, errors.New("session not exists");
	}
	return sessionValue, nil;
}

func (sess *Session) Set(value interface{}) () {
	CookieKey := sess.sess.config.Key;
	sessionId := sess.ctx.GetCookie(CookieKey);
	if sessionId == "" {
		sessionId = uuid.New().String();
	}
	sess.sess.store[sessionId] = value;
	sessionOptions := make(map[string]interface{});
	sess.ctx.SetCookie(CookieKey, sessionId, sessionOptions);
}

func (sess *Session) Clear() () {
	CookieKey := sess.sess.config.Key;
	sessionId := sess.ctx.GetCookie(CookieKey);
	if sessionId == "" {
		return;
	}
	delete(sess.sess.store, sessionId);
}

func MW(app *omg.AppContext) omg.MW {
	iconfig, _ := ioc.Config(ConfigKey);
	config := iconfig.(ConfigInfo);

	globalSessionInstance := &globalSession{
		store: make(map[string]interface{}),
		config: config,
	}

	return func(ctx *omg.Context, next omg.Next) (string, error) {
		ctx.SetPlugin("session", globalSessionInstance.GetContextInstance(ctx));
		return next();
	}
}