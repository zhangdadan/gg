package unit

import (
	"dakesolo.com/gg/share/tool/unique"
	"net/http"
)


type Context struct {
	Action Action
	Response http.ResponseWriter
	Request *http.Request
	Ai string
}
type Action func(*Context) string

func (b *Context) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//动态拦截器，比如，可以对Response进行设置
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	//设置单元全局变量
	b.Request = r
	b.Response = w

	//唯一识别串
	b.Ai = unique.GetMD5()
	w.Write([]byte(b.Action(b)))
}




