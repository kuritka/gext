// Package helps with generating random numbers and guids.
package utils

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// MustGenerateNewUUID returns a Random (Version 4) UUID or panics.
func MustGenerateNewUUID() string {
	return uuid.New().String()
}

// GenerateRandomUUID returns a Random (Version 4) UUID.
func GenerateRandomUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// GenerateRandomString returns randomly generated string of the
// given number of characters.
func GenerateRandomString(n int) string {
	if n < 0 {
		n = 0
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// ParseUUID returns parsed UUID as string.
func ParseUUID(s string) (string, error) {
	uuid, err := uuid.Parse(s)
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

// GenerateRandomNumber generates random number within low, high limit
func GenerateRandomNumber(min, max int) int {
	return min + rand.Intn(max-min)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}
