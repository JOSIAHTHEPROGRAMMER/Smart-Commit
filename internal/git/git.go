package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// RunCommand executes a git command and returns the output
func RunCommand(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("git command failed: %v - %s", err, stderr.String())
	}

	return strings.TrimSpace(out.String()), nil
}

// IsGitRepo checks if the current directory is a git repository
func IsGitRepo() bool {
	_, err := RunCommand("rev-parse", "--git-dir")
	return err == nil
}

// GetStatus returns the current git status
func GetStatus() (string, error) {
	return RunCommand("status", "--short")
}

// GetStagedDiff returns the diff of staged changes
func GetStagedDiff() (string, error) {
	if !IsGitRepo() {
		return "", fmt.Errorf("not a git repository")
	}

	diff, err := RunCommand("diff", "--cached")
	if err != nil {
		return "", err
	}

	if diff == "" {
		return "", fmt.Errorf("no staged changes found. Use 'git add' to stage changes")
	}

	return diff, nil
}

// HasStagedChanges checks if there are any staged changes
func HasStagedChanges() bool {
	output, err := RunCommand("diff", "--cached", "--name-only")
	if err != nil {
		return false
	}
	return output != ""
}

// Commit creates a git commit with the given message
func Commit(message string) error {
	if !IsGitRepo() {
		return fmt.Errorf("not a git repository")
	}

	_, err := RunCommand("commit", "-m", message)
	if err != nil {
		return fmt.Errorf("failed to commit: %v", err)
	}

	return nil
}
