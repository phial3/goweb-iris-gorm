package model

type Result struct {
	Code int
	Msg  string
	Data interface{}
}

func NewResult(c int, data interface{}, m ...string) *Result {
	r := &Result{Data: data, Code: c}

	if e, ok := data.(error); ok {
		if m == nil {
			r.Msg = e.Error()
		}
	} else {
		r.Msg = "success"
	}
	if len(m) > 0 {
		r.Msg = m[0]
	}

	return r
}
