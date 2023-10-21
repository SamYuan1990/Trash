package main

import (
	"testing"

	"github.com/open-quantum-safe/liboqs-go/oqs"
)

func BenchmarkEcdsaVerify(t *testing.B) {
	t.ReportAllocs()
	msg := []byte("test")
	var signer, verifier oqs.Signature
	sigName := "Dilithium2"
	defer signer.Clean()
	defer verifier.Clean()
	// ignore potential errors everywhere
	_ = signer.Init(sigName, nil)
	_ = verifier.Init(sigName, nil)
	pubKey, _ := signer.GenerateKeyPair()
	signature, _ := signer.Sign(msg)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		isverify, _ := verifier.Verify(msg, signature, pubKey)
		if !isverify {
			t.Fatal("Verify error\n")
		}
	}
}
