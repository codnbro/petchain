// go\vc\simple\main.go
package main

import (
	"example/core"
	"fmt"
	"os"
)

// Issuer에 의한 VC 발행 예시.
func main() {
	// 키생성(ECDSA) - 향후 KMS로 대체.
	// 키 이름을 이슈어 키라고 하고 암호화 방식도 키이름으로 해줌.
	issuerKeyEcdsa := core.NewEcdsa()

	// 이슈어의 DID 생성. (퍼블릭키)
	issuerDid, _ := core.NewDID("example", issuerKeyEcdsa.PublicKeyBase58())

	// DID Document 생성.
	verificationId := fmt.Sprintf("%s#keys-1", issuerDid)
	//verificationMethod := []core.VerificationMethod{
	//	{
	//		Id:                 verificationId,
	//		Type:               "EcdsaSecp256k1VerificationKey2019",
	//		Controller:         issuerDid.String(),
	//		PublicKeyMultibase: issuerKeyEcdsa.PublicKeyMultibase(),
	//	},
	//}
	//didDocument := core.NewDIDDocument(issuerDid.String(), verificationMethod)
	// DID와 DID Document를 VDR에 올려야 하나, 현재 생략.

	// 실제 VC 객체 생성. (매우 중요)
	vc, err := core.NewVC(
		"1234567890", // 아이디
		[]string{"VerifiableCredential", "AlumniCredential"}, // 목적과 이름
		issuerDid.String(), // 누가 발행했는지 발행자 DID
		map[string]interface{}{
			"id": "1234567890",
			"alumniOf": map[string]interface{}{
				"id": "1234567",
				"name": []map[string]string{
					{
						"value": "Example University", // 영어로는
						"lang":  "en",
					}, {
						"value": "Exemple d'Université", // 프랑스어로는
						"lang":  "fr",
					},
				},
			},
		},
	)

	if err != nil {
		fmt.Println("Failed creation VC.")
		os.Exit(0)
	}

	// VC에 Issuer의 private key로 서명한다.(JWT 사용)
	// 프라이빗 키를 넣어야 서명을 해줄수 있어서 프라이빗 키를 넣어준다.
	token, err := vc.GenerateJWT(verificationId, issuerKeyEcdsa.PrivateKey)
	fmt.Println("")
	fmt.Println("")
	fmt.Println(token)
	fmt.Println("")
	fmt.Println("")

	// 생성된 VC를 검증한다.(public key를 사용해서 검증)
	// 검증은 퍼블릭 키로 한다.
	res, _ := vc.VerifyJwt(token, issuerKeyEcdsa.PublicKey)

	if res {
		fmt.Println("VC is verified.")
	} else {
		fmt.Println("VC is Not verified.")
	}

}
