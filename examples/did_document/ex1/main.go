// examples/did_document/ex1/main.go
package main

import (
	"fmt"
	"log"
	core "petchain/core"
)

func main() {
	var method = "petchain"

	// 1. 키생성(ECDSA)
	kms := new(core.ECDSAManager)
	kms.Generate()

	// 2. DID 생성.
	did, err := core.NewDID(method, kms.PublicKeyMultibase())
	if err != nil {
		log.Printf("Failed to generate DID, error: %v\n", err)
	}

	// 3. DID Document 생성.
	verificationId := fmt.Sprintf("%s#keys-1", did)
	verificationMethod := []core.VerificationMethod{
		{
			Id:                 verificationId,
			Type:               core.VERIFICATION_KEY_TYPE_SECP256K1,
			Controller:         did.String(),
			PublicKeyMultibase: kms.PublicKeyMultibase(),
		},
	}
	didDocument := core.NewDIDDocument(did.String(), verificationMethod)

	fmt.Println("### Generate DID & DID Document ###")
	fmt.Printf("did => %s\n", did)
	fmt.Printf("did document => %+v\n", didDocument)

}

/*
### Generate DID & DID Document ###
did => did:petchain:8YMjRTauuE2ojny8Y6gqwXyXwRUxfoKukBaybC7uDFZF
did document => {"@context":["https://www.w3.org/ns/did/v1"],"id":"did:petchain:8YMjRTauuE2ojny8Y6gqwXyXwRUxfoKukBaybC7uDFZF","verificationMethod":[{"id":"did:petchain:8YMjRTauuE2ojny8Y6gqwXyXwRUxfoKukBaybC7uDFZF#keys-1","type":"EcdsaSecp256k1VerificationKey2019","controller":"did:petchain:8YMjRTauuE2ojny8Y6gqwXyXwRUxfoKukBaybC7uDFZF","PublicKeyMultibase":"zaSq9DsNNvGhYxYyqA9wd2eduEAZ5AXWgJTbTKBFwsUrCkYKFvSEku2mWP8Kw6nrD8pbLjeyPjdi3TBBvsmEx68uALJwd1UCci1JKgyRDxJRAMkPYkTudX5jUXnEM"}]}
*/
