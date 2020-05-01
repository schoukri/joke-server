package service

import (
	"bytes"
	"encoding/json"
	"text/template"

	"github.com/pkg/errors"
	"github.com/schoukri/joke-server/model"
)

// mockJokeService is a mock implementation of the JokeServiceProvider interface
type mockJokeService struct{}

// NewMockJokeService returns a new mockJokeService (a mock implementation of the JokeServiceProvider interface)
func NewMockJokeService() JokeServiceProvider {
	return &mockJokeService{}
}

// GetRandomJoke returns a "random" joke for the specified person
// (this mock implementation always returns the same joke, but changes it for the specified person)
func (service *mockJokeService) GetRandomJoke(person model.Person) (*model.Joke, error) {

	tmpl := template.Must(template.New("").Parse(`{
		"type": "success",
		"value": {
			"id": 181,
			"joke": "{{ .FirstName }} {{ .LastName }}'s OSI network model has only one layer - Physical.", "categories": ["nerdy"]
		}
	}`))

	var data bytes.Buffer
	if err := tmpl.Execute(&data, person); err != nil {
		return nil, errors.Wrap(err, "executing template")
	}

	jokeResp := jokeResponse{}
	if err := json.NewDecoder(&data).Decode(&jokeResp); err != nil {
		return nil, errors.Wrap(err, "unmarshaling failed")
	}

	return jokeResp.toJoke(), nil
}
