package lib

import (
	"fmt"
	"strings"
	"syscall"

	"golang.org/x/term"
)

var COLOR_RED   = "\x1b[31m"
var COLOR_BOLD  = "\x1b[1m"
var COLOR_BLUE  = "\x1b[34m"
var COLOR_CYAN  = "\x1b[36m"
var COLOR_GREEN = "\x1b[32m"
var COLOR_RESET = "\x1b[0m"

func Info(s string, args ...interface{}) {
  f := fmt.Sprintf(s, args...)
  fmt.Printf("%s%s[*]%s %s%s\n", 
    COLOR_BLUE, COLOR_BOLD, COLOR_RESET, f, COLOR_RESET)
}

func Success(s string, args ...interface{}) {
  f := fmt.Sprintf(s, args...)
  fmt.Printf("%s%s[+]%s %s%s\n", 
    COLOR_GREEN, COLOR_BOLD, COLOR_RESET, f, COLOR_RESET)
}

func Error(s string, args ...interface{}) {
  f := fmt.Sprintf(s, args...)
  fmt.Printf("%s%s[-]%s %s%s\n", 
    COLOR_RED, COLOR_BOLD, COLOR_RESET, f, COLOR_RESET)
}

func ErrorMsg(s string, e error) {
  fmt.Printf("%s%s[-]%s %s: %s%s%s\n", 
    COLOR_RED, COLOR_BOLD, COLOR_RESET, s, COLOR_BOLD, e.Error(), COLOR_RESET)
}

func InfoVar(k string, v string) {
  Info("%s => %s%s%s", k, COLOR_BOLD, v, COLOR_RESET)
}

func Input(s string, args ...interface{}) string {
  var res string
  f := fmt.Sprintf(s, args...)
  fmt.Printf("%s%s[?]%s %s%s: ", 
    COLOR_CYAN, COLOR_BOLD, COLOR_RESET, f, COLOR_RESET)
  fmt.Scan(&res)
  return strings.TrimSpace(res)
}

func Passwd(s string, args ...interface{}) string {
  f := fmt.Sprintf(s, args...)
  fmt.Printf("%s%s[?]%s %s%s: ", 
    COLOR_CYAN, COLOR_BOLD, COLOR_RESET, f, COLOR_RESET)

  rawpwd, err := term.ReadPassword(int(syscall.Stdin))
  fmt.Println()

  if err != nil {
    return ""
  }

  return strings.TrimSpace(string(rawpwd))
}


