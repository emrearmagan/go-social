/*
signer.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package oauth1

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"hash"
	"strings"
)

// Signer signs messages to create signed OAuth1 Requests.
type Signer interface {
	// Sign signs the message of the request
	Sign(key string, message string) (string, error)
	// Name returns the name of the signing method.
	Name() string
}

const (
	HMACSHA1  = "HMAC-SHA1"
	RSASHA1   = "RSA-SHA1"
	PLAINTEXT = "Plain text"
)

//HMAC
type HMACSigner struct {
	ConsumerSecret string
}

func hmacSign(signingKey, message string, algo func() hash.Hash) (string, error) {
	mac := hmac.New(algo, []byte(signingKey))
	mac.Write([]byte(message))
	signatureBytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signatureBytes), nil
}

func (s *HMACSigner) Sign(key, message string) (string, error) {
	signingKey := s.Key(key)
	return hmacSign(signingKey, message, sha1.New)

}

// Name returns the HMAC-SHA256 method.
func (s *HMACSigner) Name() string {
	return HMACSHA1
}

// Name returns the HMAC-SHA256 method.
func (s *HMACSigner) Key(token string) string {
	return strings.Join([]string{s.ConsumerSecret, token}, "&")

}

// RSA
//TODO if needed

// RSASigner RSA PKCS1-v1_5 signs SHA1 digests of messages using the given
// RSA private key.
type RSASigner struct {
	PrivateKey *rsa.PrivateKey
}

// Name returns the RSA-SHA1 method.
func (s *RSASigner) Name() string {
	return RSASHA1
}

// Sign uses RSA PKCS1-v1_5 to sign a SHA1 digest of the given message. The
// tokenSecret is not used with this signing scheme.
func (s *RSASigner) Sign(message string) (string, error) {
	digest := sha1.Sum([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.PrivateKey, crypto.SHA1, digest[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}
