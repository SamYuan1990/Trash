package main

import (
	"fmt"

	"github.com/tuneinsight/lattigo/v3/bfv"
	"github.com/tuneinsight/lattigo/v3/rlwe"
)

// create key
func createKey(params bfv.Parameters) (sk *rlwe.SecretKey, pk *rlwe.PublicKey) {
	// Rider's keygen
	kgen := bfv.NewKeyGenerator(params)

	return kgen.GenKeyPair()
}

// encrypt
func encrypt(params bfv.Parameters, encryptor bfv.Encryptor, coeffs interface{}) *bfv.Ciphertext {
	encoder := bfv.NewEncoder(params)
	riderPlaintext := bfv.NewPlaintext(params)
	encoder.Encode(coeffs, riderPlaintext)
	return encryptor.EncryptNew(riderPlaintext)
}

// decrypt
func decrypt(params bfv.Parameters, riderSk *rlwe.SecretKey, ciphertext *bfv.Ciphertext) []uint64 {
	encoder := bfv.NewEncoder(params)
	decryptor := bfv.NewDecryptor(params, riderSk)
	return encoder.DecodeUintNew(decryptor.DecryptNew(ciphertext))
}

func main() {
	paramDef := bfv.PN13QP218
	paramDef.T = 0x3ee0001

	params, err := bfv.NewParametersFromLiteral(paramDef)
	if err != nil {
		panic(err)
	}
	riderSk, riderPk := createKey(params)

	encryptorRiderPk := bfv.NewEncryptor(params, riderPk)

	//encryptorRiderSk := bfv.NewEncryptor(params, riderSk)
	Ciphertext := encrypt(params, encryptorRiderPk, []uint64{8})
	decryptoRS := decrypt(params, riderSk, Ciphertext)
	fmt.Println(decryptoRS[0])
	evaluator := bfv.NewEvaluator(params, rlwe.EvaluationKey{})

	// add
	// 2+8=10
	Ciphertext2 := encrypt(params, encryptorRiderPk, []uint64{2})
	evaluator.Add(Ciphertext, Ciphertext2, Ciphertext)
	decryptoRS = decrypt(params, riderSk, Ciphertext)
	fmt.Println(decryptoRS[0])
	// sub
	// 10-2=8
	evaluator.Sub(Ciphertext, Ciphertext2, Ciphertext)
	decryptoRS = decrypt(params, riderSk, Ciphertext)
	fmt.Println(decryptoRS[0])
	// mux
	// 8*2=16
	Ciphertext3 := evaluator.MulNew(Ciphertext, Ciphertext2)
	decryptoRS = decrypt(params, riderSk, Ciphertext3)
	fmt.Println(decryptoRS[0])
}
