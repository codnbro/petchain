// go\cryptograph\ecdsa\ex2_sign\ecdsa_sign.go
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
R: 48439783365669796230930455383386064375961439678092708659311173964424855926272, S: 94532082945458256097695158499530859291744523618893303561311826751643238547812
Signature: 4843978336566979623093045538338606437596143967809270865931117396442485592627294532082945458256097695158499530859291744523618893303561311826751643238547812
SignatureASN1: [48 69 2 32 44 31 87 89 21 131 130 151 169 132 1 108 35 83 54 218 130 162 12 118 75 132 62 94 23 183 8 116 157 123 22 102 2 33 0 149 101 51 95 113 20 113 214 187 93 27 80 249 26 217 224 67 226 123 5 130 191 114 238 59 132 141 230 31 12 152 77]
*/
