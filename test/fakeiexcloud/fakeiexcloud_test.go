package fakeiexcloud

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNominalUsage(t *testing.T) {
	// Create a fake client and run a basic 'get' request.
	f := FakeIEXCloud{ResponseJSON: "blah"}
	s := httptest.NewServer(http.HandlerFunc(f.Handle))
	defer s.Close()

	c := http.Client{}
	resp, err := c.Get("http://" + s.Listener.Addr().String() + "/path")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Check the request and responses received.
	if got, want := f.LastURLReceived.Path, "/path"; got != want {
		t.Fatalf("Got url path %v, want %v", got, want)
	}
	if got, want := string(body), "blah"; got != want {
		t.Fatalf("Got body %v, want %v", got, want)
	}
}

func TestErrorInjection(t *testing.T) {
	// The error should take precedence over the JSON.
	f := FakeIEXCloud{ResponseJSON: "blah", ResponseHTTPStatus: http.StatusNotFound}
	s := httptest.NewServer(http.HandlerFunc(f.Handle))
	defer s.Close()

	c := http.Client{}
	resp, err := c.Get("http://" + s.Listener.Addr().String() + "/error-path")
	if err != nil {
		t.Fatal(err)
	}
	if got, want := resp.StatusCode, http.StatusNotFound; got != want {
		t.Fatalf("Got status %v, want %v", got, want)
	}
	if got, want := f.LastURLReceived.Path, "/error-path"; got != want {
		t.Fatalf("Got url path %v, want %v", got, want)
	}
}
