package service_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/schoukri/joke-server/model"
	"github.com/schoukri/joke-server/service"
	"github.com/stretchr/testify/assert"
)

// TestJokeServiceGetRandomJoke tests the real JokeService, but connected to a test http server.
func TestJokeServiceGetRandomJoke(t *testing.T) {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{ "type": "success", "value": { "id": 181, "joke": "John Doe's OSI network model has only one layer - Physical.", "categories": ["nerdy"] } }`))
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	service := service.NewJokeService(server.URL)

	person := model.Person{FirstName: "John", LastName: "Doe"}
	joke, err := service.GetRandomJoke(person)

	assert.Nil(t, err)
	assert.Equal(t, 181, joke.ID)
	assert.Equal(t, "John Doe's OSI network model has only one layer - Physical.", joke.Value)
}

// TestMockJokeServiceGetRandomJoke tests a mock implementation of the JokeServiceProvider interface.
func TestMockJokeServiceGetRandomJoke(t *testing.T) {

	service := service.NewMockJokeService()

	person := model.Person{FirstName: "Chuck", LastName: "Norris"}
	joke, err := service.GetRandomJoke(person)

	assert.Nil(t, err)
	assert.Equal(t, 181, joke.ID)
	assert.Equal(t, "Chuck Norris's OSI network model has only one layer - Physical.", joke.Value)
}
