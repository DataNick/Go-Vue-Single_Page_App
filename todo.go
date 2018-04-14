package main

import (
  "database/sql"
  //"net/http"
  "github.com/labstack/echo"
  _"github.com/mattn/go-sqlite3"
)


func main(){

  db := initDB("storage.db")
  migrate(db)

  //A new instance of Echo
  e := echo.New()

  e.GET("/tasks", func(c echo.Context) error { return c.JSON(200, "GET Tasks") })
  e.PUT("/tasks", func(c echo.Context) error { return c.JSON(200, "PUT Tasks") })
  e.DELETE("/tasks/:id", func(c echo.Context) error {return c.JSON(200, "DELETE Task " + c.Param("id"))})

  //Start web server

  e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
  db, err := sql.Open("sqlite3", filepath)

  // Check for db errors

  if err != nil {
    panic(err)
  }

  //  No errors and no db connection --  exit
  if db == nil {
    panic("db nil")
  }

  return db
}

func migrate(db *sql.DB) {
  sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL
    );
    `
  _, err := db.Exec(sql)

  //Exit if SQL statement above errors out
  if err != nil {
    panic(err)
  }
}