package lib

import "fmt"

var MAJOR = 3
var MINOR = 1

func GetVersion() string {
  return fmt.Sprintf("%d.%d", MAJOR, MINOR)
}
