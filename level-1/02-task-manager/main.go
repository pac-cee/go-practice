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
