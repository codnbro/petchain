// examples/cryptograph/ecdsa/ex1/main.go
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
Private Key: b3899c0e882b1ef575c7b6fb980ebbe998cf6efa69fabf03c10e82cb2b0cb2bd
===== Public Key(X, Y) =====
X=115102666118831156187238838443187960212824889632454153919827240386472476980183 Y=97590057592975384469387556263311691163557371748490357142733889849854708200008
  Hex: X=fe79cce348eeacdf9ec510ff7d9f733b2e45d36ab7c30241e0dff179b638b3d7 Y=d7c203e2b2b2bd6c62f504400d094d34cc1d775eba335ae086a32f18c864b248
*/
