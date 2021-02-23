package main

import (
	"fmt"
    "database/sql"
    _ "github.com/lib/pq"
)

func main() {
    connStr := "host=localhost port=5433 user=zapgis dbname=geodados"
	connStr = "postgres://zapgis:.senha.@localhost:5433/geodados"
    db, err := sql.Open("postgres", connStr)

    if err != nil {
        fmt.Println("Fail to connect to the database with parameters", connStr)
    }

    //muni := "3550308"
	//max := 100

    fmt.Println("Querying database...")
    rows, err := db.Query("SELECT setor_id, lograd_nome, numero FROM censo2010.cnefe_lotes WHERE setor_id like '3550308%' LIMIT 10")

	if err != nil {
        fmt.Println("Error executing query", err)
    }
    
    for rows.Next() {
        var (
            setor_id string
            lograd_nome string
            numero int
        )
        if err := rows.Scan(&setor_id, &lograd_nome, &numero); err != nil {
            fmt.Println("Fail to recouvery data", err)
        }
        fmt.Printf("\nsetor_id %s, logradouro %s, %d", setor_id, lograd_nome, numero)
	}
    fmt.Println()

}
