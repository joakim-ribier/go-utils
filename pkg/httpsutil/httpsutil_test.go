package httpsutil

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/joakim-ribier/go-utils/pkg/jsonsutil"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

var exposeHost string = "0.0.0.0:3333"

func TestMain(m *testing.M) {
	if os.Getenv("ENV_MODE") == "ci" {
		// on Github action use directly the services container
		// without dockertest (to have more control on the steps)
		os.Exit(m.Run())
	} else {
		// uses a sensible default on windows (tcp/http) and linux/osx (socket)
		pool, err := dockertest.NewPool("")
		if err != nil {
			log.Fatalf("Could not construct pool: %s", err)
		}

		// uses pool to try to connect to Docker
		err = pool.Client.Ping()
		if err != nil {
			log.Fatalf("Could not connect to Docker: %s", err)
		}

		resource, err := pool.RunWithOptions(&dockertest.RunOptions{
			Repository:   "joakimribier/mockapic",
			Tag:          "latest",
			Env:          []string{"MOCKAPIC_PORT=3333"},
			ExposedPorts: []string{"3333"},
		}, func(config *docker.HostConfig) {
			// set AutoRemove to true so that stopped container goes away by itself
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
		})
		if err != nil {
			log.Fatalf("Could not start resource: %s", err)
		}

		resource.Expire(30) // hard kill the container in 3 minutes (180 Seconds)
		exposeHost = net.JoinHostPort("0.0.0.0", resource.GetPort("3333/tcp"))

		if err := pool.Retry(func() error {
			req, err := NewHttpRequest(fmt.Sprintf("http://%s/", exposeHost), "")
			if err != nil {
				return err
			}
			_, err = req.Timeout("150ms").Call()
			return err
		}); err != nil {
			log.Fatalf("Could not connect to mockapic server: %s", err)
		}

		code := m.Run()

		// cannot defer this because os.Exit doesn't care for defer
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}

		os.Exit(code)
	}
}

// TestCallHTTPS calls a https URL,
// checking for a valid return value.
func TestCallHTTPS(t *testing.T) {
	req, err := NewHttpRequest("https://go.dev/", "")
	if err != nil {
		t.Fatalf("Could not build a http request struct: %s", err)
	}

	resp, _ := req.InsecureSkipVerify().Call()
	if resp.StatusCode != 200 || err != nil {
		t.Errorf(`result: {%d} but expected: {%s}`, resp.StatusCode, "200")
	}
}

// TestCall calls httpsutil.NewHttpRequest(string, string).Call(),
// checking for a valid return value.
func TestCall(t *testing.T) {
	uuid := createNewRequest(exposeHost, "", 200, "Hello World")

	req, err := NewHttpRequest(fmt.Sprintf("http://%s/v1/%s", exposeHost, uuid), "")
	if err != nil {
		t.Fatalf("Could not build a http request struct: %s", err)
	}

	resp, err := req.AsJson().Call()
	if resp.StatusCode != 200 || err != nil || string(resp.Body) != "Hello World" {
		t.Errorf(`result: {%d} but expected: {%s}`, resp.StatusCode, "200")
	}
}

// TestCallWithTimeout calls httpsutil.NewHttpRequest(string, string).Call(),
// checking for a valid return value.
func TestCallWithTimeout(t *testing.T) {
	uuid := createNewRequest(exposeHost, "POST", 200, "Hello World")

	req, err := NewHttpRequest(fmt.Sprintf("http://%s/v1/%s?delay=50ms", exposeHost, uuid), "")
	if err != nil {
		t.Fatalf("Could not build a http request struct: %s", err)
	}

	resp, err := req.Timeout("10ms").AsJson().Call()
	if !strings.Contains(err.Error(), "(Client.Timeout exceeded while awaiting headers)") {
		t.Errorf(`result: {%v} but expected: {%v}`, resp, "Client.Timeout exceeded...")
	}
}

// TestCallWithNoTimeout calls httpsutil.NewHttpRequest(string, string).Call(),
// checking for a valid return value.
func TestCallWithNoTimeout(t *testing.T) {
	uuid := createNewRequest(exposeHost, "POST", 200, "Hello World")

	req, err := NewHttpRequest(fmt.Sprintf("http://%s/v1/%s?delay=100ms", exposeHost, uuid), "")
	if err != nil {
		t.Fatalf("Could not build a http request struct: %s", err)
	}

	resp, err := req.NoTimeout().AsJson().Call()
	if err != nil ||
		req.timeout != 0 ||
		resp.StatusCode != 200 ||
		resp.TimeInMillis < 100 || resp.TimeInMillis > 110 {
		t.Errorf(`result: {%v} but expected: {%s}`, err, "200")
	}
}

// TestCallAndTruncateBody calls httpsutil.NewHttpRequest(string, string).Call(),
// checking for a valid return value.
func TestCallAndTruncateBody(t *testing.T) {
	uuid := createNewRequest(exposeHost, "", 200, "Hello World")

	req, err := NewHttpRequest(fmt.Sprintf("http://%s/v1/%s", exposeHost, uuid), "")
	if err != nil {
		t.Fatalf("Could not build a http request struct: %s", err)
	}

	resp, err := req.AsJson().Call()

	if err != nil ||
		resp.StatusCode != 200 ||
		resp.TruncateBody(5) != "Hello" {
		t.Errorf(`result: {%v} but expected: {%s}`, resp, "Hello")
	}
}

func createNewRequest(hostAndPort string, method string, status int, body string) string {
	httpRequest, err := NewHttpRequest(fmt.Sprintf("http://%s/v1/new", hostAndPort), fmt.Sprintf(`{
    	"status": %d,
    	"contentType": "text/plain",
    	"charset": "UTF-8",
    	"body": "%s"
	}`, status, body))
	if err != nil {
		log.Fatalf("Could not create a http server: %s", err)
	}

	if method != "" {
		httpRequest.Method(method)
	}

	httpResponse, err := httpRequest.AsJson().Call()
	if err != nil {
		log.Fatalf("Could not create a http server: %s", err)
	}

	values, _ := jsonsutil.Unmarshal[map[string]interface{}](httpResponse.Body)
	return values["uuid"].(string)
}
