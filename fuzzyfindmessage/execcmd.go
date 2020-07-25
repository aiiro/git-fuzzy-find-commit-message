package fuzzyfindmessage

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

var (
	execCommand = exec.Command
	commandRun  = func(c *exec.Cmd) error {
		return c.Run()
	}
	commandOutput = func(c *exec.Cmd) ([]byte, error) {
		return c.Output()
	}
)

func _gitCommit(fileName string) error {
	c := execCommand("git", "commit", "-F", fileName, "-e")
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return commandRun(c)
}

func _lastCommitMessage() (string, error) {
	c := execCommand("git", "log", "-1", "--pretty='%B'")
	out, err := commandOutput(c)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func _gitStatus() ([]string, error) {
	var diffFiles []string

	c := execCommand("git", "status", "--porcelain")
	out, err := commandOutput(c)
	if err != nil {
		return nil, err
	}

	outStr := string(out)
	scanner := bufio.NewScanner(strings.NewReader(outStr))
	for scanner.Scan() {
		splitStatus := strings.Split(scanner.Text(), " ")
		diffFiles = append(diffFiles, splitStatus[len(splitStatus)-1])
	}

	return diffFiles, nil
}

func _gitDiff(fileName string) (string, error) {
	c := execCommand("git", "diff", "--no-color", "--", fileName)
	out, err := commandOutput(c)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func _gitAdd(files []string) error {
	args := append([]string{"add"}, files...)
	c := execCommand("git", args...)
	err := commandRun(c)
	if err != nil {
		return err
	}

	return nil
}
