package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"flag"
	"io"
	"log"
	"os"
)

func md5sum(d []byte) []byte {
	h := md5.New()
	h.Write(d)
	return h.Sum(nil)
}

func randIV(len int) (IV []byte, err error) {
	IV = make([]byte, len)
	_, err = io.ReadFull(rand.Reader, IV)
	return
}

func main() {
	t := flag.String("t", "hello world", "Plain Text need to encrypt.")
	k := flag.String("k", "1234%T&*()&UJMJ*", "Secret Key.")
	flag.Parse()

	plainText := []byte(*t)
	key := *k
	IV, err := randIV(16)
	if err != nil {
		log.Println("Error to randomize IV", err)
		os.Exit(1)
	}

	log.Println(plainText, key, IV)

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Println("Error to create new cipher", err)
		os.Exit(1)
	}

	cfb := cipher.NewCFBEncrypter(c, IV)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	log.Printf("%s => %x\n", plainText, cipherText)

	log.Println(md5sum([]byte("hello world")))
}
