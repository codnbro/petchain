// go\cryptograph\ecdsa\ex1\main.go
package main

import (
	"crypto/ecdsa"    // ECDSA 알고리즘을 사용하기 위한 패키지
	"crypto/elliptic" // 타원 곡선 암호화를 사용하기 위한 패키지
	"crypto/rand"     // 안전한 난수 생성기를 제공하는 패키지
	"fmt"             // 출력을 위한 패키지
	"log"             // 로그 출력을 위한 패키지
)

func main() {
	// ECDSA 키 쌍을 생성합니다. 이 때 elliptic.P256() 타원 곡선을 사용합니다.
	// rand.Reader는 안전한 난수 생성기를 제공합니다.
	pvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil { // 키 생성에 실패한 경우
		log.Println("ECDSA Keypair generation was Fail!")
	}
	pbKey := &pvKey.PublicKey // 생성된 개인 키로부터 공개 키 정보를 가져옵니다.

	// 키 정보를 출력합니다.
	fmt.Printf("########## Key pair ##########\n")
	fmt.Printf("===== Private Key =====\n")
	// 개인 키 출력. pvKey.D는 개인 키의 D 값(big.Int)을 나타냅니다.
	fmt.Printf("Private Key: %x\n", pvKey.D)
	fmt.Printf("===== Public Key(X, Y) =====\n")
	// 공개 키의 X, Y 좌표를 문자열 형태로 출력합니다.
	// 공개 키의 X, Y 좌표는 big.Int 타입입니다.
	fmt.Printf("X=%s Y=%s\n", pbKey.X, pbKey.Y)
	// 공개 키의 X, Y 좌표를 16진수 형태로 출력합니다.
	fmt.Printf("  Hex: X=%x Y=%x\n\n", pbKey.X.Bytes(), pbKey.Y.Bytes())
}

/*
########## Key pair ##########
===== Private Key =====
Private Key: 2453921946fcdd9afd1f069947c1d0ba82333bf9cc853574b874414f0db7c7d3
===== Public Key(X, Y) =====
X=82244380591317224841127915819760090559931434811866224211940124113623427589984 Y=39947260302688777166872715786644428200613305892706303968527425946669829490275
  Hex: X=b5d4ab777ce7b10126e2b239b58743cb67e84445c2267edf412cb12265402760 Y=5851591c2daeafa4517bc9ffe40cc1933aad4203f4052acd27d841ef0f4e8663
*/
