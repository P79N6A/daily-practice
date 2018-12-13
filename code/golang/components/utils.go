package components

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"testing"
)

// HandleError (err error)
func HandleError(err error) {
	if err != nil {
		panic(err)
		// log.Printf("%v\n", err)
	}
}

// HandleTesting (err error)
func HandleTesting(t *testing.T, err error, msg ...interface{}) {
	if err != nil {
		t.Error(err)
	} else {
		if len(msg) == 0 {
			log.Println("[OK]")
		} else {
			log.Println(msg...)
		}
	}
}

// HashPassword ...
func HashPassword(passwd string) string {
	h := md5.New()
	io.WriteString(h, passwd)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// HmacSha1 ...
func HmacSha1(plaintext string) string {
	//hmac ,use sha1
	key := []byte(secretKey)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(plaintext))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

// StartProfiling ...
func StartProfile(fpath string) {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
}
