package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CheckConsent() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	consentFile := filepath.Join(home, ".smartcommit_consent")

	// Check if consent already given
	if _, err := os.Stat(consentFile); err == nil {
		return nil
	}

	// Show warning
	fmt.Println("\nSECURITY NOTICE:")
	fmt.Println("SmartCommit sends your git diffs to Google Gemini for analysis.")
	fmt.Println("Sensitive data is filtered, but you should review output before committing.")
	fmt.Println("\nDo NOT use on repositories with highly sensitive/proprietary code.")
	fmt.Print("\nDo you consent to sending diffs to Google Gemini? (yes/no): ")

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	response = strings.TrimSpace(strings.ToLower(response))

	if response != "yes" && response != "y" {
		return fmt.Errorf("consent not given")
	}

	// Save consent
	return os.WriteFile(consentFile, []byte("consented"), 0600)
}
