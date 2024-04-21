package util

import (
	"math/rand"
	"time"
)

func init() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// RandomOwner generates a random owner email
func RandomOwnerEmail() string {
	return RandomString(6) + "@gmail.com"
}

func RandomRole() string {
	roles := []string{"user", "guide", "leadguide", "admin"}
	n := len(roles)
	return roles[rand.Intn(n)]
}
