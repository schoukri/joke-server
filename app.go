package main

import (
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_chi"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/schoukri/joke-server/service"
)

// App defines the dependencies necessary for our http server to retrieve data from external services and route incoming requests.
type App struct {
	personService service.PersonServiceProvider
	jokeService   service.JokeServiceProvider
	router        *chi.Mux
}

// NewApp returns an http server that will return a random joke for a random person.
func NewApp(personService service.PersonServiceProvider, jokeService service.JokeServiceProvider) *App {

	a := &App{
		personService: personService,
		jokeService:   jokeService,
	}

	// router
	router := chi.NewRouter()

	// TODO: in order to prevent exceeding the rate limit of 10 requests/second of the external Names API (http://uinames.com/api/)
	// I have limited this service to the same request rate. Before we could allow a higher request rate for our server, we would
	// first have to "workaround" the Names API. One idea is to cache the successful responses from the Names API and then use the
	// cache as a secondary source of random names when the Names API request limit has been exceeded.
	limiter := tollbooth.NewLimiter(10, nil)
	limiter.SetIPLookups([]string{"X-Forwarded-For", "RemoteAddr", "X-Real-IP"})
	router.Use(tollbooth_chi.LimitHandler(limiter))

	// middleware
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(10 * time.Second))
	router.Use(render.SetContentType(render.ContentTypePlainText))

	// routes
	router.Get("/", a.randomJokeHander())

	a.router = router

	return a

}

// Start starts the joke server on the specified address.
func (a *App) Start(addr string) error {
	return http.ListenAndServe(addr, a.router)
}

func (a *App) randomJokeHander() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// get a random person from the person Service
		person, err := a.personService.GetRandomPerson()
		// (if the service returns an error, return the original status code)
		if err != nil {
			re, ok := err.(*service.RequestError)
			if ok {
				http.Error(w, http.StatusText(re.StatusCode), re.StatusCode)
			} else {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		// get a random joke from the Joke service for the specified person
		// (if the service returns an error, return the original status code)
		joke, err := a.jokeService.GetRandomJoke(*person)
		if err != nil {
			re, ok := err.(*service.RequestError)
			if ok {
				http.Error(w, http.StatusText(re.StatusCode), re.StatusCode)
			} else {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		w.Write([]byte(joke.Value))

	}
}
