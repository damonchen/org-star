package web

import (
	"github.com/damonchen/org-star/internal/web/view"
	"github.com/damonchen/org-star/internal/web/view/auth"
	"github.com/damonchen/org-star/internal/web/view/fav"
	"github.com/go-chi/chi"
)

func NewRouter(svr *Server) *chi.Mux {
	r := chi.NewRouter()
	r.Use(svr.Auth)
	r.Get("/", view.Index)

	r.Route("/auth", func(r chi.Router) {
		r.Get("/authorize", auth.Authorize)
		r.Post("/login", auth.Login)
		r.Post("/logout", auth.Logout)
		r.Post("/redirect", auth.Redirect)
	})

	r.Route("/fav", func(r chi.Router) {
		r.Get("/", fav.Fav)
	})

	return r
}
