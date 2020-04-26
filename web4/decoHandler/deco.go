package decoHandler

import "net/http"

type DecoratorFunc func(http.ResponseWriter, *http.Request, http.Handler)

//DecoHandler
type DecoHandler struct {
	h  http.Handler
	fn DecoratorFunc
}

func (self *DecoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	self.fn(w, r, self.h)
}

//NewDecoHandler
func NewDecoHandler(h http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoHandler{
		h:  h,
		fn: fn,
	}
}
