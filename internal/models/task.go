package models

import (
	"errors"
	"github.com/pol-cova/task-tracker-cli/internal/fileManager"
	"time"
)

var tasks []Task

const (
	// Status constants
	TASK_STATUS_TODO        = "todo"
	TASK_STATUS_DONE        = "done"
	TASK_STATUS_IN_PROGRESS = "in-progress"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Get all tasks method in json format
func GetAll(file string) ([]Task, error) {
	err := fileManager.ReadTasks(file, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// Get all task by status method
func GetByStatus(status string, file string) ([]Task, error) {
	err := fileManager.ReadTasks(file, &tasks)
	if err != nil {
		return nil, err
	}

	var tasksByStatus []Task
	for _, task := range tasks {
		if task.Status == status {
			tasksByStatus = append(tasksByStatus, task)
		}
	}
	return tasksByStatus, nil
}

// Add a new task method
func Add(description string, file string) (Task, error) {
	err := fileManager.ReadTasks(file, &tasks)
	if err != nil {
		return Task{}, err
	}

	// Generate a new ID
	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	// Validate the description
	if description == "" {
		return Task{}, errors.New("description cannot be empty")
	}

	// Create the new task
	newTask := Task{
		ID:          newID,
		Description: description,
		Status:      TASK_STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Append the new task to the tasks slice
	tasks = append(tasks, newTask)

	// Save the tasks
	err = fileManager.SaveTasks(file, tasks)
	if err != nil {
		return Task{}, err
	}
	return newTask, nil
}

// Update a task method
func Update(id int, description string, file string) (Task, error) {
	err := fileManager.ReadTasks(file, &tasks)
	if err != nil {
		return Task{}, err
	}

	// Find the task
	var task Task
	for i, t := range tasks {
		if t.ID == id {
			task = tasks[i]
			break
		}
	}

	// Validate the task
	if task.ID == 0 {
		return Task{}, errors.New("task not found")
	}

	// Validate the description
	if description == "" {
		return Task{}, errors.New("description cannot be empty")
	}

	// Update the task
	task.Description = description
	task.UpdatedAt = time.Now()

	// Save the tasks
	err = fileManager.SaveTasks(file, tasks)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

// Delete a task method
func Delete(id int, file string) error {
	err := fileManager.ReadTasks(file, &tasks)
	if err != nil {
		return err
	}

	// Find the task
	var task Task
	for i, t := range tasks {
		if t.ID == id {
			task = tasks[i]
			break
		}
	}

	// Validate the task
	if task.ID == 0 {
		return errors.New("task not found")
	}

	// Remove the task
	tasks = append(tasks[:task.ID-1], tasks[task.ID:]...)

	// Save the tasks
	err = fileManager.SaveTasks(file, tasks)
	if err != nil {
		return err
	}
	return nil
}

// Mark done a task method
func MarkDone(id int, file string) (Task, error) {
	err := fileManager.ReadTasks(file, &tasks)
	if err != nil {
		return Task{}, err
	}

	// Find the task
	var task *Task
	for i := range tasks {
		if tasks[i].ID == id {
			task = &tasks[i]
			break
		}
	}

	// Validate the task
	if task == nil {
		return Task{}, errors.New("task not found")
	}

	// Update the task
	task.Status = TASK_STATUS_DONE
	task.UpdatedAt = time.Now()

	// Save the tasks
	err = fileManager.SaveTasks(file, tasks)
	if err != nil {
		return Task{}, err
	}
	return *task, nil
}

// Mark in progress a task method
func MarkInProgress(id int, file string) (Task, error) {
	err := fileManager.ReadTasks(file, &tasks)
	if err != nil {
		return Task{}, err
	}

	// Find the task
	var task *Task
	for i := range tasks {
		if tasks[i].ID == id {
			task = &tasks[i]
			break
		}
	}

	// Validate the task
	if task == nil {
		return Task{}, errors.New("task not found")
	}

	// Update the task
	task.Status = TASK_STATUS_IN_PROGRESS
	task.UpdatedAt = time.Now()

	// Save the tasks
	err = fileManager.SaveTasks(file, tasks)
	if err != nil {
		return Task{}, err
	}
	return *task, nil
}
