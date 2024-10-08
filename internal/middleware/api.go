package handler

import(
	"github.com/go-chi/chi.git"
	chimiddle "github.com/go-chi/middleware"
	"github.com/interlucas/Proj6_GO_calculator_api/internal/middleware"
)

func Handler(r *chi.Mux){
	//Global middleware
	r.User(chimiddle.StripSlashes)
}