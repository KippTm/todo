package main

func main() {
    todos := Todos{}
    storage := NewStorage[Todos]("/fun/storage/tasks/todos.json")
    storage.Load(&todos)

    cmdFlags := NewCommandFlags()
    cmdFlags.Execute(&todos)

    storage.Save(todos)
}
