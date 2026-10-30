package api

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
)

type HTTPserver struct {
	httpHandlers *HTTPHandler
}

func NewHTTPServer(handler *HTTPHandler) *HTTPserver {
	return &HTTPserver{httpHandlers: handler}
}

func (s *HTTPserver) StartServer() error {
	router := chi.NewRouter()
	router.Post("/tasks", s.httpHandlers.HandlerCreateTask)
	router.Get("/tasks/{title}", s.httpHandlers.HandlerGetTask)
	router.Get("/tasks", s.httpHandlers.HandlerGetAllTasks)
	router.Patch("/tasks/{title}", s.httpHandlers.HandlerCompleteTask)
	router.Delete("/tasks/{title}", s.httpHandlers.HandlerDeleteTask)

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
