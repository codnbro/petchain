// examples/cryptograph/ecdsa/main.go
package main

import (
	"crypto/sha256" // SHA-256 해시 함수를 사용하기 위한 패키지
	"fmt"           // 표준 입출력을 위한 패키지
	"petchain/core" // ECDSA 관련 기능을 담은 사용자 정의 패키지
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

/*
### Start Main() ###
########## Key pair ##########
===== Private Key =====
Private Key: 40476e71364abf879c74db29e4ccae6591d19f005cfb81d1205f55e560f8391b
===== Public Key(X, Y) =====
X=77630462281671674068344944796292394392909475495939538351801756480451709277490 Y=14515582290040653137612206093981360160140428241476510052109045341513716066837
  Hex: X=aba148d9c243a15d364b016be5079724669aaa3bf27188453e811f7c8aec9532 Y=201787471c3b704623e9102632d864613440a2d0ed616c20445eda31aeb3ea15

########## Sign ##########
===== Message =====
Msg: Welcome to petchain
Digest: a57bf424cf62cf68b5e3ab2dc079c4dc4601c370ab690890c2456074f87d92f3
R: 36924836527336152708517808397388754058450812760064433535927151260435823631962, S: 46787721747355349076604268835117451914787831570388274709715963144922241691453

########## Verification ##########
Signature verifies


########## SignASN1 ##########
===== Message =====
Msg: Welcome to petchain
Digest: a57bf424cf62cf68b5e3ab2dc079c4dc4601c370ab690890c2456074f87d92f3
Signature: 304402203c6bd53511a5c618d3bb8c09912d9d855ca341d18ec101ff92534605c325a63302201da0ab3b82e7b324d72fa8e7c97a70d2c40625eacfd1e4742ccbd04d7d8e1e6a

########## Verification ASN1 ##########
SignatureASN1 verifies
*/
