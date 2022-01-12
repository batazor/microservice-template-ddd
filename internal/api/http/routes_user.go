package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"google.golang.org/protobuf/encoding/protojson"

	"microservice-template-ddd/internal/user/infrastructure/rpc"
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
	w.Header().Add("Content-type", "application/json")

	resp, err := api.UserService.Get(r.Context(), &user_rpc.GetRequest{Id: "test@user"})
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

func (api *API) GetUser(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (api *API) DeleteUser(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
