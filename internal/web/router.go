package web

import (
	"github.com/damonchen/org-star/internal/web/view"
	"github.com/go-chi/chi"
)

func NewRouter(svr *Server) *chi.Mux {
	r := chi.NewRouter()
	r.Use(svr.Auth)
	r.Get("/", view.Index)

	r.Get("/oss/{provider}", svr.Handle)
	return r
}
