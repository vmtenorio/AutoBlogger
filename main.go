package main

import (
  "fmt"
  fu "autoBlogger/fileutils"
)


func main() {
  fmt.Println("Parsing files")
  var sett fu.Settings
  fu.ParseJson("settings.json", &sett)

  var creds fu.Creds
  fu.ParseJson("creds.json", &creds)

  email_pass := fu.DecodePassword(creds.User_password)

}
