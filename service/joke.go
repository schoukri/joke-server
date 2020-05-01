package service

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/schoukri/joke-server/model"
)

// JokeServiceProvider is the interface that wraps the GetRandomJoke method.
type JokeServiceProvider interface {
	GetRandomJoke(person model.Person) (*model.Joke, error)
}

// jokeService implements the JokeServiceProvider interface.
type jokeService struct {
	url        string
	httpClient *http.Client
}

// jokeResponse defines the response from the external Joke service.
type jokeResponse struct {
	Type  string `json:"type"`
	Value struct {
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}

// NewJokeService returns a new jokeService.
func NewJokeService(url string) JokeServiceProvider {
	return &jokeService{
		url:        url,
		httpClient: NewClient(),
	}
}

// GetRandomJoke returns a random Joke from the external Joke service, for the specified Person.
func (service *jokeService) GetRandomJoke(person model.Person) (*model.Joke, error) {

	serviceURL, err := url.Parse(service.url + "/random")
	if err != nil {
		return nil, errors.Wrap(err, "could not parse service url")
	}

	// build the query string
	query := url.Values{}
	query.Set("firstName", person.FirstName)
	query.Set("lastName", person.LastName)
	query.Set("limitTo", `\[nerdy\]`)
	serviceURL.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", serviceURL.String(), nil)
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
			Err:        errors.Wrap(err, "joke api request failed"),
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, &RequestError{
			StatusCode: resp.StatusCode,
			Err:        errors.Errorf("joke api request returned status code %d", resp.StatusCode),
		}
	}

	jokeResp := jokeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&jokeResp); err != nil {
		return nil, &RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.Wrap(err, "unmarshaling failed"),
		}
	}

	return jokeResp.toJoke(), nil
}

// toJoke converts the response from the external Joke service to our internal Joke model.
func (j *jokeResponse) toJoke() *model.Joke {
	return &model.Joke{
		ID:    j.Value.ID,
		Value: j.Value.Joke,
	}
}
