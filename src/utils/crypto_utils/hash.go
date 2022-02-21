package crypto_utils

import "github.com/zeebo/xxh3"

// Port of the xxh3 library to Go. Among the fastest hashing algo around.
//https://github.com/zeebo/xxh3

// HashString  returns the hash of the string slice. Always returns the same hash for the same input.
// Use this for verification.
func HashString(s string) uint64 {
	return xxh3.HashString(s)
}

// HashBytes  returns the hash of the byte slice. Always returns the same hash for the same input.
// Use this for verification.
func HashBytes(b []byte) uint64 {
	return xxh3.Hash(b)
}

// HashRandString returns a 128-bit random hash of the string slice using a new random seed for each call.
// Use this as random UUID!
// Important: This function returns a different hash for identical strings on each call because of the random seed and thus the has cannot be used for verification.
func HashRandString(s string) uint64 {
	newSeed := getSeedUInt() // The "unpredictability" from changing the seed should come from the seed itself.
	return xxh3.HashStringSeed(s, newSeed)
}

// HashRandBytes returns a 128-bit random hash of the byte slice using a new random seed for each call.
// Use this as random UUID!
// Important: This function returns a different hash for identical inputs on each call because of the random seed and thus the hash cannot be used for verification.
func HashRandBytes(b []byte) uint64 {
	newSeed := getSeedUInt()
	return xxh3.HashSeed(b, newSeed)
}
