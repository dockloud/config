package secret

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type DockerHandler struct{}

func NewDockerHandler() *DockerHandler {
	return &DockerHandler{}
}

func (h *DockerHandler) ReadSecret(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("empty secret name")
	}

	if filepath.Base(name) != name || name == "." || name == ".." {
		return "", fmt.Errorf("invalid secret name: %q", name)
	}

	secretsDir := os.Getenv("SECRETS_DIR")
	if secretsDir == "" {
		secretsDir = "/run/secrets"
	}

	fullPath := filepath.Join(secretsDir, name)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to get secret %q (%s): %w", name, fullPath, err)
	}

	return strings.TrimSpace(string(data)), nil
}
