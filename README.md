[![Godoc Reference](https://godoc.org/github.com/aead/siphash?status.svg)](https://godoc.org/github.com/aead/siphash)

## The SipHash pseudo-random function

SipHash is a family of pseudorandom functions (a.k.a. keyed hash functions) optimized for speed on short messages.  
SipHash computes a 64-bit message authentication code from a variable-length message and 128-bit secret key.
This implementation uses the recommended parameters c=2 and d=4.

### Installation
Install in your GOPATH: `go get -u github.com/aead/siphash`  
