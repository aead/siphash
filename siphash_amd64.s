// Copyright (c) 2016 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// +build amd64, !gccgo, !appengine

#define ROUND(v0, v1, v2, v3) \
	ADDQ v1, v0;  \
	RORQ $51, v1; \
	ADDQ v3, v2;  \
	XORQ v0, v1;  \
	RORQ $48, v3; \
	RORQ $32, v0; \
	XORQ v2, v3;  \
	ADDQ v1, v2;  \
	ADDQ v3, v0;  \
	RORQ $43, v3; \
	RORQ $47, v1; \
	XORQ v0, v3;  \
	XORQ v2, v1;  \
	RORQ $32, v2

// core(hVal *[4]uint64, msg []byte)
TEXT ·core(SB), 4, $0-32
	MOVQ hVal+0(FP), AX
	MOVQ msg_base+8(FP), SI
	MOVQ msg_len+16(FP), BX
	MOVQ 0(AX), R9
	MOVQ 8(AX), R10
	MOVQ 16(AX), R11
	MOVQ 24(AX), R12
	ANDQ $-8, BX

loop:
	MOVQ 0(SI), DX
	XORQ DX, R12
	ROUND(R9, R10, R11, R12)
	ROUND(R9, R10, R11, R12)
	XORQ DX, R9

	LEAQ 8(SI), SI
	SUBQ $8, BX
	JNZ  loop

	MOVQ R9, 0(AX)
	MOVQ R10, 8(AX)
	MOVQ R11, 16(AX)
	MOVQ R12, 24(AX)
	RET

// finalize64(hVal *[4]uint64, block *[8]byte) uint64
TEXT ·finalize64(SB), 4, $0-24
	MOVQ hVal+0(FP), SI
	MOVQ block+8(FP), BX
	MOVQ 0(BX), CX
	MOVQ 0(SI), R9
	MOVQ 8(SI), R10
	MOVQ 16(SI), R11
	MOVQ 24(SI), R12

	XORQ CX, R12
	ROUND(R9, R10, R11, R12)
	ROUND(R9, R10, R11, R12)
	XORQ CX, R9

	NOTB R11
	ROUND(R9, R10, R11, R12)
	ROUND(R9, R10, R11, R12)
	ROUND(R9, R10, R11, R12)
	ROUND(R9, R10, R11, R12)
	XORQ R9, R10
	XORQ R11, R12
	XORQ R10, R12

	MOVQ R12, ret+16(FP)
	RET

// finalize128(tag *[16]byte, hVal *[4]uint64, block *[8]byte)
TEXT ·finalize128(SB), 4, $0-24
	MOVQ tag+0(FP), DI
	MOVQ hVal+8(FP), SI
	MOVQ block+16(FP), BX
	MOVQ 0(BX), CX
	MOVQ 0(SI), R9
	MOVQ 8(SI), R10
	MOVQ 16(SI), R11
	MOVQ 24(SI), R12

	XORQ CX, R12
	ROUND(R9, R10, R11, R12)
	ROUND(R9, R10, R11, R12)
	XORQ CX, R9

	XORQ $0xEE, R11
	ROUND(R9, R10, R11, R12)
	ROUND(R9, R10, R11, R12)
	ROUND(R9, R10, R11, R12)
	ROUND(R9, R10, R11, R12)
	MOVQ R9, R13
	XORQ R10, R13
	XORQ R11, R13
	XORQ R12, R13

	XORQ $0xDD, R10
	ROUND(R9, R10, R11, R12)
	ROUND(R9, R10, R11, R12)
	ROUND(R9, R10, R11, R12)
	ROUND(R9, R10, R11, R12)
	XORQ R9, R10
	XORQ R11, R12
	XORQ R10, R12

	MOVQ R13, 0(DI)
	MOVQ R12, 8(DI)
	RET
