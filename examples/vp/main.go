// go\vp\main.go
package main

import (
	core "core"
	"fmt"
)

func main() {
	issuerKeyEcdsa := core.NewEcdsa()
	holderKeyEcdsa := core.NewEcdsa()

	// Issuer DID 생성.
	issuerDid, _ := core.NewDID("petchain", issuerKeyEcdsa.PublicKeyBase58())

	// Issuer DID Document 생성.
	verificationId := fmt.Sprintf("%s#keys-1", issuerDid)
	verificationMethod := []core.VerificationMethod{
		{
			Id:                 verificationId,
			Type:               "EcdsaSecp256k1VerificationKey2019",
			Controller:         issuerDid.String(),
			PublicKeyMultibase: issuerKeyEcdsa.PublicKeyMultibase(),
		},
	}
	didDocument := core.NewDIDDocument(issuerDid.String(), verificationMethod)
	core.RegisterDid(issuerDid.String(), didDocument.String())
	fmt.Println("Issuer DID Document: ", didDocument)

	// Holder DID 생성.
	holderDid, _ := core.NewDID("petchain", holderKeyEcdsa.PublicKeyBase58())

	// Holder DID Document 생성.
	verificationIdForHolder := fmt.Sprintf("%s#keys-1", holderDid)
	verificationMethodForHolder := []core.VerificationMethod{
		{
			Id:                 verificationIdForHolder,
			Type:               "EcdsaSecp256k1VerificationKey2019",
			Controller:         holderDid.String(),
			PublicKeyMultibase: holderKeyEcdsa.PublicKeyMultibase(),
		},
	}
	didDocumentForHolder := core.NewDIDDocument(holderDid.String(), verificationMethodForHolder)
	core.RegisterDid(holderDid.String(), didDocumentForHolder.String())

	// VC 생성.
	vc1, _ := core.NewVC(
		"1234567890",
		[]string{"VerifiableCredential", "AlumniCredential"},
		issuerDid.String(),
		map[string]interface{}{
			"id": "1234567890",
			"alumniOf": map[string]interface{}{
				"id": "1234567",
				"name": []map[string]string{
					{
						"value": "Example University",
						"lang":  "en",
					}, {
						"value": "Exemple d'Université",
						"lang":  "fr",
					},
				},
			},
		},
	)

	vc2, _ := core.NewVC(
		"123456789012345",
		[]string{"VerifiableCredential", "CertificationOfEmployee"},
		issuerDid.String(),
		map[string]interface{}{
			"id":          "123456789012121212",
			"company":     "petchain Co.",
			"joiningDate": "2020.01.01",
			"name":        "Hong kil-dong",
			"birthDate":   "2000.01.01",
		},
	)
	vc1Token, _ := vc1.GenerateJWT(verificationId, issuerKeyEcdsa.PrivateKey)
	vc2Token, _ := vc2.GenerateJWT(verificationId, issuerKeyEcdsa.PrivateKey)
	vcs := []string{vc1Token, vc2Token}

	vp, err := core.NewVP(
		"121212121",
		[]string{"VerifiablePresentation", "TestPresentation"},
		holderDid.String(),
		vcs,
	)

	if err != nil {
		panic(fmt.Sprintf("Error: ", err))
	}

	fmt.Println("VP: ", vp)

	vpToken := vp.GenerateJWT(verificationIdForHolder, holderKeyEcdsa.PrivateKey)
	fmt.Println("VP Token: ", vpToken)
}
