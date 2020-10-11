package fileutils

import (
  "os"
  "encoding/json"
  "io/ioutil"
  b64 "encoding/base64"
  "github.com/gomarkdown/markdown"
  "log"
  "strings"
)

type Settings struct {
  SmtpServer string
  SmtpPort string
  MailTemplate string
  PostTemplate string
  MailList string
}

type Creds struct {
  UserEmail string
  UserPassword string
}

func ParseJson(filename string, content interface{}) {
  json_file, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }

  byte_Content, _ := ioutil.ReadAll(json_file) //type []uint8
  json.Unmarshal([]byte(byte_Content), content)

  json_file.Close()
}

func DecodePassword (enc_password string) string {
  comp_password := enc_password + "=="
  dec_password, err := b64.StdEncoding.DecodeString(comp_password)
  if err != nil {
    log.Fatal(err)
  }
  return string(dec_password)
}

const PLACEHOLDER_TEXT = "{% TEXT %}"

func BuildFromTemplate (templateFilename string, contentsFile string) string {
  template, err := ioutil.ReadFile(templateFilename)
  if err != nil {
    log.Fatal(err)
  }

  content, err := ioutil.ReadFile(contentsFile)
  if err != nil {
    log.Fatal(err)
  }

  contentHTML := markdown.ToHTML([]byte(content), nil, nil)

  return strings.Replace(string(template), PLACEHOLDER_TEXT, string(contentHTML), 1)
}

