// go\cryptograph\ecdsa\ex3_verify\ecdsa_verify.go
package main

import (
	"crypto/ecdsa"    // ECDSA 알고리즘을 위한 패키지
	"crypto/elliptic" // 타원 곡선 암호화를 위한 패키지
	"crypto/rand"     // 암호학적으로 안전한 난수 생성기 패키지
	"crypto/sha256"   // SHA-256 해시 함수를 위한 패키지
	"fmt"             // 출력을 위한 패키지
	"log"             // 로그 출력을 위한 패키지
	"math/big"        // 큰 수를 다루기 위한 패키지
)

// Signature 구조체는 서명의 R, S 값을 저장합니다.
type Signature struct {
	R *big.Int
	S *big.Int
}

// String 메서드는 Signature 구조체의 문자열 표현을 반환합니다.
func (s Signature) String() string {
	return s.R.String() + s.S.String()
}

// sign 함수는 주어진 메시지의 다이제스트와 개인 키를 사용하여 ECDSA 서명을 생성합니다.
func sign(digest []byte, pvKey *ecdsa.PrivateKey) (*Signature, error) {
	r := big.NewInt(0)
	s := big.NewInt(0)

	// ecdsa.Sign 함수는 ECDSA 서명을 생성합니다.
	r, s, err := ecdsa.Sign(rand.Reader, pvKey, digest)
	if err != nil {
		return nil, err
	}

	// Signature 구조체에 R, S 값을 설정합니다.
	signature := &Signature{
		R: r,
		S: s,
	}
	return signature, nil
}

// SignASN1 함수는 ASN.1 인코딩된 서명을 생성합니다.
func SignASN1(digest []byte, pvKey *ecdsa.PrivateKey) ([]byte, error) {

	signature, err := ecdsa.SignASN1(rand.Reader, pvKey, digest)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

// SignToString 함수는 서명을 문자열로 변환합니다.
func SignToString(digest []byte, pvKey *ecdsa.PrivateKey) (string, error) {
	signature, err := sign(digest, pvKey)
	if err != nil {
		return "", err
	}

	return signature.String(), nil
}

// verify 함수는 서명이 주어진 메시지의 다이제스트와 일치하는지 검증합니다.
func verify(signature *Signature, digest []byte, pbKey *ecdsa.PublicKey) bool {
	return ecdsa.Verify(pbKey, digest, signature.R, signature.S)
}

// verifyASN1 함수는 ASN.1 인코딩된 서명을 검증합니다.
func verifyASN1(signature []byte, digest []byte, pbKey *ecdsa.PublicKey) bool {
	return ecdsa.VerifyASN1(pbKey, digest, signature)
}

func main() {
	// ECDSA 키 쌍을 생성합니다.
	pvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Println("ECDSA Keypair generation was Fail!")
	}

	// 서명할 메시지를 정의하고, 그 메시지의 SHA-256 다이제스트를 계산합니다.
	msg := "Welcome to PetChain!"
	digest := sha256.Sum256([]byte(msg))

	// sign 함수를 사용하여 메시지에 대한 서명을 생성합니다.
	signature, err := sign(digest[:], pvKey)
	if err != nil {
		log.Printf("Failed to sign msg.")
	}

	// 서명 정보를 출력합니다.
	fmt.Printf("########## Sign ##########\n")
	fmt.Printf("===== Message =====\n")
	fmt.Printf("Msg: %s\n", msg)
	fmt.Printf("Digest: %x\n", digest)
	fmt.Printf("R: %s, S: %s\n", signature.R, signature.S)
	fmt.Printf("Signature: %+v\n", signature.String())

	// 서명을 검증합니다.
	pbKey := &pvKey.PublicKey
	ret := verify(signature, digest[:], pbKey)
	fmt.Println("########## Verification ##########")
	if ret {
		fmt.Println("Signature verifies")
	} else {
		fmt.Println("Signature does not verify")
	}

	// ASN.1 인코딩된 서명을 생성하고 검증합니다.
	signatureASN1, err := SignASN1(digest[:], pvKey)
	if err != nil {
		log.Printf("Failed to sign msg.")
	}
	ret = verifyASN1(signatureASN1, digest[:], pbKey)

	fmt.Println("########## Verification 2 ##########")
	if ret {
		fmt.Println("Signature verifies")
	} else {
		fmt.Println("Signature does not verify")
	}

	// 변경된 메시지에 대한 검증을 시도합니다.
	msg2 := "Hello, PetChain World."
	digest2 := sha256.Sum256([]byte(msg2))
	ret = verify(signature, digest2[:], pbKey)

	fmt.Println("\n########## Verification 3: Other message ##########")
	if ret {
		fmt.Printf("Signature verifies")
	} else {
		fmt.Printf("Signature does not verify")
	}

	// 다른 키를 사용하여 검증을 시도합니다.
	pvKey2, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pbKey2 := &pvKey2.PublicKey
	ret = verify(signature, digest[:], pbKey2)

	fmt.Println("\n########## Verification 4: Other key ##########")
	if ret {
		fmt.Printf("Signature verifies")
	} else {
		fmt.Printf("Signature does not verify")
	}
}

/*
########## Sign ##########
===== Message =====
Msg: Welcome to PetChain!
Digest: 936a1e8aa8b2bf2fb8ad38ef8510f32cf4d237fb95da4696f4e363a8a8f8bc7e
R: 47837728373606562309343831003310252244374679233020665829834211827484942289147, S: 53084870536044270369299892138993173495918676263042610284987754719788595287328
Signature: 4783772837360656230934383100331025224437467923302066582983421182748494228914753084870536044270369299892138993173495918676263042610284987754719788595287328
########## Verification ##########
Signature verifies
########## Verification 2 ##########
Signature verifies

########## Verification 3: Other message ##########
Signature does not verify
########## Verification 4: Other key ##########
Signature does not verify
*/
