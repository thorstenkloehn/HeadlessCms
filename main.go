package main

import (
    "fmt"

)

// Die Hauptfunktion des Programms.
func main() {
    // Rufen Sie die Dbinfo-Funktion aus dem datenbank-Paket auf, um eine Datenbankverbindungszeichenkette zu erhalten.
    start := dbinfo()

    // Drucken Sie die Datenbankverbindungszeichenkette auf die Konsole.
    fmt.Println(start)
}