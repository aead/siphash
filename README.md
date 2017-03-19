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
Write_8-4     679MB/s ± 0%   3,51
Write_1K-4   2.07GB/s ± 0%   1,12
Sum64_8-4     258MB/s ± 0%   9,24
Sum64_1K-4   2.00GB/s ± 0%   1,16
Sum128_8-4   75.7MB/s ± 0%  31,79 
Sum128_1K-4  1.72GB/s ± 0%   1,35
```
