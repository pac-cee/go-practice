package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "os"
)

type Record struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    // Example: Write CSV
    file, _ := os.Create("data.csv")
    writer := csv.NewWriter(file)
    writer.Write([]string{"Name", "Age"})
