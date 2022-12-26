package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomName generates a random name
func RandomName() string {
	return RandomString(6)
}

// RandomDate generates a random date
func RandomDate() time.Time {
	date := time.Unix(time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), 0).
		UTC().AddDate(0, 0, rand.Intn(1000))
	return date
}

// RandomEmail generates a random email
func RandomEmail() string {
	return RandomString(6) + "@gmail.com"
}

// RandomS
func RandomMobile() string {
	return strconv.FormatInt(RandomInt(1_000_000_000, 9_999_999_999), 10)
}

// RandomAddress generates a random address
func RandomAddress() string {
	return strconv.FormatInt(RandomInt(1, 200), 10) + " " + RandomString(8)
}

// RandomStudentID generates a random student id
func RandomStudentID() string {
	return "V2021" + strconv.FormatInt(RandomInt(10_000, 99_999), 10)
}
