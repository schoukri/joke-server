package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/schoukri/joke-server/service"
	"github.com/stretchr/testify/assert"
)

// TestAppServer tests our internal App server using a test http server mock external Person and Joke services.
func TestAppServer(t *testing.T) {

	// setup our app with Mock services
	app := NewApp(
		service.NewMockPersonService(),
		service.NewMockJokeService(),
	)

	// create a test http server that will serve our routes
	ts := httptest.NewServer(app.router)
	defer ts.Close()

	// get a random joke from the test server
	// (not so random -- it will return the same one every time)
	resp, err := http.Get(ts.URL + "/")
	assert.Nil(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("Content-Type"))

	respBody, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, string(respBody), `Hasina Tanweer's OSI network model has only one layer - Physical.`)

	// GET an unknown resource (should see a Not Found error)
	resp2, err := http.Get(ts.URL + "/foobar")
	assert.Nil(t, err)
	defer resp2.Body.Close()
	assert.Equal(t, http.StatusNotFound, resp2.StatusCode)

	// POST to "/" (should get an Unsupported Method error)
	resp3, err := http.Post(ts.URL+"/", "", nil)
	assert.Nil(t, err)
	defer resp3.Body.Close()
	assert.Equal(t, http.StatusMethodNotAllowed, resp3.StatusCode)

}
