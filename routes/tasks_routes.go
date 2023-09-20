package routes

import (
	"TarefasCalendario/controllers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func LoadRoutes(r *chi.Mux) {
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/tasks", controllers.GetTasks)
	r.Post("/tasks", controllers.CreateTask)
	r.Get("/tasks/{taskID}", controllers.GetTask)
	r.Delete("/tasks/{taskID}", controllers.DeleteTask)

	// Docs
	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

}
