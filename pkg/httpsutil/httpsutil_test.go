package httpsutil

import (
	"testing"
)

// TestCallGetURL calls httpsutil.NewHttpRequest.Call on {GET} URL,
// checking for a valid return value.
func TestCallGetURL(t *testing.T) {
	if r, err := NewHttpRequest("https://github.com/joakim-ribier", ""); r != nil {
		if r, err := r.Call(); r.StatusCode != 200 {
			if err != nil {
				t.Log(err)
			}
			t.Errorf(`result: {%d} but expected: {%s}`, r.StatusCode, "200")
		}
	} else {
		t.Errorf(`result: {%s} but expected: {%s}`, err, "200")
	}
}
