package server

import (
	"fmt"
	"os/exec"
)

func prettyRun(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	fmt.Println(cmd.String())
	return cmd.Run()
}

func prettyRunOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	fmt.Println(cmd.String())
	bytes, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
