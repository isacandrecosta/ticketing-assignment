package views

type Response struct {
	StatusCode int
	Body       interface{}
}

func (r Response) GetStatusCode() int {
	return r.StatusCode
}

func (r Response) GetBody() interface{} {
	return r.Body
}
