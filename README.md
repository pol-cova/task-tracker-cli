# Task Tracker Project

**Task Tracker** is a project from backend `roadmap.sh`

## Features

- **Add**: Create a new task.
- **Update**: Modify an existing task's description.
- **Delete**: Remove a task from the list.
- **Mark as Done**: Mark a task as completed.
- **Mark as In Progress**: Mark a task as currently being worked on.
- **List Tasks**: View all tasks, filter by status (todo, in-progress, done).

## Task Properties

Each task in the Task Tracker has the following properties:

- **id**: A unique identifier for the task.
- **description**: A short description of the task.
- **status**: The current status of the task (`todo`, `in-progress`, `done`).
- **createdAt**: The date and time when the task was created.
- **updatedAt**: The date and time when the task was last updated.

## Getting Started

### Prerequisites

- A compatible programming language environment (e.g., Go).
- Command-line interface.

### Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/pol-cova/task-tracker-cli.git
    cd task-tracker-cli
    ```

2. **Build the project** (if applicable, based on your programming language):
    ```bash
    go build -o tasky
    ```

3. **Run the CLI**:
    ```bash
    ./tasky
    ```

### Commands

- **Add a Task**:
    ```bash
    ./tasky add "Your task description"
    ```

- **Update a Task**:
    ```bash
    ./tasky update <task_id> "New task description"
    ```

- **Delete a Task**:
    ```bash
    ./tasky delete <task_id>
    ```

- **Mark a Task as Done**:
    ```bash
    ./tasky mark-done <task_id>
    ```

- **Mark a Task as In Progress**:
    ```bash
    ./tasky mark-in-progress <task_id>
    ```

- **List All Tasks**:
    ```bash
    ./tasky list
    ```

- **List Tasks by Status** (e.g., `todo`, `in-progress`, `done`):
    ```bash
    ./tasky list <status>
    ```

## Implementation Details

- **Data Storage**: Tasks are stored in a JSON file (`tasks.json`) located in the current directory. This file is created automatically if it does not exist.
- **Error Handling**: The application gracefully handles errors and edge cases to ensure smooth operation.

## Project Roadmap url

https://roadmap.sh/projects/task-tracker