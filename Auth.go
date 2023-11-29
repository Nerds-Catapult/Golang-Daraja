package darajaAuth

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
)

type AuthResponse struct {
	AccessToken string `json:"access_token"` // The access token to be used in subsequent API calls
	ExpiresIn   string `json:"expires_in"`   // The number of seconds before the access token expires
}

type Authorization struct {
	Response AuthResponse
}

func NewAuthorization(consumerKey, consumerSecret string, env Environment) (*Authorization, error) {
	authHeader := map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(consumerKey + ":" + consumerSecret)),
	}

	netPackage := newRequestPackage(nil, endpointAuth, http.MethodGet, authHeader, env)
	respBody, err := sendRequest(netPackage)
	if err != nil {
		return nil, err
	}

	var authResponse AuthResponse
	err = json.Unmarshal(respBody, &authResponse)
	if err != nil {
		return nil, err
	}

	auth := &Authorization{
		Response: authResponse,
	}

	return auth, nil
}

func sendRequest(netPackage *RequestPackage) ([]byte, error) {
	// Simulate sending request and receiving response (replace with actual code)
	// For demonstration, returning dummy response body
	dummyResponseBody := []byte(`{"access_token": "dummy_token", "expires_in": "3600"}`)
	return dummyResponseBody, nil
}
