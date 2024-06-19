package main

import (
	"crypto"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	jose "gopkg.in/square/go-jose.v2"
	"log"
	"os"
)


func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Need to specify infile")
		return
	}
	infile := os.Args[1]
	pembytes, err := os.ReadFile(infile)
	if err != nil {
		log.Fatalf("Error reading pem key: %v\n", err)
		return
	}
	pemBlock, _ := pem.Decode(pembytes)
	key, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)

	if err != nil {
		log.Fatalf("Error reading pem pubkey: %v\n", err)
	}

	kid, err := keyIDFromPublicKey(key)
	alg := jose.RS256

	var keys []jose.JSONWebKey
	keys = append(keys, jose.JSONWebKey{
		Key:       key,
		KeyID:     kid,
		Algorithm: string(alg),
		Use:       "sig",
	})

	keySet, err := json.MarshalIndent(jose.JSONWebKeySet{Keys: keys}, "", "    ")
  _, err = fmt.Println(string(keySet))
  if err != nil {
    log.Fatalf("Error printing key to stdout: %v", err)
  }
}

// taken from https://github.com/kubernetes/kubernetes/blob/master/pkg/serviceaccount/jwt.go
func keyIDFromPublicKey(publicKey any) (string, error) {
	publicKeyDERBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", fmt.Errorf("failed to serialize public key to DER format: %v", err)
	}

	hasher := crypto.SHA256.New()
  _, err = hasher.Write(publicKeyDERBytes)
  if err != nil {
    return "", err
  }
	publicKeyDERHash := hasher.Sum(nil)

	keyID := base64.RawURLEncoding.EncodeToString(publicKeyDERHash)

	return keyID, nil
}
