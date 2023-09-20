// models/task.go

package models

import (
	"TarefasCalendario/db"
	"context"
	"fmt"
)

type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
	Time string `json:"time"`
}

type TaskPayload struct {
	Name string `json:"name"`
	Date string `json:"date"`
	Time string `json:"time"`
}

func CreateTasksTable() error {
	sql := `
		CREATE TABLE IF NOT EXISTS "tasks" (
			"id" serial PRIMARY KEY,
			"name" text NOT NULL,
			"date" date NOT NULL,
			"time" time NOT NULL
		);
	`

	fmt.Println("Criando Tabela de tarefas")

	_, err := db.GetDatabase().Exec(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("erro ao criar tabela de tarefas: %s", err)
	} else {
		fmt.Println("Tabela de tarefas criada")
	}

	return nil
}
