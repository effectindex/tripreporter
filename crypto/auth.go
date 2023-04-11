// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package crypto

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/nacl/secretbox"
)

const (
	saltLen   = 16
	secretLen = 32
	authLen   = saltLen + secretLen + secretbox.Overhead
)

var zeroNonce [24]byte

func generate(buf []byte) {
	// if the system entropy source is broken, something is seriously wrong
	if _, err := rand.Read(buf); err != nil {
		panic(fmt.Errorf("rand.Read: %w", err)) // TODO: Use Zap panic (#28)
	}
}

func pkdf(password, salt []byte) (key [32]byte) {
	h := argon2.IDKey(password, salt, 1, 16*1024, 4, uint32(len(key)))
	copy(key[:], h)
	return
}

func Register(password []byte) (auth, secret []byte) {
	secret = make([]byte, secretLen)
	generate(secret)
	auth = Seal(secret, password)
	return
}

func Seal(secret, password []byte) (auth []byte) {
	if len(secret) != secretLen {
		panic("invalid secret length") // TODO: Use Zap panic (#28)
	}

	auth = make([]byte, saltLen, authLen)
	generate(auth)

	// the key is freshly generated from a random salt, and used only once (here), so we're fine using an all-zero nonce
	key := pkdf(password, auth)
	return secretbox.Seal(auth, secret, &zeroNonce, &key)
}

func Open(auth, password []byte) (secret []byte, ok bool) {
	if len(auth) != authLen {
		return
	}

	key := pkdf(password, auth[:saltLen])
	return secretbox.Open(nil, auth[saltLen:], &zeroNonce, &key)
}
