package git

import (
	"regexp"
)

func FilterSensitiveData(diff string) string {
	// Match common secret key/value pairs (more permissive charset)
	secretPattern := regexp.MustCompile(`(?i)(api[_-]?key|token|secret|password|passwd|auth|access[_-]?key)\s*[:=]\s*["']?[a-zA-Z0-9_\-\/\+=\.]{8,}["']?`)
	diff = secretPattern.ReplaceAllString(diff, "$1: [REDACTED]")

	// Match Bearer tokens
	bearerPattern := regexp.MustCompile(`(?i)authorization\s*[:=]\s*["']?bearer\s+[a-zA-Z0-9\-\._~\+\/]+=*["']?`)
	diff = bearerPattern.ReplaceAllString(diff, "authorization: [REDACTED]")

	// Match GitHub tokens
	githubPattern := regexp.MustCompile(`\b(ghp_[a-zA-Z0-9]{36}|github_pat_[a-zA-Z0-9_]{50,})\b`)
	diff = githubPattern.ReplaceAllString(diff, "[GITHUB_TOKEN_REDACTED]")

	// Match AWS Access Key IDs
	awsAccessKey := regexp.MustCompile(`\bAKIA[0-9A-Z]{16}\b`)
	diff = awsAccessKey.ReplaceAllString(diff, "[AWS_ACCESS_KEY_REDACTED]")

	// Match private key blocks
	privateKeyPattern := regexp.MustCompile(`-----BEGIN [A-Z ]+PRIVATE KEY-----[\s\S]+?-----END [A-Z ]+PRIVATE KEY-----`)
	diff = privateKeyPattern.ReplaceAllString(diff, "[PRIVATE_KEY_REDACTED]")

	// Emails
	emailPattern := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
	diff = emailPattern.ReplaceAllString(diff, "[EMAIL_REDACTED]")

	// Credit card numbers (still basic)
	ccPattern := regexp.MustCompile(`\b\d{4}[\s-]?\d{4}[\s-]?\d{4}[\s-]?\d{4}\b`)
	diff = ccPattern.ReplaceAllString(diff, "[CC_REDACTED]")

	return diff
}
