package main

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/lkn/acme-put+/internal/acme"
)

func main() {
	content, err := acme.GetCurrentAcmeWindowContent()
	if err != nil {
		acme.WriteToStderr("Failed to get acme window content: %v\n", err)
		os.Exit(1)
	}

	windowName, err := acme.GetCurrentAcmeWindowName()
	if err != nil {
		acme.WriteToStderr("Failed to get acme window name: %v\n", err)
		os.Exit(1)
	}

	if err := writeWithPrivileges(content, windowName); err != nil {
		acme.WriteToStderr("Failed to write file with privileges: %v\n", err)
		os.Exit(1)
	}

	acme.WriteToStderr("Successfully wrote to: %s\n", windowName)
}

func writeWithPrivileges(content []byte, filePath string) error {
	filePath = filepath.Clean(filePath)
	
	cmd := exec.Command("pkexec", "tee", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	if _, err := stdin.Write(content); err != nil {
		stdin.Close()
		cmd.Wait()
		return err
	}

	if err := stdin.Close(); err != nil {
		cmd.Wait()
		return err
	}

	return cmd.Wait()
}