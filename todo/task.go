package todo

import (
    "fmt"
)

type Task struct {
    ID   int
    Name string
    Done bool
}

func AddTask(tasks []Task, name string) Task {
    id := len(tasks) + 1
    return Task{ID: id, Name: name, Done: false}
}

func ListTasks(tasks []Task) {
    for _, task := range tasks {
        status := " "
        if task.Done {
            status = "✔"
        }
        fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Name)
    }
}

func MarkTaskDone(tasks []Task, id int) {
    for i := range tasks {
        if tasks[i].ID == id {
            tasks[i].Done = true
            return
        }
    }
}

func DeleteTask(tasks []Task, id int) []Task {
    var updated []Task
    for _, task := range tasks {
        if task.ID != id {
            updated = append(updated, task)
        }
    }
    return updated
}

func ClearCompleted(tasks []Task) []Task {
    var result []Task
    for _, t := range tasks {
        if !t.Done {
            result = append(result, t)
        }
    }
    return result
}
