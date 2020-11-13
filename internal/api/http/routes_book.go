package http

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Routes creates a REST router
func (api *API) BookRoutes() chi.Router {
	r := chi.NewRouter()

	// CRUD
	r.Post("/", api.AddBook)
	r.Get("/", api.ListBook)
	r.Get("/{bookId}", api.GetBook)
	r.Delete("/{bookId}", api.DeleteBook)

	// RENT
	r.Post("/rent/{bookId}", api.RentBook)

	return r
}

// CRUD ================================================================================================================
func (api *API) AddBook(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (api *API) ListBook(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (api *API) GetBook(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (api *API) DeleteBook(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

// RENT ================================================================================================================
func (api *API) RentBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	panic("implement me")
}
