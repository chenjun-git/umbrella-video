package handler

import (
	"github.com/go-chi/chi"

	"business/video/handler/v1.0"
)

func RegisterVideoRouter() *chi.Mux {
	router := chi.NewRouter()
	v1_0.RegisterRouter(router)
	return router
}
