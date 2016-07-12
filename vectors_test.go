// Copyright (c) 2016 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package siphash

import (
	"encoding/hex"
	"testing"
)

func fromHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

var vectors = []struct {
	key, msg string
	hash     uint64
}{
	// Test vector from https://131002.net/siphash/siphash.pdf
	{
		key:  "000102030405060708090a0b0c0d0e0f",
		msg:  "000102030405060708090a0b0c0d0e",
		hash: uint64(0xa129ca6149be45e5),
	},
}

// Tests the SipHash implementation
func TestVectors(t *testing.T) {
	for i, v := range vectors {
		var key [16]byte
		copy(key[:], fromHex(v.key))

		msg := fromHex(v.msg)

		h := New(&key)
		_, err := h.Write(msg)
		if err != nil {
			t.Fatalf("Test vector %d: Spihash write failed: %s", i, err)
		}

		sum := h.Sum64()
		if sum != v.hash {
			t.Fatalf("Test vector %d: Hash values don't match - found %x expected %x", i, sum, v.hash)
		}
		sum = Sum64(msg, &key)
		if err != nil {
			t.Fatalf("Test vector %d: Failed to calculate MAC: %s", i, err)
		}
		if sum != v.hash {
			t.Fatalf("Hash values don't match - found %x expected %x", sum, v.hash)
		}
	}
}
