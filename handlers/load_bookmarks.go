package handlers

import (
  "net/http"
  "database/sql"
  "encoding/json"
  _ "github.com/lib/pq"
)

func LoadBookmarks(w http.ResponseWriter, r *http.Request) {
  // get connection string for database from file
  _, connStr := ShareData()

  // connect to db
  db, err := sql.Open("postgres", connStr)
  Check(err)
  defer db.Close()

  // query db
  rows, err := db.Query("SELECT * FROM bookmarks")
  Check(err)
  defer rows.Close()

  // load query into package
  var response []Package
  for rows.Next() {
    bookmark := Package{}
    err = rows.Scan(&bookmark.Id, &bookmark.Url, &bookmark.Description)
    Check(err)
    response = append(response, bookmark)
  }

  // send response package
  w.Header().Set("content-type", "application/json")
  output, _ := json.Marshal(response)
  w.Write(output)
}
