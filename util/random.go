package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixMicro()))
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generate random owner name
func RandomOwner() string {
	return RandomString(6)
}

// Generate Random Money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Generate random Currency
func RandomCurrency() string {
	currecies := []string{"EUR", "USD", "KSH"}

	n := len(currecies)

	return currecies[rand.Intn(n)]
}
