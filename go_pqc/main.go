package main

import "github.com/open-quantum-safe/liboqs-go/oqs"

func main() {
	msg := []byte("test")
	var signer, verifier oqs.Signature
	sigName := "Dilithium2"
	defer signer.Clean()
	defer verifier.Clean()
	// ignore potential errors everywhere
	_ = signer.Init(sigName, nil)
	_ = verifier.Init(sigName, nil)
	pubKey, _ := signer.GenerateKeyPair()
	for i := 0; i < 100; i++ {
		go func() {
			for {
				signature, _ := signer.Sign(msg)
				verifier.Verify(msg, signature, pubKey)
			}
		}()
	}
}
