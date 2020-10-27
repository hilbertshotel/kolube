package handlers

import (
  "os"
  "fmt"
  "time"
  "bufio"
)


// INITIALIZE IP AND DATABASE STRINGS FOR GLOBAL USE
var IP string
var connStr string

func init() {
  IP, connStr = ConnectionData()
}
func ShareData() (string, string) {
  return IP, connStr
}


// COMMUNICATION STRUCT
type Package struct {
  Id int `json:id`
  Url string `json:url`
  Description string `json:description`
}


// READ IP AND DATABASE STRINGS
func ConnectionData() (string, string) {
  file, err := os.Open("local/data.txt")
  Check(err)

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  var data []string
  for scanner.Scan() {
    data = append(data, scanner.Text())
  }

  file.Close()
  IP, connStr := data[0], data[1]
  return IP, connStr
}


// ERROR HANDLING AND LOGGING                 WRITE MUTEX HERE
func Check(err error) {
  if err != nil {

    // prepare log
    t := time.Now()
    log := fmt.Sprintf("%s | %s\n", t.Format("Mon Jan _2 15:04:05 2006"), err)

    // open file
    file, err := os.OpenFile("local/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
      fmt.Println(err)
    }
    defer file.Close()

    // write to file
    _, err = file.WriteString(log)
    if err != nil {
      fmt.Println(err)
    }
  }
}
