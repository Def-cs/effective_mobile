package api

import (
	logger "effective_mobile/logs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type Handle struct {
	method map[string]func(r *http.Request) ([]byte, int, error)
	path   string
	logger logger.LogInterface
}

func newHandle(method string, path string, handle func(r *http.Request) ([]byte, int, error), logger logger.LogInterface) {
	if _, ok := urls[path]; ok {
		urls[path].method[method] = handle
	} else {
		methods := make(map[string]func(r *http.Request) ([]byte, int, error))
		methods[method] = handle
		urls[path] = &Handle{
			method: methods,
			path:   path,
			logger: logger,
		}
	}
}

func (h *Handle) Handle(w http.ResponseWriter, r *http.Request) {
	if handler, ok := h.method[r.Method]; ok {
		json, status, err := handler(r)
		if err != nil {
			h.logger.Error(err.Error() + r.Method + " " + r.URL.Path)
		}
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)

		h.logger.Info("Done request: " + r.Method + " " + r.URL.Path)
	} else {
		h.logger.Error("Method not allowed" + r.Method + " " + r.URL.Path)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func RegisterMux(logger logger.LogInterface) *mux.Router {
	createHandlers(logger)
	router := mux.NewRouter()
	for path, handle := range urls {
		router.HandleFunc(path, handle.Handle)
	}

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return router
}
