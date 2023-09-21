package mappers

import "TarefasCalendario/models"

func MapTasksToPayload(tasks []*models.Task) []*models.TaskPayload {
	taskPayloads := make([]*models.TaskPayload, len(tasks))
	for i, task := range tasks {
		taskPayloads[i] = MapTaskToPayload(task)
	}
	return taskPayloads
}

func MapTaskToPayload(task *models.Task) *models.TaskPayload {
	return &models.TaskPayload{
		ID:   task.ID,
		Name: task.Name,
		Date: task.Date,
		Time: task.Time,
	}
}
