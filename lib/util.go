package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	rand2 "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

func JoinURL(src, path string) string {
  if src[len(src)-1] == '/' {
    return fmt.Sprintf("%s%s", src, path)
  }
  return fmt.Sprintf("%s/%s", src, path)
}

func GetSHA256(str string) string {
  sum := sha256.Sum256([]byte(str))
  return fmt.Sprintf("%x", sum)
}

func GetMD5(str string) string {
  sum := md5.Sum([]byte(str))
  return fmt.Sprintf("%x", sum)
}

func MkRandom(size int) string {
  ret := make([]rune, size)
  for i := range ret {
    ret[i] = chars[rand.Intn(len(chars))]
  }
  return string(ret)
}

func Mkdirp(path string) error {
  err := os.Mkdir(path, os.ModePerm)
  if err != nil && !os.IsExist(err){
    return err
  }

  if err != nil && os.IsExist(err){
    return nil 
  }

  return err
}

func Encrypt(key, plain string) (string, error) {
  aes, err := aes.NewCipher([]byte(key))
  if err != nil {
    return "", err
  }

  gcm, err := cipher.NewGCM(aes)
  if err != nil {
    return "", err
  }

  nonce := make([]byte, gcm.NonceSize())
  _, err = rand2.Read(nonce)
  if err != nil {
    return "", err
  }

  cipher := gcm.Seal(nonce, nonce, []byte(plain), nil)
  return base64.StdEncoding.EncodeToString(cipher), nil
}


func Decrypt(key, enc string) (string, error) {
  dec, err := base64.StdEncoding.DecodeString(enc)
  if err != nil {
    return "", err
  }

  aes, err := aes.NewCipher([]byte(key))
  if err != nil {
    return "", err 
  }

  gcm, err := cipher.NewGCM(aes)
  if err != nil {
    return "", err
  }

  ns := gcm.NonceSize()
  nonce, text := dec[:ns], dec[ns:]

  plain, err := gcm.Open(nil, []byte(nonce), []byte(text), nil)
  if err != nil {
    return "", err
  }

  return string(plain), nil
}
