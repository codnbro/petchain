package main

import (
	"fmt"

	"github.com/comnics/did-example/util"
)

func main() {
	str := "Hello, PetChain."
	hash1 := util.MakeHash(str)
	hash2 := util.MakeHashBase58(str)
	hash3 := util.MakeHashHex(str)
	fmt.Println("Plain Text: ", str)
	fmt.Println("MakeHash: ", hash1)
	fmt.Println("HashBase58: ", hash2)
	fmt.Println("MakeHashHex: ", hash3)
}

/*
Plain Text:  Hello, PetChain.
MakeHash:  [197 116 236 85 57 34 244 144 230 0 42 249 202 199 206 66 83 64 229 121 168 208 189 192 93 179 143 16 110 223 232 0]
HashBase58:  EHnizE9DMF449znNc7t1qDddPj3BvrFPfffZesSxaFNF
MakeHashHex:  c574ec553922f490e6002af9cac7ce425340e579a8d0bdc05db38f106edfe800
*/
