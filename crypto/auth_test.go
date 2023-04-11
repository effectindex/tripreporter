// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package crypto

import (
	"bytes"
	"testing"
)

func TestRegisterOpen(t *testing.T) {
	auth, secret := Register([]byte("hunter2"))

	if secret_, ok := Open(auth, []byte("hunter2")); !ok {
		t.Error("Open failed")
	} else if !bytes.Equal(secret, secret_) {
		t.Error("Open didn't return the same secret")
	}

	if _, ok := Open(auth, []byte("*******")); ok {
		t.Error("Open didn't reject wrong password")
	}
}

func TestOpenInvalid(t *testing.T) {
	if _, ok := Open(make([]byte, 128), []byte("hunter2")); ok {
		t.Error("Open didn't reject wrong-length auth data")
	}
}

func TestSealInvalid(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("didn't panic on invalid secret length")
		}
	}()

	Seal(make([]byte, 64), []byte("hunter2"))
}
