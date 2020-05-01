package service_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/schoukri/joke-server/service"
	"github.com/stretchr/testify/assert"
)

// TestPersonServiceGetRandomPerson tests the real PersonService, but connected to a test http server.
func TestPersonServiceGetRandomPerson(t *testing.T) {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"first_name":"Hasina","last_name":"Tanweer"}`))
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	service := service.NewPersonService(server.URL)

	person, err := service.GetRandomPerson()

	assert.Nil(t, err)
	assert.Equal(t, "Hasina", person.FirstName)
	assert.Equal(t, "Tanweer", person.LastName)
}

// TestMockPersonServiceGetRandomPerson tests a mock implementation of the JokeServiceProvider interface.
func TestMockPersonServiceGetRandomPerson(t *testing.T) {

	service := service.NewMockPersonService()

	person, err := service.GetRandomPerson()

	assert.Nil(t, err)
	assert.Equal(t, "Hasina", person.FirstName)
	assert.Equal(t, "Tanweer", person.LastName)
}
