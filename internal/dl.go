package internal

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type DL interface {
	Search() (string, error)
}

type dl struct {
	Key       string
	MaxResult string
	URL       string
}

func NewDL(key, keyword, maxResult string) DL {
	return &dl{
		Key:       key,
		MaxResult: maxResult,
		URL:       keyword,
	}
}

func (d *dl) Search() (string, error) {
	currentDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}

	cmd := "node"
	args := []string{filepath.Join(currentDirectory, "internal/node-package"), d.MaxResult, d.Key, d.URL}
	process := exec.Command(cmd, args...)
	stdin, err := process.StdinPipe()

	if err != nil {
		fmt.Println(err)
	}
	defer stdin.Close()
	buf := new(bytes.Buffer)
	process.Stdout = buf
	process.Stderr = os.Stderr

	if err = process.Start(); err != nil {
		return "", errors.New(fmt.Sprintf("error occured: %s", err))
	}

	process.Wait()
	return buf.String(), nil
}
