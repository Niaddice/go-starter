package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var src = rand.NewSource(time.Now().UnixNano())

const (
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func RandSalt(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return string(b)
}

func GenerateHashPassword(pwd string, salt string) string {
	bytes := []byte(pwd)
	sum := []byte(string(sha1.New().Sum(bytes)) + salt)
	return fmt.Sprintf("%x", md5.Sum(sum))
}

func TestApproach6(t *testing.T) {
	fmt.Println(RandSalt(10))
}

func BenchmarkApproach6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandSalt(10)
	}
}
