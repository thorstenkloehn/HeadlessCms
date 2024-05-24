package main

import (
    "encoding/json"
    "fmt"
    "log"

    "model"
)

func main() {
    restaurant := model.Restaurant{
        ID:        1,
        Name:      "Restaurant Name",
        Kategorie: []string{"Italian", "Pizza"},
        Location:  "POINT(12.34 56.78)",
    }

    jsonData, err := json.Marshal(restaurant)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(jsonData))
}