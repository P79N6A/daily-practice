package main

import (
    "fmt"
    "encoding/json"
)

type Movie struct {
    Title   string
    Year    int     `json:"released"`
    Color   bool    `json:"color,omitempty"`
    Actors  []string
}

var movies = []Movie{
    {Title: "Casablanca", Year: 1942, Color: false,
        Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
    {Title: "Cool Hand Luke", Year: 1967, Color: true,
        Actors: []string{"Paul Newman"}},
    {Title: "Bullitt", Year: 1968, Color: true,
        Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
    jsonString, err := json.MarshalIndent(movies, "", "  ")
    if err != nil {
        fmt.Printf("%v\n", err)
        return
    }
    fmt.Printf("%s\n", jsonString)

    var objects []Movie
    if err := json.Unmarshal(jsonString, &objects); err != nil {
        fmt.Printf("%v\n", err)
        return
    }
    fmt.Printf("objects: %v\n", objects)
}
