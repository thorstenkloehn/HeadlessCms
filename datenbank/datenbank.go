package datenbank

import (
    "encoding/json"
    "fmt"
    "os"
)

// Database repräsentiert die Konfiguration einer Datenbank.
type Database struct {
    Host     *string `json:"host"`
    Port     *int    `json:"port"`
    Name     *string `json:"name"`
    Username *string `json:"username"`
    Password *string `json:"password"`
}

// Config enthält die Konfiguration für die Anwendung.
type Config struct {
    Database *Database `json:"database"`
}

// Dbinfo liest die Konfigurationsdatei und gibt eine Datenbankverbindungszeichenkette zurück.
// Es gibt eine leere Zeichenkette zurück, wenn es einen Fehler beim Lesen der Konfigurationsdatei gibt oder wenn einer der erforderlichen Werte fehlt.
func Dbinfo() string {
    file, err := os.Open("config.json")
    if err != nil {
        fmt.Println(err)
        return ""
    }
    defer file.Close()

    config := &Config{}
    decoder := json.NewDecoder(file)
    err = decoder.Decode(config)
    if err != nil {
        fmt.Println(err)
        return ""
    }

    if config.Database.Host != nil && config.Database.Port != nil && config.Database.Username != nil && config.Database.Password != nil && config.Database.Name != nil {
        dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			*config.Database.Host, *config.Database.Port, *config.Database.Username, *config.Database.Password, *config.Database.Name)
        // Setzen Sie config auf nil, um den Speicher freizugeben
        config = nil
    
        return dbInfo
    } else {
        return ""
    }
}