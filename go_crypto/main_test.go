package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"
)

func BenchmarkEcdsaVerify(t *testing.B) {
	t.ReportAllocs()
	msg := []byte("test")
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	r, s, err := ecdsa.Sign(rand.Reader, priv, msg)
	if err != nil {
		t.Fatal(err)
	}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		isverify := ecdsa.Verify(&priv.PublicKey, msg, r, s)
		if !isverify {
			t.Fatal("Verify error\n")
		}
	}
}
