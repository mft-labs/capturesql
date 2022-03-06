package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"os/exec"
)

func isFileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}

func StreamOutput(src io.Reader, dst *bytes.Buffer, out io.Writer) error {
	mw := io.MultiWriter(dst)
	if out!=nil {
		mw = io.MultiWriter(dst, out)
	}
	s := bufio.NewReader(src)
	for {
		var line []byte
		line, err := s.ReadSlice('\n')
		if err == io.EOF && len(line) == 0 {
			break // done
		}
		if err == io.EOF {
			return fmt.Errorf("Improper termination: %v", line)
		}
		if err != nil {
			return err
		}
		mw.Write(line)
	}

	return nil
}

func RunCommand(app string,params []string, wg *sync.WaitGroup, verbose bool ) (text string, err error) {
	defer wg.Done()
	var wg2 sync.WaitGroup
	wg2.Add(1)
	cmd := exec.Command(app, params...)
	cmd.Env = append(os.Environ(),"")
	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()
	err = cmd.Start()
	if err != nil {
		wg2.Done()
		return "",err
	}
	var stdout, stderr bytes.Buffer
	go func() {
		if verbose {
			StreamOutput(stdoutPipe, &stdout, os.Stdout)
		} else {
			StreamOutput(stdoutPipe, &stdout,nil)
		}

		wg2.Done()
	}()
	wg2.Add(1)
	go func() {
		StreamOutput(stderrPipe, &stderr, os.Stderr)
		wg2.Done()
	}()
	wg2.Wait()
	err = cmd.Wait()
	if err != nil {
		return "",err
	}
	text = stdout.String()
	if strings.Contains(text,"Usage of "){
		return "", fmt.Errorf("Invalid usage")
	}
	if strings.Contains(text,"Error '1' "){
		return "", fmt.Errorf("Error occurred")
	}

	errStr :=  stderr.String()
	if len(errStr) >0 {
		return "",fmt.Errorf("Error:\n%s\n", errStr)
	}
	return text,nil
}
