package services

import (
	"TarefasCalendario/db"
	"TarefasCalendario/models"
	"context"
	"fmt"
)

func GetTask(taskID string) (*models.Task, error) {
	db := db.GetDatabase()

	sqlTask := `
		SELECT id, name, date, time FROM tasks WHERE id = $1;
	`

	task := &models.Task{}
	err := db.QueryRow(context.Background(), sqlTask, taskID).Scan(&task.ID, &task.Name, &task.Date, &task.Time)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Erro ao obter tarefa: %s", err)
	}

	return task, nil
}

func CreateTask(task *models.Task) (*models.Task, error) {
	db := db.GetDatabase()

	sqlTask := `
		INSERT INTO tasks (id, name, date, time) VALUES ($1, $2, $3, $4);
	`

	_, err := db.Exec(context.Background(), sqlTask, task.ID, task.Name, task.Date, task.Time)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Erro ao criar tarefa: %s", err)
	}

	return task, nil
}

func GetAllTasks() ([]*models.Task, error) {
	db := db.GetDatabase()

	tasks := []*models.Task{}

	sqlTask := `
		SELECT id, name, date, time FROM tasks;
	`

	rows, err := db.Query(context.Background(), sqlTask)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Erro ao obter tarefas: %s", err)
	}

	for rows.Next() {
		task := &models.Task{}
		rows.Scan(&task.ID, &task.Name, &task.Date, &task.Time)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func DeleteTask(taskID string) error {
	db := db.GetDatabase()

	sqlTask := `
		DELETE FROM tasks WHERE id = $1;
	`

	_, err := db.Exec(context.Background(), sqlTask, taskID)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Erro ao deletar tarefa: %s", err)
	}

	return nil
}
