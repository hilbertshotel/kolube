package handlers

import (
  "net/http"
  "database/sql"
  "encoding/json"
  _ "github.com/lib/pq"
)

func LoadBookmarks(w http.ResponseWriter, r *http.Request) {
  // connect to db
  db, err := sql.Open("postgres", ConnStr)
  Check(err)
  defer db.Close()

  // query db
  rows, err := db.Query("SELECT * FROM bookmarks")
  Check(err)
  defer rows.Close()

  // load query into package
  var response []BookmarkPackage
  for rows.Next() {
    bookmark := BookmarkPackage{}
    err = rows.Scan(&bookmark.Id, &bookmark.Url, &bookmark.Description)
    Check(err)
    response = append(response, bookmark)
  }

  // send response package
  w.Header().Set("content-type", "application/json")
  output, err := json.Marshal(response)
  Check(err)
  w.Write(output)
}
