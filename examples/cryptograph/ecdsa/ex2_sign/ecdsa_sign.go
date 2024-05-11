// examples/cryptograph/ecdsa/ex2_sign/ecdsa_sign.go
package main

import (
	"crypto/ecdsa"    // ECDSA 알고리즘을 사용하기 위한 패키지
	"crypto/elliptic" // 타원 곡선을 사용하기 위한 패키지
	"crypto/rand"     // 안전한 난수 생성을 위한 패키지
	"crypto/sha256"   // SHA-256 해시 알고리즘을 사용하기 위한 패키지
	"fmt"             // 출력을 위한 패키지
	"log"             // 로그 출력을 위한 패키지
	"math/big"        // 큰 정수를 다루기 위한 패키지
)

// Signature 구조체는 ECDSA 서명의 R과 S 값을 저장합니다.
type Signature struct {
	R *big.Int
	S *big.Int
}

// String 메서드는 Signature 구조체의 R과 S 값을 문자열로 변환합니다.
func (s Signature) String() string {
	return s.R.String() + s.S.String()
}

// sign 함수는 주어진 해시 값(digest)에 대해 ECDSA 서명을 생성합니다.
func sign(digest []byte, pvKey *ecdsa.PrivateKey) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, pvKey, digest)
	if err != nil {
		return nil, err
	}
	signature := &Signature{
		R: r,
		S: s,
	}
	return signature, nil
}

// SignASN1 함수는 ASN.1 형식으로 ECDSA 서명을 생성합니다.
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

func main() {
	// ECDSA 키 쌍을 생성합니다.
	pvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Println("ECDSA Keypair generation was Fail!")
	}
	// 서명할 메시지
	msg := "Welcome to PetChain!"
	digest := sha256.Sum256([]byte(msg))
	// 메시지에 대한 서명을 생성합니다.
	signature, err := sign(digest[:], pvKey)
	if err != nil {
		log.Printf("Failed to sign msg.")
	}
	// ASN.1 형식으로 서명을 생성합니다.
	signatureASN1, err := SignASN1(digest[:], pvKey)
	if err != nil {
		log.Printf("Failed to sign msg.")
	}
	// 생성된 메시지, 해시, 서명을 출력합니다.
	fmt.Printf("########## Sign ##########\n")
	fmt.Printf("===== Message =====\n")
	fmt.Printf("Msg: %s\n", msg)
	fmt.Printf("Digest: %x\n", digest)
	fmt.Printf("R: %s, S: %s\n", signature.R, signature.S)
	fmt.Printf("Signature: %+v\n", signature.String())
	fmt.Printf("SignatureASN1: %+v\n", signatureASN1)
}

/*
########## Sign ##########
===== Message =====
Msg: Welcome to PetChain!
Digest: 936a1e8aa8b2bf2fb8ad38ef8510f32cf4d237fb95da4696f4e363a8a8f8bc7e
R: 112811599017160678090225441090347579709356993102421075155469138500751950938118, S: 54417468177765585266835917708901103345468343955932963410250667618487982300234
Signature: 11281159901716067809022544109034757970935699310242107515546913850075195093811854417468177765585266835917708901103345468343955932963410250667618487982300234
SignatureASN1: [48 69 2 33 0 151 195 191 120 11 229 213 128 74 20 13 227 238 25 64 77 20 2 129 172 91 112 183 32 176 71 22 17 229 152 177 217 2 32 51 105 199 8 255 128 145 158 106 214 121 208 73 160 151 151 58 223 166 56 44 80 12 16 255 231 187 39 201 72 143 222]

종료 코드 0(으)로 완료된 프로세스
*/
