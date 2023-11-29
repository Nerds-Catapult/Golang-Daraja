package darajaAuth

import (
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
}