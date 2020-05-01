package service

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/schoukri/joke-server/model"
)

// PersonServiceProvider is the interface that wraps the GetRandomPerson method.
type PersonServiceProvider interface {
	GetRandomPerson() (*model.Person, error)
}

// personService implements the PersonServiceProvider interface.
type personService struct {
	url        string
	httpClient *http.Client
}

// personResponse defines the response from the external Person service.
type personResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// NewPersonService returns a new personService.
func NewPersonService(url string) PersonServiceProvider {
	return &personService{
		url:        url,
		httpClient: NewClient(),
	}
}

// GetRandomPerson returns a random Person from the external Person service.
func (service *personService) GetRandomPerson() (*model.Person, error) {

	req, err := http.NewRequest("GET", service.url, nil)
	if err != nil {
		return nil, &RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.Wrap(err, "cannot build request"),
		}
	}

	addStandardHeaders(req)

	resp, err := service.httpClient.Do(req)

	if err != nil {
		return nil, &RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.Wrap(err, "person api request failed"),
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, &RequestError{
			StatusCode: resp.StatusCode,
			Err:        errors.Errorf("person api request returned status code %d", resp.StatusCode),
		}
	}

	personResp := personResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&personResp); err != nil {
		return nil, &RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.Wrap(err, "unmarshaling failed"),
		}
	}

	return personResp.toPerson(), nil
}

// toPerson converts the response from the external Person service to our internal Person model.
func (p *personResponse) toPerson() *model.Person {
	return &model.Person{
		FirstName: p.FirstName,
		LastName:  p.LastName,
	}
}
