package vault

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"encoding/json"
)

//Business logic as interface
type Service interface {
	Hash(ctx context.Context, password string) (string, error)
	Validate(ctx context.Context, password string, hash string) (bool, error)
}

//implementation with empty struct (stateless)
type vaultService struct {
}

//constructor - we can later add initialization if needed
func NewService() Service {
	return vaultService{}
}

//implementation
func (vaultService) Hash(ctx context.Context, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

//implementation
func (vaultService) Validate(ctx context.Context, password string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//request and responses should be modeled, as well as decode function for each request
type hashRequest struct {
	Password string `json:"password"`
}

type hashResponse struct {
	Hash string `json:"hash"`
	Err  string `json:"err,omitempty"`
}

func decodeHashRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req hashRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

type validateRequest struct {
	Password string `json:"password"`
	Hash     string `json:"hash"`
}

type validateResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}

func decodeValidateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req validateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

//A single func for encoding the response
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
