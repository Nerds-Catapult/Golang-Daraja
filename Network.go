package darajaAuth

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type networkPackage struct {
	Payload interface{}
	Endpoint string
	Method string
	Headers map[string]string
}

type networkResponse[T any] struct {
	Body T
	StatusCode int
}

func newRequestPackage(payload interface{}, endpoint string, method string, headers map[string]string, env Environment) *networkPackage {
	var reqUrl = baseUrlSandbox
	if env == ENVIROMENT_PRODUCTION {
		reqUrl = baseUrlLive
	}
	reqUrl = reqUrl + endpoint

	if method == http.MethodGet {
		q := url.Values{}
		var mapPayload map[string]interface{} = struct2Map(payload)
		if len(mapPayload) > 0 {
			for key, value := range mapPayload {
				q.Add(key, value.(string))
			}
			if strings.Index(reqUrl, "?") == -1 {
				reqUrl = reqUrl + "?" + q.Encode()
			} else {
				reqUrl = reqUrl + "&" + q.Encode()
			}
			reqUrl += q.Encode()
		}
	}
	return &networkPackage{
		Payload: payload,
		Endpoint: reqUrl,
		Method: method,
		Headers: headers,
	}
}

func (p *networkPackage) addHeader(key string, value string) {
	if p.Headers == nil {
		p.Headers = make(map[string]string)
	}
	p.Headers[key] = value
}

func newRequest[T any](pac *networkPackage) (*networkResponse[T], *ErrorResponse){
	netResponseHolder := &networkResponse[T]{}
	client := &http.Client{}
	var jsonDataBytes []byte
	var httpReq *http.Request

	if pac.Payload != nil {
		jsonDataBytes, _ = json.Marshal(pac.Payload)
		httpReq, _ = http.NewRequest(pac.Method, pac.Endpoint, strings.NewReader(string(jsonDataBytes)))
	} else {
		httpReq, _ = http.NewRequest(pac.Method, pac.Endpoint, nil)
	}
	for key, value := range pac.Headers {
		httpReq.Header.Add(key, value)
	}
	if pac.Method == http.MethodPost {
		httpReq.Header.Add("Content-Type", "application/json")
	}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, &ErrorResponse{error: err}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	netResponseHolder.StatusCode = resp.StatusCode
	if netResponseHolder.StatusCode >= 400 {
		if resp.Body != nil{
			var errorResponse ErrorResponse
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, &ErrorResponse{error: err}
			}
			//errorResponse.Raw = string(body)
			bodyString := string(body)

			err = json.Unmarshal([]byte(bodyString), &errorResponse)
			if err != nil {
				if bodyString != "" {
					return nil, &ErrorResponse{error: errors.New(resp.Status)}
				}
				return nil, &ErrorResponse{error: errors.New(resp.Status + " " + bodyString)}
			}
			if errorResponse.ErrorMessage != "" || errorResponse.ErrorCode != "" {
				return nil, &errorResponse
			}
			errorResponse.Raw = string(body)
			errorResponse.error = errors.New(http.StatusText(netResponseHolder.StatusCode))
			return nil, &errorResponse
		} else {
			return nil, &ErrorResponse{error: errors.New(resp.Status)}
		}
	}
	if err := json.NewDecoder(resp.Body).Decode(&netResponseHolder.Body); err != nil {
		return nil, &ErrorResponse{error: err}
	}
	return netResponseHolder, nil
}

func performSecurePostRequest[T any](payload interface{}, endpoint string, d *Daraja) (*networkResponse[T], *ErrorResponse){
	var headers = make(map[string]string)
	if d.authorization.AccessTokens == ""
}