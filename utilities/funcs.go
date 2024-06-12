package utilities

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateVerificationCode() int {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	randomInt := rand.Intn(max-min) + min
	return randomInt
}

func ConvertIntToString(number int) string {
	return strconv.Itoa(number)
}

func GenerateVerificationGracePeriod() time.Time {
	return time.Now().Add(time.Minute * 15)
}
