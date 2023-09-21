package controllers

import (
	"TarefasCalendario/mappers"
	"TarefasCalendario/models"
	"TarefasCalendario/responder"
	"TarefasCalendario/services"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskPayload models.TaskPayload
	err := json.NewDecoder(r.Body).Decode(&taskPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Erro ao decodificar dados da tarefa:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func GetTasks(w http.ResponseWriter, r *http.Request) {

	tasks, err := services.GetAllTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Erro ao finalizar criação no banco de dados:", err)
		return
	}

	responder.JSON(w, r, mappers.MapTasksToPayload(tasks), http.StatusOK)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	tasksID := chi.URLParam(r, "TaskID")
	task, err := services.GetTask(tasksID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responder.JSON(w, r, task, http.StatusOK)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "taskID")

	err := services.DeleteTask(taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
