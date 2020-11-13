package http

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Routes creates a REST router
func (api *API) UserRoutes() chi.Router {
	r := chi.NewRouter()

	// CRUD
	r.Post("/", api.AddUser)
	r.Get("/", api.ListUser)
	r.Get("/{userId}", api.GetUser)
	r.Delete("/{userId}", api.DeleteUser)

	return r
}

// CRUD ================================================================================================================
func (api *API) AddUser(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (api *API) ListUser(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (api *API) GetUser(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (api *API) DeleteUser(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
