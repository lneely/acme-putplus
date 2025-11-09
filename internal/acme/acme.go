package acme

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"9fans.net/go/acme"
)

func GetCurrentAcmeWindowContent() ([]byte, error) {
	winIDStr := os.Getenv("winid")
	if winIDStr == "" {
		return nil, fmt.Errorf("not running in acme window (winid not set)")
	}

	winID, err := strconv.Atoi(winIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid winid: %w", err)
	}

	w, err := acme.Open(winID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open acme window: %w", err)
	}
	defer w.CloseFiles()

	content, err := w.ReadAll("body")
	if err != nil {
		return nil, fmt.Errorf("failed to read acme window body: %w", err)
	}

	return content, nil
}

func GetCurrentAcmeWindowName() (string, error) {
	winIDStr := os.Getenv("winid")
	if winIDStr == "" {
		return "", fmt.Errorf("not running in acme window (winid not set)")
	}

	winID, err := strconv.Atoi(winIDStr)
	if err != nil {
		return "", fmt.Errorf("invalid winid: %w", err)
	}

	w, err := acme.Open(winID, nil)
	if err != nil {
		return "", fmt.Errorf("failed to open acme window: %w", err)
	}
	defer w.CloseFiles()

	tag, err := w.ReadAll("tag")
	if err != nil {
		return "", fmt.Errorf("failed to read acme window tag: %w", err)
	}

	tagFields := strings.Fields(string(tag))
	if len(tagFields) == 0 {
		return "", fmt.Errorf("empty tag")
	}

	windowName := tagFields[0]
	if !filepath.IsAbs(windowName) {
		wd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("failed to get working directory: %w", err)
		}
		windowName = filepath.Join(wd, windowName)
	}

	return windowName, nil
}

func WriteToStderr(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}