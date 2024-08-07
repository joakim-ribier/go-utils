package httpsutil

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
	"time"

	"github.com/joakim-ribier/go-utils/pkg/timesutil"
)

// HttpRequest struct helps to build the http.Request
type HttpRequest struct {
	Req                *http.Request
	timeout            time.Duration
	insecureSkipVerify bool
	transport          http.RoundTripper
}

// HttpResponse HTTP response
type HttpResponse struct {
	Status        string
	StatusCode    int
	ContentLength int64
	TimeInMillis  int64
	Body          []byte
}

// NewHttpRequest builds new http.Request object ('GET' by default)
func NewHttpRequest(url, body string) (*HttpRequest, error) {
	method := "GET"
	if len(body) > 0 {
		method = "POST"
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}

	return &HttpRequest{
		Req:                req,
		timeout:            15 * time.Second,
		insecureSkipVerify: false,
		transport:          nil,
	}, nil
}

// AsJson adds "Content-Type: application/json" header
func (r *HttpRequest) AsJson() *HttpRequest {
	return r.Header("Content-Type", "application/json")
}

// Header adds new header
func (r *HttpRequest) Header(key string, value string) *HttpRequest {
	r.Headers(map[string]string{key: value})
	return r
}

// Method defines method
func (r *HttpRequest) Method(method string) *HttpRequest {
	r.Req.Method = method
	return r
}

// SetBasicAuth sets basic authentication
func (r *HttpRequest) SetBasicAuth(username string, password string) *HttpRequest {
	r.Req.SetBasicAuth(username, password)
	return r
}

// Header adds new headers
func (r *HttpRequest) Headers(params map[string]string) *HttpRequest {
	for k, v := range params {
		r.Req.Header.Set(k, v)
	}
	return r
}

// Timeout sets a timeout, a {timeout} of zero means no timeout
func (r *HttpRequest) Timeout(timeout string) *HttpRequest {
	if v, err := time.ParseDuration(timeout); err == nil {
		r.timeout = v
	}
	return r
}

// NoTimeout sets a timeout of zero
func (r *HttpRequest) NoTimeout() *HttpRequest {
	r.timeout = 0
	return r
}

// InsecureSkipVerify sets {insecureSkipVerify} to {true}
func (r *HttpRequest) InsecureSkipVerify() *HttpRequest {
	r.insecureSkipVerify = true
	return r
}

// Transport sets a {transport} value
func (r *HttpRequest) Transport(t *http.Transport) *HttpRequest {
	r.transport = t
	return r
}

func (r *HttpRequest) getTransport() http.RoundTripper {
	var transport http.RoundTripper = nil
	if t := r.transport; t != nil {
		transport = t
	} else {
		if r.insecureSkipVerify {
			transport = &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
		}
	}
	return transport
}

// Call sends the HTTP request and returns the HTTP response
func (r *HttpRequest) Call() (*HttpResponse, error) {
	client := &http.Client{
		Timeout:   r.timeout,
		Transport: r.getTransport(),
	}

	resp, err := timesutil.WithExecutionTime(func() (*http.Response, error) {
		return client.Do(r.Req)
	})
	if err != nil {
		return nil, err
	}
	defer resp.T.Body.Close()

	body, err := io.ReadAll(resp.T.Body)
	if err != nil {
		return nil, err
	}

	return &HttpResponse{
		Status:        resp.T.Status,
		StatusCode:    resp.T.StatusCode,
		Body:          body,
		ContentLength: resp.T.ContentLength,
		TimeInMillis:  resp.TimeInMillis,
	}, nil
}

// TruncateBody transforms {r.Body} to string and truncates it with the {max}.
func (r *HttpResponse) TruncateBody(max int) string {
	if len(r.Body) > max {
		return string(r.Body[:max])
	}
	return string(r.Body)
}
