package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

func main() {
	msg := []byte("test")
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	for i := 0; i < 100; i++ {
		go func() {
			for {
				r, s, _ := ecdsa.Sign(rand.Reader, priv, msg)
				ecdsa.Verify(&priv.PublicKey, msg, r, s)
			}
		}()
	}
}
