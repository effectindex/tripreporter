// SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package util

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

// GenerateSaltedPasswordHash generates a 32-bit hash from a password and salt with the RFC-recommended settings
func GenerateSaltedPasswordHash(p []byte, s []byte) []byte {
	return argon2.IDKey(p, s, 1, 64*1024, 4, 32)
}

// GenerateSalt will generate a salt with a set length of tLen and will try to fit rLen into tLen with s as padding for leftover space.
// Consider rLen to be the minimum random bytes you will get, with s being added if there is space, and the space until tLen being filled with random, if empty
func GenerateSalt(rLen, tLen int, s string) ([]byte, error) {
	if rLen > tLen {
		rLen = tLen
	}

	// Generate the initial random base
	if b, err := GenerateRandomBytes(rLen); err != nil {
		return nil, err
	} else {
		// Append padding string bytes
		b = append(b, []byte(s)...)

		// If too big, trim
		if len(b) > tLen {
			return b[:tLen], nil
		}

		// If too short, fill the rest with random
		if len(b) < tLen {
			if f, err := GenerateRandomBytes(tLen - len(b)); err != nil {
				return nil, err
			} else {
				b = append(b, f...)
			}
		}

		return b, nil
	}
}

// GenerateRandomBytes will generate n bytes from the crypto/rand lib
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
