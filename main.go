package main

import (
  "gopkg.in/hlandau/passlib.v1"
  "flag"
  "fmt"
  "os"
  "strings"
  "io/ioutil"
  )

type someuser struct {
  user string
  hash string
}

var user string
var password string
var shadow string
var user_hash string

func init() {
  flag.StringVar(&user, "user", "", "user to verify")
  flag.StringVar(&password, "password", "", "password to verify")
  flag.StringVar(&shadow, "shadow", "/etc/shadow", "shadow file")
  flag.Parse()
}

func main() {

  userdata := getUserHash(user)
  user_hash := userdata.hash
  // fmt.Println("user: %s", user)
  // fmt.Println("password: %s", password)
  _, err := passlib.Verify(password, user_hash)
  if err != nil {
    fmt.Println("invalid")
    os.Exit(1)
  } else {
    fmt.Println("valid")
    os.Exit(0)
  }
}

func getUserHash(user string) someuser {
  buf, err := ioutil.ReadFile(shadow);
  if err != nil {
    panic(err)
  }
  s := string(buf)
    users := strings.Split(s, "\n")

  var userdata []string

  for _, element := range users {
    userdata = strings.Split(element, ":")

    if userdata[0] == user {
      return someuser{user:userdata[0], hash:userdata[1]}
    }
  }

  return someuser{user: "not found"}
}
