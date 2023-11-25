package util

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxy"

func init() {
	rand.Int63n(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
func RandomOwner() string {
	return RandomString(8)
}
func RandomMoney() int64 {
	return RandomInt(500, 10000)
}

func RandomCurrency() string {
	currencies := []string{
		"INR", "USD", "EUR",
	}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func ErrorLogFatal(err error, msg ...any) {

	if err != nil {
		log.Fatal(err)
	}
}
