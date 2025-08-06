package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	router := chi.NewRouter()
	router.Use(middleware.Logger) // Middleware PRIMEIRO

	return &WebServer{
		Router:        router,
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebServer) AddHandlerWithMethod(path string, handler http.HandlerFunc, method string) {
	switch method {
	case "GET":
		s.Router.Get(path, handler)
	case "POST":
		s.Router.Post(path, handler)
	case "PUT":
		s.Router.Put(path, handler)
	case "DELETE":
		s.Router.Delete(path, handler)
	default:
		s.Router.Handle(path, handler)
	}
}

// loop through the handlers and add them to the router
// start the server
func (s *WebServer) Start() {
	if len(s.Handlers) > 0 {
		for path, handler := range s.Handlers {
			s.Router.Handle(path, handler)
		}
	}

	http.ListenAndServe(s.WebServerPort, s.Router)
}
