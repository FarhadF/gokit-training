package vault

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"encoding/json"
)
//func (e Endpoints) register (r *httprouter.Router) {
//	r.POST("/hash", e.handleHash)
//	r.POST("/validate", e.handleValidate)
//}

/*func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
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

}*/

//using http router, register func will do the routing path registration
func (e Endpoints) Register (r *httprouter.Router) {
	r.Handle("POST", "/hash", e.HandleHashPost)
	r.Handle("POST", "/validate", e.HandleValidatePost)
}

//each method needs a http handler handlers are registered in the register func
func (e Endpoints) HandleHashPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decodedHashReq, err := decodeHashRequest(e.Ctx, r)
	if err != nil {
		respondError(w, 500, err)
		return
	}
	resp, err := e.HashEndpoint(e.Ctx, decodedHashReq.(hashRequest))
	if err != nil {
		respondError(w, 500, err)
		return
	}
	respondSuccess(w, resp.(hashResponse))
}

//each method needs a http handler
func(e Endpoints) HandleValidatePost(w http.ResponseWriter,r *http.Request, _ httprouter.Params) {
	decodeValidateReq, err := decodeValidateRequest(e.Ctx, r)
	if err != nil {
		respondError(w, 500, err)
		return
	}
	resp,err := e.ValidateEndpoint(e.Ctx, decodeValidateReq.(validateRequest))
	if err != nil {
		respondError(w, 500, err)
		return
	}
	respondSuccess(w, resp.(validateResponse))
}

// respondError in some canonical format.
func respondError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":       err,
		"status_code": code,
		"status_text": http.StatusText(code),
	})
}

// respondSuccess in some canonical format.
func respondSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}