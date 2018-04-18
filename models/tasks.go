package models

import (
  "database/sql"

  _ "github.com/mattn/go-sqlite3"

)

// Task is a struct containing Task data
/*
Go lets you add meta data to variables using backticks.
We define what we want each field to look like once it is converted to JSON.
This allows the "c.Bind" function in our handlers to know where to map JSON data when populating a new Task.
*/

type Task struct {
  ID int      `json:"id"`
  Name string `json:"name"`
}
 /*
Second type is a collection of Task items
Return all tasks in the database
 */

//TaskCollection is collection of Tasks

type TaskCollection struct {
  Tasks []Task `json:"items"`
}

//GetTasks selects all tasks from the database, places them in a new collection, and returns them

func GetTasks(db *sql.DB) TaskCollection {
  sql := "SELECT * FROM tasks"
  row, err := db.Query(sql)
  //Exit if SQL doesn't work
  if err != nil {
    panic(err)
  }
  //Cleanup when program exits
  defer rows.Close()

  result := TaskCollection{}
  for rows.Next() {
    task := Task{}
    err2 := rows.Scan(&task.ID, &task.Name)
    //Exit if error
    if err2 != nil {
      panic(err2)
    }

    result.Tasks = append(result.Tasks, task)
  }
  return result
}

//PutTask inserts a new task into the database and returns either the new id on success or a panic on failure

func PutTask(db *sql.DB, name string) (int64, error) {
  sql := "INSERT INTO tasks(name) VALUES(?)"

  //Create a prepared SQL statement
  stmt, err := db.Prepare(sql)
  //Exit on error
  if err != nil {
    panic(err)
  }

  //Cleanup after program exits
  defer stmt.Close()

  //Replace '?' in SQL statement with 'name'
  result, err2 := stmt.Exec(name)
  //Exit on error
  if err2 != nil {
    panic(err2)
  }

  return result.LastInsertId()
}