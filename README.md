[![Godoc Reference](https://godoc.org/github.com/aead/siphash?status.svg)](https://godoc.org/github.com/aead/siphash)
[![Build Status](https://travis-ci.org/aead/siphash.svg?branch=master)](https://travis-ci.org/aead/siphash)

## The SipHash pseudo-random function

SipHash is a family of pseudo-random functions (a.k.a. keyed hash functions) optimized for speed on short messages.  
SipHash computes a 64-bit or 128 bit message authentication code from a variable-length message and 128-bit secret key.
This implementation uses the recommended parameters c=2 and d=4.

### Installation
Install in your GOPATH: `go get -u github.com/aead/siphash`  

### Performance
**AMD64**  
Hardware: Intel i7-6500U 2.50GHz x 2  
System: Linux Ubuntu 16.04 - kernel: 4.4.0-67-generic  
Go version: 1.8.0  
```
name         speed           cpb
Write_8-4     688MB/s ± 0%   3,47
Write_1K-4   2.09GB/s ± 5%   1,11
Sum64_8-4     244MB/s ± 1%   9,77
Sum64_1K-4   2.06GB/s ± 0%   1,13
Sum128_8-4    189MB/s ± 0%  12,62
Sum128_1K-4  2.03GB/s ± 0%   1,15
```
