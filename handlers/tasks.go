package handlers

import (
  "database/sql"
  "net/http"
  "strconv"

  "go-echo-vue/models"

  "github.com/labstack/echo"
)

//return arbitary JSON in our response
//a map with strings as keys and anything as values
//"interface" keyword represents anything from a primitive datatype to a user defined type or struct
type H map[string]interface{}

//GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    //return c.JSON(http.StatusOK, "tasks")
    //Fetch tasks using model
    return c.JSON(http.StatusOK, models.GetTasks(db))
  }
}

//PutTask endpoint
func PutTask(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    // return c.JSON(http.StatusCreated, H{
    //   "created": 123,
    //   })
    //Instantiate new task
    var task models.Task
    //Map incoming JSON body to the new task
    c.Bind(&task)
    //c.Bind takes a JSON formatted body sent in a PUT request and maps it to a Task struct.
    //The Task struct will be defined in our models package.
    //Add task using model
    id, err := models.PutTask(db, task.Name)
    //Return JSON response if successfull creation of task
    if err == nil {
      return c.JSON(http.StatusCreated, H{
        "created": id,
        })
      //Handle errors
    } else {
      return err
    }

  }
}
//DeleteTask
func DeleteTask(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    // return c.JSON(http.StatusOK, H{
    //   "deleted": id,
    //   })
    //Use model to delete a task
    _, err := models.DeleteTask(db, id)
    //Return JSON response if successful deletion of record
    if err == nil {
      return c.JSON(http.StatusOK, H{
        "deleted": id,
        })
      //Handle errors
    } else {
      return err
    }
  }
}