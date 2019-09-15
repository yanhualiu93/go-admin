package custom

import (
	"fmt"
	"go-admin/common/errors"
	"go-admin/common/page"
	"net/http"
)

type Request struct {
	Params Map
}

func (r *Request) SetParam(key string, val interface{}) *Request {
	if r.Params == nil {
		r.Params = make(Map)
	}
	r.Params.Set(key, val)
	return r
}
func (r *Request) GetParam(key string) (interface{}, error) {
	return r.Params.Get(key)
}
func (r *Request) GetInt64(key string) (int64, error) {
	return r.Params.GetInt64(key)
}
func (r *Request) GetString(key string) (string, error) {
	return r.Params.GetString(key)
}
func (r *Request)GetPage(keys ...string) *page.Info {
	var PageI interface{}
	if len(keys) ==1{
		PageI,_ = r.Params.Get(keys[0])
	}
	switch info := PageI.(type) {
	case page.Info:
		return &info
	default:
		return nil
	}
}
func (r *Request)GetMap(key string)(map[string]interface{}, error)  {
	 mI,err:= r.Params.Get(key)
	if err==nil {
		return  nil ,err
	}
	switch m:= mI.(type) {
	case map[string]interface{}:
		return m,nil
	default:
		return nil,  errors.New(http.StatusBadRequest,fmt.Sprintf("request params %s,can not case map[string]interface",key))
	}
}