package handler

import (
	"html/template"

	"github.com/gorilla/mux"
	"github.com/khusainnov/edulab/pkg/service"
)

var tml *template.Template

func init() {
	tml = template.Must(template.ParseGlob("/Users/rustamkhusainov/DocumentsAir/joba-workspace/edulab/frontend/html_templates/*.html"))
}

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", h.GreetingPage)
	r.HandleFunc("/about", h.About)

	r.HandleFunc("/auth/sign-up", h.SignUp).Methods("GET", "POST")
	r.HandleFunc("/auth/sign-in", h.SignIn)
	r.HandleFunc("/logout", h.Logout)

	r.HandleFunc("/courses", h.Courses)

	api := r.PathPrefix("/api").Subrouter()
	api.Use(h.UserIdentity)
	{
		api.HandleFunc("/profile", h.Profile)

		profile := api.PathPrefix("/profile").Subrouter()
		{
			profile.HandleFunc("/settings", nil).Methods("GET", "POST")
		}
	}

	return r
}
