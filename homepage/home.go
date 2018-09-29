package homepage

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

const message = "Server is online"

//Handlers expects a logger to be injected
type Handlers struct {
	logger *log.Logger
}

//SetupRoutes is where we setup all routing
//associated in this domain/package
func (h *Handlers) SetupRoutes(router *chi.Mux) {
	router.Get("/", h.Logger(h.Home))
}

//Home is the homepage handler
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

//Logger is an example of a generic middleware
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer h.logger.Printf("Logger Middleware Sample")
		next(w, r)
	}
}

//New returns a new Specific
//Domain Handler with the injected logger
func New(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
