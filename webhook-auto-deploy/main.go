package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const (
	secret       = "MY_SUPER_SECRET"
	projectsRoot = "/root/projects"
	targetBranch = "refs/heads/main"
)

type GitHubPayload struct {
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

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	signature := r.Header.Get("X-Hub-Signature-256")
	if !verifySignature(secret, body, signature) {
		log.Println("❌ Invalid signature")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var payload GitHubPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Println("❌ JSON parse error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if payload.Ref != targetBranch {
		log.Println("ℹ️ Not main branch. Ignored.")
		w.WriteHeader(http.StatusOK)
		return
	}

	repoName := payload.Repository.Name
	projectPath := projectsRoot + "/" + repoName

	deployScript := projectPath + "/.deploy.sh"

	// بررسی وجود فایل deploy
	if _, err := os.Stat(deployScript); os.IsNotExist(err) {
		log.Printf("❌ Deploy script not found for project %s\n", repoName)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.Printf("🚀 Deploying project: %s\n", repoName)

	cmd := exec.Command("bash", deployScript)
	cmd.Dir = projectPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Printf("❌ Deploy failed for %s: %v\n", repoName, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Deploy completed for %s\n", repoName)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/deploy", handler)
	log.Println("🔥 Multi-project webhook server running on :9001")
	log.Fatal(http.ListenAndServe(":9001", nil))
}
