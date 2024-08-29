package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/pol-cova/task-tracker-cli/internal/models"
	"github.com/spf13/cobra"
	"strconv"
)

const file = "tasks.json"

// rootCMD
var rootCMD = &cobra.Command{
	Use:   "tasky",
	Short: "Tasky is a simple CLI to manage your tasks",
	Long:  `Tasky is a simple CLI to manage your tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		// Print help
		cmd.Help()
	},
}

func Execute() error {
	return rootCMD.Execute()
}

func init() {
	rootCMD.AddCommand(addCMD)
	rootCMD.AddCommand(listCMD)
	rootCMD.AddCommand(updateCMD)
	rootCMD.AddCommand(deleteCMD)
	rootCMD.AddCommand(markDoneCMD)
	rootCMD.AddCommand(markInProgress)
}

var addCMD = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Add task
		t, err := models.Add(args[0], file)
		if err != nil {
			panic(err)
		}
		// Return Task added successfully (ID: Task id)
		fmt.Printf("Task added successfully (ID: %d)", t.ID)
	},
}

var listCMD = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all tasks`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []models.Task
		var err error

		if len(args) == 1 {
			// Get tasks by status
			tasks, err = models.GetByStatus(args[0], file)
		} else {
			// Get all tasks
			tasks, err = models.GetAll(file)
		}

		if err != nil {
			panic(err)
		}

		// check if task are null
		if len(tasks) == 0 {
			fmt.Println("No tasks with the specified status :(")
			return
		}

		// Marshal tasks to JSON
		tasksJSON, err := json.MarshalIndent(tasks, "", "  ")
		if err != nil {
			panic(err)
		}

		// Print JSON string
		fmt.Println(string(tasksJSON))
	},
}

var updateCMD = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Long:  `Update a task`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Update task
		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		_, err = models.Update(id, args[1], file)
		if err != nil {
			panic(err)
		}

		// Return Task updated successfully
		fmt.Println("Task updated successfully")
	},
}

var deleteCMD = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  `Delete a task`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Delete task
		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		err = models.Delete(id, file)
		if err != nil {
			panic(err)
		}

		// Return Task deleted successfully
		fmt.Println("Task deleted successfully")
	},
}

var markDoneCMD = &cobra.Command{
	Use:   "mark-done",
	Short: "Mark a task as done",
	Long:  `Mark a task as done`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Mark done task
		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		t, err := models.MarkDone(id, file)
		if err != nil {
			panic(err)
		}

		// Return Task marked as done successfully
		fmt.Printf("Task marked as done successfully (ID: %d)", t.ID)
	},
}

var markInProgress = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "Mark a task as in progress",
	Long:  `Mark a task as in progress`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Mark in progress task
		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		t, err := models.MarkInProgress(id, file)
		if err != nil {
			panic(err)
		}

		// Return Task marked as in progress successfully
		fmt.Printf("Task marked as in progress successfully (ID: %d)", t.ID)
	},
}
