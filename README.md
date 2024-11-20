# Golang CLI todo-list 

## Setup
    Compile the binary by running "go build". Then add the "todo" binary to your local bin with "sudo mv todo /usr/local/bin", to make it a global command.
    
## Usage
    todo [flag] [arg]

## Flags
    -a [str arg]: add a single task
    -m [str arg]: add multiple tasks, seperated by a ,
    -e [str arg]: edit a task, arg given in format "index:update"
    -d [int arg]: delete a task 
    -t [int arg]: toggle a task as complete/uncomplete
    -l: list the tasks
    -c: clear all tasks
    
