package handlers

import (
  "net/http"
  "io/ioutil"
  "database/sql"
  "encoding/json"
  _ "github.com/lib/pq"
)

func AddBookmark(w http.ResponseWriter, r *http.Request) {
  // pick up request data
  newBookmark := Package{}
  request, err := ioutil.ReadAll(r.Body)
  Check(err)
  json.Unmarshal(request, &newBookmark)

  // connect to database
  db, err := sql.Open("postgres", ConnStr)
  Check(err)
  defer db.Close()

  // add new bookmark to database
  _, err = db.Exec(`INSERT INTO bookmarks (url, description)
  VALUES ($1, $2)`, newBookmark.Url, newBookmark.Description)
  Check(err)
}
