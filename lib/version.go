package lib

import "fmt"

var MAJOR = 3
var MINOR = 2

func GetVersion() string {
  return fmt.Sprintf("%d.%d", MAJOR, MINOR)
}
