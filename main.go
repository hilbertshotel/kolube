package main

import (
  "fmt"
  "net/http"
  "kolube/handlers"
)

func index(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "./templates/index.html")
}

func main() {
  // get IP from file
  IP, _ := handlers.ShareData()

  // serve css and js files
  http.Handle("/static/",
    http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

  // routes
  http.HandleFunc("/", index)
  http.HandleFunc("/loadBookmarks", handlers.LoadBookmarks)
  http.HandleFunc("/addBookmark", handlers.AddBookmark)
  http.HandleFunc("/delBookmark", handlers.DelBookmark)

  // listen
  fmt.Println("Now serving @", IP)
  http.ListenAndServe(IP, nil)
}
