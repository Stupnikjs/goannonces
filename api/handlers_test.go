package api

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var theTests = []struct {
		name                    string
		url                     string
		expectedStatusCode      int
		expectedURL             string
		expectedFirstStatusCode int
	}{
		{"home", "/", http.StatusOK, "/", http.StatusOK},
		{"404", "/fish", http.StatusNotFound, "/fish", http.StatusNotFound},
	}

	routes := app.Routes()

	// create a test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// client without certificate
	client := http.Client{
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// range through the test data
	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}
		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s expected status %d, but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}

		if resp.Request.URL.Path != e.expectedURL {
			t.Errorf("%s: epxected final url of %s but got %s", e.name, e.expectedURL, resp.Request.URL.Path)
		}

		resp2, _ := client.Get(ts.URL + e.url)

		if resp2.StatusCode != e.expectedFirstStatusCode {
			t.Errorf("%s expected first returned status code to be %d but got %d", e.name, e.expectedFirstStatusCode, resp2.StatusCode)
		}

	}

}

func TestAppHome(t *testing.T) {
	var tests = []struct {
		name         string
		putInSession string
		expectedHTML string
	}{
		{"first visit", "", "<small>From Session:"},
		{"second visit", "hello world!", "<small>From Session:hello world!"},
	}

	for _, e := range tests {
		// create a request
		req, _ := http.NewRequest("GET", "/", nil)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(app.RenderAccueil)

		handler.ServeHTTP(rr, req)

		// check status code
		if rr.Code != http.StatusOK {
			t.Errorf("TestAppHome returned wrong status code; expected 200 but got %d", rr.Code)
		}

		// check the body

		body, _ := io.ReadAll(rr.Body)

		if !strings.Contains(string(body), e.expectedHTML) {
			t.Errorf("%s: did not find %s in response body", e.name, e.expectedHTML)
		}

	}

}

/*



mux.Get("/urls", app.ListObjectHandler)
mux.Post("/upload", app.UploadFile)


mux.Get("/stream/sound/{id}", app.StreamTrackFromGCPHandler)
test.mp3 pushed dans setup.go


mux.Post("/api/track/upload", app.UploadTrackListHandler)
mux.Post("/api/track/tag", app.UpdateTrackTagHandler)
mux.Post("/api/track/remove", app.DeleteTrackHandler)

*/
