// go\cryptograph\ecdsa\main.go
package main

import (
	"crypto/sha256" // SHA-256 해시 함수를 사용하기 위한 패키지
	"fmt"           // 표준 입출력을 위한 패키지
	"go/core"       // ECDSA 관련 기능을 담은 사용자 정의 패키지
)

func main() {
	fmt.Println("### Start Main() ###")

	// ECDSA secp256k1
	var ecdsa core.ECDSAManager // ECDSAManager 구조체 인스턴스 생성
	ecdsa.Generate()            // ECDSA 키 쌍(공개/개인 키) 생성

	// 생성된 키 쌍 출력
	fmt.Printf("########## Key pair ##########\n")
	fmt.Printf("===== Private Key =====\n")
	fmt.Printf("Private Key: %x\n", ecdsa.PrivateKey.D) // 개인 키 출력
	fmt.Printf("===== Public Key(X, Y) =====\n")
	fmt.Printf("X=%s Y=%s\n", ecdsa.PublicKey.X, ecdsa.PublicKey.Y) // 공개 키 출력
	fmt.Printf("  Hex: X=%x Y=%x\n\n", ecdsa.PublicKey.X.Bytes(), ecdsa.PublicKey.Y.Bytes())

	msg := "Welcome to petchain"
	digest := sha256.Sum256([]byte(msg)) // 메시지에 대한 SHA-256 해시 생성

	// 메시지 서명
	signature, err := ecdsa.Sign(digest[:]) // 생성된 해시에 대해 서명
	if err != nil {
		fmt.Printf("Fail to sign to msg.")
	}

	// 서명 결과 출력
	fmt.Printf("########## Sign ##########\n")
	fmt.Printf("===== Message =====\n")
	fmt.Printf("Msg: %s\n", msg)
	fmt.Printf("Digest: %x\n", digest)
	fmt.Printf("R: %s, S: %s\n\n", signature.R, signature.S)

	// 서명 검증
	ret := ecdsa.Verify(signature, digest[:])

	// 검증 결과 출력
	fmt.Printf("########## Verification ##########\n")
	if ret {
		fmt.Printf("Signature verifies\n")
	} else {
		fmt.Printf("Signature does not verify\n")
	}

	// ASN1 형식으로 서명 및 검증
	signatureASN1, err := ecdsa.SignASN1(digest[:]) // ASN1 형식으로 서명
	if err != nil {
		fmt.Printf("Fail to sign to msg.")
	}

	// ASN1 서명 결과 출력
	fmt.Printf("\n\n########## SignASN1 ##########\n")
	fmt.Printf("===== Message =====\n")
	fmt.Printf("Msg: %s\n", msg)
	fmt.Printf("Digest: %x\n", digest)
	fmt.Printf("Signature: %x \n\n", signatureASN1)

	ret = ecdsa.VerifyASN1(signatureASN1, digest[:]) // ASN1 서명 검증

	// ASN1 검증 결과 출력
	fmt.Printf("########## Verification ASN1 ##########\n")
	if ret {
		fmt.Printf("SignatureASN1 verifies\n")
	} else {
		fmt.Printf("Signature does not verify\n")
	}
}
