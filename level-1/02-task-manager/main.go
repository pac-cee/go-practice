package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

type TaskManager struct {
	tasks    []Task
	filename string
}

func NewTaskManager(filename string) *TaskManager {
	return &TaskManager{
		tasks:    make([]Task, 0),
		filename: filename,
	}
}

func (tm *TaskManager) LoadTasks() error {
	data, err := ioutil.ReadFile(tm.filename)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &tm.tasks)
}

func (tm *TaskManager) SaveTasks() error {
	data, err := json.MarshalIndent(tm.tasks, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(tm.filename, data, 0644)
}

func (tm *TaskManager) AddTask(title, description string, dueDate time.Time) {
	task := Task{
		ID:          len(tm.tasks) + 1,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Status:      "pending",
	}
	tm.tasks = append(tm.tasks, task)
	tm.SaveTasks()
}

func (tm *TaskManager) ListTasks() {
	if len(tm.tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("\nTasks:")
	fmt.Println("----------------------------------------")
	for _, task := range tm.tasks {
		fmt.Printf("ID: %d\n", task.ID)
		fmt.Printf("Title: %s\n", task.Title)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Printf("Due Date: %s\n", task.DueDate.Format("2006-01-02"))
