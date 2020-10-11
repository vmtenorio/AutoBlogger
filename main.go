package main

import (
  "fmt"
  fu "autoBlogger/fileutils"
  mu "autoBlogger/mailutils"
  "os"
)


func main() {
  fmt.Println("Parsing files")
  var sett fu.Settings
  fu.ParseJson("settings_test.json", &sett)

  var creds fu.Creds
  fu.ParseJson("creds.json", &creds)

  emailPass := fu.DecodePassword(creds.UserPassword)

  // Publish post through Blogger API

  // Send email to email list
  // Contents file from argv
  contents_file := os.Args[1]
  mailContent := fu.BuildFromTemplate(sett.MailTemplate, contents_file)

  mu.SendEmail(mailContent, creds.UserEmail, emailPass, sett)
/*
  postContent := fu.BuildFromTemplate(sett.PostTemplate, contents_file)
  fmt.Println(postContent)
  */
}
