package main

import (
    "flag"
    "os"
    "fmt"
    "strings"
    "strconv"
)

type cmdFlags struct {
    Add string
    MultiAdd string
    Del int
    Edit string
    Toggle int
    List bool
    Clear bool
}

func NewCommandFlags() *cmdFlags {
    cf := cmdFlags{}

    flag.StringVar(&cf.Add, "a", "", "Add new task, specify title")
    flag.StringVar(&cf.Edit, "e", "", "Edit a task")
    flag.StringVar(&cf.MultiAdd, "m", "", "Add multiple tasks, seperated by ','")
    flag.IntVar(&cf.Del, "d", -1, "Delete a task")
    flag.IntVar(&cf.Toggle, "t", -1, "Change status of a task")
    flag.BoolVar(&cf.List, "l", false, "Show list of tasks")
    flag.BoolVar(&cf.Clear, "c", false, "Empty list of tasks")

    flag.Parse()
    return &cf
}

func (cf *cmdFlags) Execute (tasks *Todos){
    switch {
        case cf.List:
            tasks.print()
        case cf.Add != "":
            tasks.add(cf.Add)
        case cf.Edit != "":
            parts := strings.SplitN(cf.Edit, ":", 2)
            if len(parts) != 2 {
                fmt.Println("Edit should follow format id:new task")
                os.Exit(1)
            }
            index, err := strconv.Atoi(parts[0])
            if err != nil {
                fmt.Println("Error, inavlid index")
                os.Exit(1)
            }
            tasks.edit(index, parts[1])
        case cf.MultiAdd != "":
            parts := strings.Split(cf.MultiAdd, ",")
            tasks.multiadd(parts)
        case cf.Toggle != -1:
            tasks.toggle(cf.Toggle)
        case cf.Del != -1:
            tasks.delete(cf.Del)
        case cf.Clear:
            tasks.clear()
    }
}
