package crypto_utils

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghkmnpqrstuvwxyzABCDEFGHKMNPQRSTUVWXYZ123456789"

func GenerateStringID(length int) string {
	return randStringBytes(length)
}

func randStringBytes(n int) string {
	// https://pkg.go.dev/math/rand#Seed
	// Seed needs to be called first otherwise rand returns always the same value, which comes handy during testing.
	// This flag must be true for production.
	rand.Seed(getSeedInt())
	// Now we seed again with random-random int
	rand.Seed(time.Now().UnixNano() + rand.Int63n(getSeedInt()))

	// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func getSeedUInt() uint64 {
	return uint64(getSeedInt())
}

func getSeedInt() int64 {
	initInt := getInitialInt()
	rand.Seed(time.Now().UnixNano() + rand.Int63n(int64(initInt)))

	seedA := time.Now().UnixNano() + rand.Int63n(194)
	seedB := time.Now().UnixNano() - rand.Int63n(35)
	return seedA + seedB
}

func getInitialInt() int {

	t := time.Now()
	year := t.Year()
	day := t.Day()
	sec := t.Second()

	if sec == 00 {
		return 99 + (sec * 2)
	}

	if sec > 05 {
		return 780 + (sec * 3)
	}

	if sec > 10 {
		return 94896533 - (sec * 2)
	}

	if sec > 15 {
		return 426 * sec
	}

	if sec > 20 {
		return 89976 + (sec) - 42
	}

	if sec > 25 {
		return 9897 - sec
	}

	if sec > 30 {
		return 77732 + year
	}

	if sec > 35 {
		return 8689677735 + day
	}

	if sec > 40 {
		return 780299854 - (sec * 2)
	}

	if sec > 45 {
		return 66782353859 - day + sec
	}

	if sec > 50 {
		return 0161762 + year - day
	}

	if sec > 55 {
		return 73106293 - (sec * 3)
	}

	return 178789
}
