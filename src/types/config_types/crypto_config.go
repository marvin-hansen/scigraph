package config_types

import (
	"scigraph/src/utils/crypto_utils"
	"strconv"
)

const defaultKeyLength = 128

type CryptoConfig struct {
	seedRandom bool
	sessionKey bool
	masterKey  string
	keyHash    uint64
	configHash uint64
}

func NewCryptoConfig(masterKey string, seedRandom, generateSessionKey bool) *CryptoConfig {
	if generateSessionKey {
		sessionKey := crypto_utils.GenerateStringID(defaultKeyLength)
		encKey := crypto_utils.EncodeBase64(sessionKey)
		keyHash := crypto_utils.HashString(encKey)
		configHash := createConfigHash(encKey, seedRandom, generateSessionKey, keyHash)
		return &CryptoConfig{seedRandom: seedRandom, sessionKey: generateSessionKey, masterKey: encKey, keyHash: keyHash, configHash: configHash}
	} else {
		encKey := crypto_utils.EncodeBase64(masterKey)
		keyHash := crypto_utils.HashString(encKey)
		configHash := createConfigHash(encKey, seedRandom, generateSessionKey, keyHash)
		return &CryptoConfig{seedRandom: seedRandom, sessionKey: false, masterKey: encKey, keyHash: keyHash, configHash: configHash}
	}
}

// GetEncodedMasterKey returns the base64 encoded master key. Must be decoded before usage.
func (c *CryptoConfig) GetEncodedMasterKey() string {
	return c.masterKey
}

// SeedRandom returns the seed random config flag.
func (c *CryptoConfig) SeedRandom() bool {
	return c.seedRandom
}

// SessionKey returns true only if the key is a generated one-time key.
func (c *CryptoConfig) SessionKey() bool {
	return c.sessionKey
}

// Hash returns the 128bit hash of the encoded master key. Use as UUID
func (c *CryptoConfig) Hash() uint64 {
	return c.keyHash
}

// VerifyMasterHash returns true only if the keyHash parameter equals th verified stored hash of the encoded master key.
func (c *CryptoConfig) VerifyMasterHash(verifyHash uint64) bool {

	// re-creates  masterHash
	intHash := crypto_utils.HashString(c.masterKey)
	// compares the verifyHash parameter to the re-created intHash and the actual key hash to triple verify that everything is consistent.
	if verifyHash == c.keyHash && c.keyHash == intHash && verifyHash == intHash {
		return true
	} else {
		return false
	}
}

func (c *CryptoConfig) VerifyConfigHash() bool {

	h1 := createConfigHash(c.masterKey, c.seedRandom, c.sessionKey, c.keyHash)
	h2 := c.configHash

	if h1 == h2 {
		return true
	} else {
		return false
	}
}

func createConfigHash(encKey string, seedRandom, generateSessionKey bool, keyHash uint64) uint64 {

	s1 := encKey
	s2 := strconv.FormatBool(seedRandom)
	s3 := strconv.FormatBool(generateSessionKey)
	s4 := strconv.FormatUint(keyHash, 10)

	s5 := s1 + s2 + s3 + s4
	return crypto_utils.HashString(s5)
}
