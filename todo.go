package main

import (
  //"net/http"
  "github.com/labstack/echo"
  // "github.com/labstack/echo/engine/standard"


)


func main(){
  //A new instance of Echo

  e := echo.New()

  e.GET("/tasks", func(c echo.Context) error { return c.JSON(200, "GET Tasks") })
  e.PUT("/tasks", func(c echo.Context) error { return c.JSON(200, "PUT Tasks") })
  e.DELETE("/tasks/:id", func(c echo.Context) error {return c.JSON(200, "DELETE Task " + c.Param("id"))})

  //Start web server

  // e.Run(standard.New(":8000"))
  e.Logger.Fatal(e.Start(":8000"))
}