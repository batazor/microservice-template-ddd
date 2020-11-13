package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/golang/protobuf/jsonpb"

	billing_rpc "robovoice-template/internal/billing/domain"
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
		w.Write([]byte(`{"error": "error 0_o"}`))
		return
	}

	m := jsonpb.Marshaler{}
	err = m.Marshal(w, resp)
	if err != nil {
		api.Log.Error(err.Error())
		w.Write([]byte(`{"error": "error 0_o"}`))
	}
}

func (api *API) GetBilling(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (api *API) DeleteBilling(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
