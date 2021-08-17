package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func getCurrentAbPathByCaller() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	return path, nil
}
func main() {
	_, _, err := ExecuteCommandFile("./test.sh")
	fmt.Println(err)
}

func ExecuteCommandFile(name string) (stdoutBuf, stderrBuf bytes.Buffer, err error) {
	cmd := exec.Command("/bin/bash", "-c", name)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err = cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()
	err = cmd.Wait()
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	return stdoutBuf, stderrBuf, err
}
