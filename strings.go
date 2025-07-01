package u

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode"
)

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}
	firstLetter := rune(s[0])

	return unicode.IsUpper(firstLetter)
}

func ToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	var chunks []string
	var currentChunk []rune
	for i, letter := range s {
		if i == 0 || !unicode.IsUpper(letter) {
			currentChunk = append(currentChunk, unicode.ToLower(letter))
		} else {
			prevIsLower := unicode.IsLower(rune(s[i-1]))
			nextIsLower := true
			if i != len(s)-1 {
				nextIsLower = unicode.IsLower(rune(s[i+1]))
			}
			if prevIsLower || nextIsLower {
				chunks = append(chunks, string(currentChunk))
				currentChunk = currentChunk[:0]
			}
			currentChunk = append(currentChunk, unicode.ToLower(letter))
		}
	}

	if len(currentChunk) > 0 {
		chunks = append(chunks, string(currentChunk))
	}

	return strings.Join(chunks, "_")
}

func Capitalize(s string) string {
	if s == "" {
		return s
	}

	letters := []rune(s)

	letters[0] = unicode.ToUpper(letters[0])

	return string(letters)
}

func RandomString(byteLength int) (string, error) {
	b := make([]byte, byteLength)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// Encode the random bytes into a URL-safe base64 string
	return base64.URLEncoding.EncodeToString(b), nil
}

func AddMissingSuffix(s, suf string) string {
	if !strings.HasSuffix(s, suf) {
		return s + suf
	}
	return s
}

func AddMissingPrefix(s, pref string) string {
	if !strings.HasPrefix(s, pref) {
		return pref + s
	}
	return s
}

type indentConfig struct {
	Indent            string
	ScannerBufferSize int
}

type IndentOption func(*indentConfig)

func WithIndent(indent string) IndentOption {
	return func(c *indentConfig) {
		c.Indent = indent
	}
}

func WithScannerBufferSize(size int) IndentOption {
	return func(c *indentConfig) {
		c.ScannerBufferSize = size
	}
}

func IndentString(text string, opts ...IndentOption) (string, error) {
	if text == "" {
		return "", nil
	}
	cfg := indentConfig{
		Indent:            "  ",
		ScannerBufferSize: 64 * 1024,
	}

	for _, opt := range opts {
		opt(&cfg)
	}
	reader := strings.NewReader(text)
	var writer strings.Builder
	scanner := bufio.NewScanner(reader)
	buf := make([]byte, cfg.ScannerBufferSize)
	scanner.Buffer(buf, cfg.ScannerBufferSize)

	for scanner.Scan() {
		line := scanner.Text()
		_, err := writer.WriteString(cfg.Indent + line + "\n")
		if err != nil {
			return "", fmt.Errorf("Failed to write indented line: %w", err)
		}
	}
	return writer.String(), scanner.Err()
}

func IndentErrors(description string, errs error) error {
	if errs == nil {
		return nil
	}

	var sb strings.Builder
	sb.WriteString(description)
	sb.WriteString(":\n")

	indentedErrs, err := IndentString(errs.Error())
	if err != nil {
		fmt.Printf("Internal error indenting string: %v\n", err)
		return errs
	}

	sb.WriteString(indentedErrs)

	return errors.New(sb.String())
}

func ValidateDurationString(durationStr string) error {
	_, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("Invalid duration format '%s': %w", durationStr, err)
	}
	return nil
}

func SortString(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}

	return 0
}
