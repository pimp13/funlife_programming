package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

const (
	secret       = "YOUR_SECRET_KEY"
	projectsRoot = "/media/qut"
	targetBranch = "refs/heads/main"
)

type GithubPayload struct {
	Repository struct {
		Name string `json:"name"`
	} `json:"repository"`
	Ref string `json:"ref"`
}

func verifySignature(secret string, body []byte, signature string) bool {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	expectedMAC := mac.Sum(nil)

	sig := strings.TrimPrefix(signature, "sha256=")
	decodedSig, err := hex.DecodeString(sig)
	if err != nil {
		return false
	}

	return hmac.Equal(decodedSig, expectedMAC)
}

func main() {

}
