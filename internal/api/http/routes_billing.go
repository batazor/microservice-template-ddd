package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"google.golang.org/protobuf/encoding/protojson"

	"microservice-template-ddd/internal/billing/infrastructure/rpc"
)

// Routes creates a REST router
func (api *API) BillingRoutes() chi.Router {
	r := chi.NewRouter()

	// CRUD
	r.Post("/", api.AddBilling)
	r.Get("/", api.ListBilling)
	r.Get("/{BillingId}", api.GetBilling)
	r.Delete("/{BillingId}", api.DeleteBilling)

	return r
}

// CRUD ================================================================================================================
func (api *API) AddBilling(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (api *API) ListBilling(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	resp, err := api.BillingService.Get(r.Context(), &billing_rpc.GetRequest{Id: "test@user"})
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

func (api *API) GetBilling(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (api *API) DeleteBilling(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
