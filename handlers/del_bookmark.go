package handlers

import (
  "net/http"
  "io/ioutil"
  "database/sql"
  "encoding/json"
  _ "github.com/lib/pq"
)

func DelBookmark(w http.ResponseWriter, r *http.Request) {
  // pick up request data
  var idToDelete int
  request, err := ioutil.ReadAll(r.Body)
  Check(err)
  json.Unmarshal(request, &idToDelete)

  // get connection string for database from file
  _, connStr := ShareData()

  // connect to database
  db, err := sql.Open("postgres", connStr)
  Check(err)
  defer db.Close()

  // del bookmark from database
  _, err = db.Exec("DELETE FROM bookmarks WHERE id = $1", idToDelete)
  Check(err)
}
