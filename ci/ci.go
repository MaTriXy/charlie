package ci

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/google/go-github/github"
)

type Server struct {
}

// New Server.
func New() http.Handler {
	return &Server{}
}

// ServeHTTP implements the http.Handler interface.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload github.WebHookPayload
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "could not parse body", http.StatusBadRequest)
		return
	}

	// Create a temp directory.
	log.Println("creating temp directory")
	tempDir, err := ioutil.TempDir("", *payload.Repo.FullName)
	if err != nil {
		http.Error(w, "could not clone repo", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempDir)
	log.Println("created", tempDir)

	// Clone the repository.
	log.Println("cloning", *payload.Repo.SSHURL)
	if err := exec.Command("git", "clone", *payload.Repo.SSHURL, tempDir).Run(); err != nil {
		http.Error(w, "could not clone repo", http.StatusInternalServerError)
		return
	}
	log.Println("cloned", *payload.Repo.SSHURL)

	// Build the project.
	log.Println("changing directory", tempDir)
	if err := os.Chdir(tempDir); err != nil {
		http.Error(w, "could not build repo", http.StatusInternalServerError)
		return
	}
	log.Println("building project")
	buildCmd := exec.Command("./gradlew", "build")
	lw := &logWriter{}
	buildCmd.Stdout = lw
	buildCmd.Stderr = lw
	if err := buildCmd.Run(); err != nil {
		http.Error(w, "could not build project", http.StatusInternalServerError)
		return
	}
	log.Println("succesfully built repo")

	fmt.Fprintf(w, "Built %s", payload.Repo)
}

// An io.Writer implementation writes to the default logger.
type logWriter struct {
}

func (l *logWriter) Write(p []byte) (n int, err error) {
	log.Print(string(p))
	return len(p), nil
}
