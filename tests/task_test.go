package tests

func TestAddTask(t *testing.T) {
    tasks := []todo.Task{}
    task := todo.AddTask(tasks, "Test task")

    if task.ID != 1 {
        t.Errorf("Expected ID 1, got %d", task.ID)
    }
    if task.Name != "Test task" {
        t.Errorf("Expected name 'Test task', got '%s'", task.Name)
    }
    if task.Done {
        t.Error("New task should not be marked as done")
    }
}

func TestMarkTaskDone(t *testing.T) {
    tasks := []todo.Task{
        {ID: 1, Name: "Task 1", Done: false},
        {ID: 2, Name: "Task 2", Done: false},
    }

    todo.MarkTaskDone(tasks, 2)

    if !tasks[1].Done {
        t.Errorf("Expected task with ID 2 to be marked as done")
    }
}
