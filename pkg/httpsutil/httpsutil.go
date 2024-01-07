package httpsutil

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

// HttpRequest struct helps to build the http.Request
type HttpRequest struct {
	Req     *http.Request
	timeout int
}

// HttpResponse HTTP response
type HttpResponse struct {
	Status     string
	StatusCode int

	Body []byte
}

// NewHttpRequest builds new http.Request object ('GET' by default)
func NewHttpRequest(url string, body string) (*HttpRequest, error) {
	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}
	return &HttpRequest{
		Req:     req,
		timeout: 15,
	}, nil
}

// AsJson adds "Content-Type: application/json" header
func (r *HttpRequest) AsJson() *HttpRequest {
	r.Headers(map[string]string{"Content-Type": "application/json"})
	return r
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

// Timeout sets a timeout in seconds
func (r *HttpRequest) Timeout(seconds int) *HttpRequest {
	if seconds > -1 {
		r.timeout = seconds
	}
	return r
}

// Call sends the HTTP request and returns the HTTP response
func (r *HttpRequest) Call() (*HttpResponse, error) {
	timeout := time.Duration(r.timeout) * time.Second
	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(r.Req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &HttpResponse{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Body:       body,
	}, nil
}
