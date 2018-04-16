package handlers

import (
  "database/sql"
  "net/http"
  "strconv"

  "github.com/labstack/echo"
)

//return arbitary JSON in our response
//a map with strings as keys and anything as values
//"interface" keyword represents anything from a primitive datatype to a user defined type or struct
type H map[string]interface{}

//GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    return c.JSON(http.StatusOK, "tasks")
  }
}

//PutTask endpoint
func PutTask(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    return c.JSON(http.StatusCreated, H{
      "created": 123,
      })
  }
}
//DeleteTask
func DeleteTask(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    return c.JSON(http.StatusOK, H{
      "deleted": id,
      })
  }
}