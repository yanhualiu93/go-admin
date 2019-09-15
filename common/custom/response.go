package custom

type Response struct {
	Body Map
}

func (r *Response)SUCCESS() *Response {
	return r.SetValue("message","success")
}
func (r *Response)SetMessage(v interface{}) *Response  {
	return r.SetValue("message",v)
}

func (r *Response) SetValue(key string, val interface{}) *Response {
	if r.Body == nil {
		r.Body = make(Map)
	}
	r.Body.Set(key, val)
	return r
}
func (r *Response) GetValue(key string) (interface{}, error) {
	return r.Body.Get(key)
}
func (r *Response) GetString(key string) (string, error) {
	return r.Body.GetString(key)
}
func (r *Response) GetInt64(key string) (int64, error) {
	return r.Body.GetInt64(key)
}
func (r *Response) GetFloat64(key string) (float64, error) {
	return r.Body.GetFloat64(key)
}
