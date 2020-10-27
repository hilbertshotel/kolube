package handlers

import "fmt"

type Package struct {
  Id int `json:id`
  Url string `json:url`
  Description string `json:description`
}

func Check(err error) {
  if err != nil {
    fmt.Println(err)
  }
}
