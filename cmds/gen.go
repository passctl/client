package cmds

import (
	"math/rand"
	"strconv"

	"github.com/passctl/client/lib"
)

func GenPwd(s int) string {
  chars := [][]rune{
    []rune("0123456789"),
    []rune("abcdefghijklmnopqrstuvwxyz"),
    []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
    []rune("!\"#$%&\\'()*+,-./:;<=>?@[\\]^_`{|}~"),
  }

  c := 0
  res := make([]rune, s)
  rand.Shuffle(len(chars), func(i, j int) { 
    chars[i], chars[j] = chars[j], chars[i] })

  for i := 0; i < s; i++ {
    rand.Shuffle(len(chars), func(i, j int) { 
      chars[i], chars[j] = chars[j], chars[i] })

    if c > 3 {
      c = 0
    }

    res[i] += chars[c][rand.Intn(len(chars[c]))]
    c++
  }

  return string(res)
}

func Gen() bool {
START:
  l, err := strconv.Atoi(lib.Input("Enter a length for the password"))
  if err != nil {
    lib.Error("Please enter a valid number")
    goto START
  }

  if l <= 4 || l > 120 {
    lib.Error("Length must be larger then %s4%s and smaller than %s120%s",
      lib.COLOR_BOLD, lib.COLOR_RESET, lib.COLOR_BOLD, lib.COLOR_RESET)
    goto START
  }

  pwd := GenPwd(l) 
  lib.Info("Password: %s%s%s", 
    lib.COLOR_BOLD, pwd, lib.COLOR_RESET)
  return true
}
