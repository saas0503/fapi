package interceptor

import "net/http"

const HeaderFlagDoNotIntercept = "do_not_intercept"

type ExcludeHeaderWrite struct {
	http.ResponseWriter
	ExcludeHeaders []string
}

func (w *ExcludeHeaderWrite) WriteHeader(code int) {
	for _, header := range w.ExcludeHeaders {
		w.Header().Del(header)
	}
	w.ResponseWriter.WriteHeader(code)
}

type RoutingStatusWriter struct {
	http.ResponseWriter

	NotFoundHandle       func() bool
	MethodNotAllowHandle func() bool

	statusCode  int
	intercepted bool
}

func (w *RoutingStatusWriter) WriteHeader(statusCode int) {
	if w.intercepted {
		return
	}
	w.statusCode = statusCode
	if (w.NotFoundHandle() && statusCode == http.StatusNotFound) ||
		(w.MethodNotAllowHandle() && statusCode == http.StatusMethodNotAllowed) {
		w.intercepted = true
		return
	}

	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *RoutingStatusWriter) Write(b []byte) (int, error) {
	if w.intercepted {
		return 0, nil
	}
	return w.ResponseWriter.Write(b)
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	interceptor := &RoutingStatusWriter{
		ResponseWriter: &ExcludeHeaderWrite{
			ResponseWriter: w,
			ExcludeHeaders: []string{HeaderFlagDoNotIntercept},
		},
		NotFoundHandle: func() bool {
			return w.Header().Get(HeaderFlagDoNotIntercept) == ""
		},
		MethodNotAllowHandle: func() bool {
			return w.Header().Get(HeaderFlagDoNotIntercept) == ""
		},
	}
}
