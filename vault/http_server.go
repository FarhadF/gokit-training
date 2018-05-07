package vault

import (
	"net/http"
	"context"
	//"github.com/julienschmidt/httprouter"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)
//func (e Endpoints) register (r *httprouter.Router) {
//	r.POST("/hash", e.handleHash)
//	r.POST("/validate", e.handleValidate)
//}

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/hash", httptransport.NewServer(
		endpoints.HashEndpoint,
		decodeHashRequest,
		encodeResponse,
	))
	m.Handle("/validate", httptransport.NewServer(
		endpoints.ValidateEndpoint,
		decodeValidateRequest,
		encodeResponse,
	))
	m.Handle("/metrics", promhttp.Handler())
	return m

}
