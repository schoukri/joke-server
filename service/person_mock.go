package service

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
	"github.com/schoukri/joke-server/model"
)

// mockPersonService is a mock implementation of the PersonServiceProvider interface
type mockPersonService struct{}

// NewMockPersonService returns a new mockPersonService (a mock implementation of the PersonServiceProvider interface)
func NewMockPersonService() PersonServiceProvider {
	return &mockPersonService{}
}

// GetRandomPerson returns a "random" Person
// (this mock implementation always returns the same person)
func (service *mockPersonService) GetRandomPerson() (*model.Person, error) {

	data := strings.NewReader(`{"first_name":"Hasina","last_name":"Tanweer"}`)

	personResp := personResponse{}
	if err := json.NewDecoder(data).Decode(&personResp); err != nil {
		return nil, errors.Wrap(err, "unmarshaling failed")
	}

	return personResp.toPerson(), nil
}
