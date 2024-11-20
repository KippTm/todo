package main

import (
    "time"
    "fmt"
    "errors"
    "os"
    "github.com/aquasecurity/table"
    "strconv"
)

type Todo struct {
    Task string
    Completed bool
    CreatedAt time.Time
    CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(task string) {
    todo := Todo {
        Task:  task,
        Completed:  false,
        CreatedAt:  time.Now(),
        //CompletedAt:  nil,
    }
    *todos = append(*todos, todo)
}

func (todos *Todos) multiadd(tasks []string) {
    for _, t := range tasks {
        todos.add(t)
    }
}

func (todos *Todos) validator(index int) error {
    if index < 0 || index >= len(*todos) {
        err := errors.New("Invalid index")
        fmt.Println(err)
        return err
    }
    return nil
}

func (todos *Todos) delete(index int) error {
    t := *todos
    err := t.validator(index)
    if err != nil {
        return err
    }

    *todos = append(t[:index], t[index+1:]...)
    return nil
}

func (todos *Todos) toggle(index int) error {
    t := *todos
    err := t.validator(index)
    if err != nil {
        return err
    }

    t[index].Completed = !t[index].Completed 
    return nil 
}

func (todos *Todos) edit(index int, task string) error {
    t := *todos 
    err := t.validator(index)
    if err != nil {
        return err
    }

    t[index].Task = task
    return nil
}

func (todos *Todos) print(){
    table := table.New(os.Stdout)
    //table.SetRowLines(false)
    table.SetHeaders("#", "Task", "Completed", "Created At")
    
    for i, t := range *todos {
        completed := "--->🟥<---"
        
        if t.Completed {
            completed = "--->🟩<---"
        }

        table.AddRow(strconv.Itoa(i), t.Task, completed, t.CreatedAt.Format(time.RFC822))
    }

    table.Render()
}

func (todos *Todos) clear() {
    *todos = nil
}
