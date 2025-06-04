package main

import (
    "fmt"
    "os"
    "strconv"

    "go-todo-cli/storage"
    "go-todo-cli/todo"
)

func main() {
    tasks, _ := storage.LoadTasks("tasks.json")

    if len(os.Args) < 2 {
        fmt.Println("Usage: [add|list|done|delete|clear] [arguments]")
        return
    }

    command := os.Args[1]

    switch command {
    case "add":
        if len(os.Args) < 3 {
            fmt.Println("Please provide a task name.")
            return
        }
        name := os.Args[2]
        task := todo.AddTask(tasks, name)
        tasks = append(tasks, task)

    case "list":
        todo.ListTasks(tasks)

    case "done":
        if len(os.Args) < 3 {
            fmt.Println("Please provide task ID.")
            return
        }
        id, _ := strconv.Atoi(os.Args[2])
        todo.MarkTaskDone(tasks, id)

    case "delete":
        if len(os.Args) < 3 {
            fmt.Println("Please provide task ID.")
            return
        }
        id, _ := strconv.Atoi(os.Args[2])
        tasks = todo.DeleteTask(tasks, id)

    case "clear":
        tasks = todo.ClearCompleted(tasks)
    }

    storage.SaveTasks("tasks.json", tasks)
}
