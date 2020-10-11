package fileutils

import (
  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
  b64 "encoding/base64"
)

type Settings struct {
  SMTP_server string
  SMTP_port string
  EMAIL_template string
  POST_template string
  Mail_list string
}

type Creds struct {
  User_email string
  User_password string
}

func ParseJson(filename string, content interface{}) {
  json_file, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }

  byte_Content, _ := ioutil.ReadAll(json_file) //type []uint8
  json.Unmarshal([]byte(byte_Content), content)

  json_file.Close()
}

func DecodePassword (enc_password string) string {
  comp_password := enc_password + "=="
  dec_password, err := b64.StdEncoding.DecodeString(comp_password)
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
  return string(dec_password)
}
