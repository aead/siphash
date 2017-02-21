// Copyright (c) 2016 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// Package siphash implements a hash / MAC function
// developed Jean-Philippe Aumasson and Daniel Bernstein.
// SipHash computes a 64-bit message authentication
// code from a variable-length message and a 128-bit secret
// key. It was designed to be efficient even for short inputs,
// with performance comparable to non-cryptographic hash
// functions. This package implements SipHash with the
// recommended parameters: c = 2 and d = 4.
package siphash // import "github.com/aead/siphash"

import (
	"crypto/subtle"
	"encoding/binary"
	"errors"
	"hash"
)

const (
	// KeySize is the size of the SipHash secret key in bytes.
	KeySize = 16
	// TagSize is the size of the SipHash authentication tag in bytes.
	TagSize = 8
)

// The four initialization constants
const (
	c0 = uint64(0x736f6d6570736575)
	c1 = uint64(0x646f72616e646f6d)
	c2 = uint64(0x6c7967656e657261)
	c3 = uint64(0x7465646279746573)
)

var errKeySize = errors.New("siphash: bad key length")

// Verify checks whether the given sum is equal to the
// computed checksum of msg. This function returns true
// if and only if the computed checksum is equal to the
// given sum.
func Verify(sum *[TagSize]byte, msg []byte, key *[16]byte) bool {
	var out [TagSize]byte
	Sum(&out, msg, key)
	return subtle.ConstantTimeCompare(sum[:], out[:]) == 1
}

// Sum generates an authenticator for msg with a 128 bit key
// and puts the 64 bit result into out.
func Sum(out *[TagSize]byte, msg []byte, key *[16]byte) {
	r := Sum64(msg, key)
	binary.LittleEndian.PutUint64(out[:], r)
}

// Sum64 generates and returns the 64 bit authenticator
// for msg with a 128 bit key.
func Sum64(msg []byte, key *[16]byte) uint64 {
	k0 := binary.LittleEndian.Uint64(key[:])
	k1 := binary.LittleEndian.Uint64(key[8:])

	var hVal [4]uint64
	hVal[0] = k0 ^ c0
	hVal[1] = k1 ^ c1
	hVal[2] = k0 ^ c2
	hVal[3] = k1 ^ c3

	n := len(msg)
	ctr := byte(n)

	if n >= TagSize {
		n &= (^(TagSize - 1))
		core(&hVal, msg[:n])
		msg = msg[n:]
	}

	var block [TagSize]byte
	copy(block[:], msg)
	block[7] = ctr

	return finalize(&hVal, &block)
}

// New returns a hash.Hash64 computing the SipHash checksum with a 128 bit key.
func New(key *[16]byte) hash.Hash64 {
	h := new(digest)
	h.key[0] = binary.LittleEndian.Uint64(key[:])
	h.key[1] = binary.LittleEndian.Uint64(key[8:])
	h.Reset()
	return h
}

// The siphash hash struct implementing hash.Hash
type digest struct {
	hVal  [4]uint64
	key   [2]uint64
	block [TagSize]byte
	off   int
	ctr   byte
}

func (d *digest) BlockSize() int { return TagSize }

func (d *digest) Size() int { return TagSize }

func (d *digest) Reset() {
	d.hVal[0] = d.key[0] ^ c0
	d.hVal[1] = d.key[1] ^ c1
	d.hVal[2] = d.key[0] ^ c2
	d.hVal[3] = d.key[1] ^ c3

	d.off = 0
	d.ctr = 0
}

func (d *digest) Write(p []byte) (int, error) {
	n := len(p)
	d.ctr += byte(n)

	if d.off > 0 {
		dif := TagSize - d.off
		if n > dif {
			d.off += copy(d.block[d.off:], p[:dif])
			p = p[dif:]
			core(&(d.hVal), d.block[:])
			d.off = 0
		} else {
			d.off += copy(d.block[d.off:], p)
			return n, nil
		}
	}

	if nn := len(p); nn >= TagSize {
		nn &= (^(TagSize - 1))
		core(&(d.hVal), p[:nn])
		p = p[nn:]
	}

	if len(p) > 0 {
		d.off = copy(d.block[:], p)
	}
	return n, nil
}

func (d *digest) Sum64() uint64 {
	hVal := d.hVal
	block := d.block
	for i := d.off; i < TagSize-1; i++ {
		block[i] = 0
	}
	block[7] = d.ctr
	return finalize(&hVal, &block)
}

func (d *digest) Sum(b []byte) []byte {
	r := d.Sum64()

	var out [TagSize]byte
	binary.LittleEndian.PutUint64(out[:], r)
	return append(b, out[:]...)
}
