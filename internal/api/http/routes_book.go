package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/encoding/protojson"

	"microservice-template-ddd/internal/book/infrastructure/rpc"
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
	w.Header().Add("Content-type", "application/json")

	resp, err := api.BookService.Get(r.Context(), &book_rpc.GetRequest{Id: "Hello World"})
	if err != nil {
		api.Log.Error(err.Error())
		_, _ = w.Write([]byte(`{"error": "error 0_o"}`))
		return
	}

	m := protojson.MarshalOptions{}
	payload, err := m.Marshal(resp)
	if err != nil {
		api.Log.Error(err.Error())
		_, _ = w.Write([]byte(`{"error": "error 0_o"}`))
	}

	_, _ = w.Write(payload)
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

	resp, err := api.BookService.Rent(r.Context(), &book_rpc.RentRequest{Id: "Hello World"})
	if err != nil {
		api.Log.Error(err.Error())
		_, _ = w.Write([]byte(`{"error": "error 0_o"}`))
		return
	}

	m := protojson.MarshalOptions{}
	payload, err := m.Marshal(resp)
	if err != nil {
		api.Log.Error(err.Error())
		_, _ = w.Write([]byte(`{"error": "error 0_o"}`))
	}

	_, _ = w.Write(payload)
}
