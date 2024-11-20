package main

import (
    "os"
    "path/filepath"
    "encoding/json"
)

type storage[T any] struct{
    fileName string
}

func NewStorage[T any](fileName string) *storage[T]{
    configDir, err := os.UserHomeDir() // Use UserConfigDir for user-specific config
    if err != nil {
        panic(err)
    }

    fullPath := filepath.Join(configDir, fileName) // e.g., ~/.config/myapp/todos.json

    err = os.MkdirAll(filepath.Dir(fullPath), 0755) // Ensure the directory exists
    if err != nil {
        panic(err)
    }

    return &storage[T]{fileName: fullPath}
}

func (s *storage[T]) Save(data T) error {
    fileData, err := json.MarshalIndent(data, "", "    ")

    if err != nil {
        return err
    }
    
    return os.WriteFile(s.fileName, fileData, 0644)
}

func (s *storage[T]) Load(data *T) error {
    fileData, err := os.ReadFile(s.fileName)

    if err != nil {
        return err
    }

    return json.Unmarshal(fileData, data)
}
