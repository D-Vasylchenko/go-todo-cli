package storage

import (
    "encoding/json"
    "os"
    "go-todo-cli/todo"
)

func LoadTasks(filename string) ([]todo.Task, error) {
    file, err := os.Open(filename)
    if err != nil {
        return []todo.Task{}, nil
    }
    defer file.Close()

    var tasks []todo.Task
    json.NewDecoder(file).Decode(&tasks)
    return tasks, nil
}

func SaveTasks(filename string, tasks []todo.Task) error {
    data, _ := json.MarshalIndent(tasks, "", "  ")
    return os.WriteFile(filename, data, 0644)
}
