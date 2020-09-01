package main

import (
  _ "github.com/lib/pq"
  "database/sql"
  "fmt"
  "log"
  "strings"
)

const (
    hostname = "localhost"
    host_port = 5432
    username = "yollotl"
    password = "yollotl" 
    database_name = "catalog"
)
type movie struct {
    name string
    year int
    time int
    genre string
    original bool
}

func connect(query string , PWD string) string {
    pg_con_string := fmt.Sprintf("port=%d host=%s user=%s "+"password=%s dbname=%s sslmode=disable",host_port, hostname, username, PWD, database_name)
    db, err :=  sql.Open("postgres",pg_con_string) 
    if err != nil {
        log.Fatal(err)
    }

    query = strings.ToLower(query)
    query = strings.ReplaceAll(query,"'","''")
    query = fmt.Sprintf("SELECT * FROM netflix_en WHERE name = '%s'",query)
    rows, err := db.Query(query)
    if err != nil{
        log.Fatal(err)
    }
    defer rows.Close()
    val := make([] *movie,0)  

    for rows.Next() { 
        mov := new(movie)
        err = rows.Scan(&mov.name, &mov.year ,&mov.time ,&mov.genre ,&mov.original)
        if err != nil {
            log.Fatal(err)
        }
        val = append(val,mov)
    }
    err = rows.Err()
    if err != nil {
        panic(err)
    }

    for _, bk := range val {
        return fmt.Sprintf("%s %d %d %s %d \n",bk.name, bk.year,bk.time, bk.genre, bk.original)
    }
    if rows.Next() == false{
        return ("0 rows")
    }
    return "o"
}
