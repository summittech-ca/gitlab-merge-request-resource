package in

import (
	"fmt"
	"os"
	"os/exec"
)

type GitRunner interface {
	Run(args ...string) error
}

func NewRunner() GitRunner {
	return DefaultRunner{}
}

type DefaultRunner struct {
}

func (r DefaultRunner) Run(args ...string) error {
	cmd := "git"
	command := exec.Command(cmd, args...)
	fmt.Print(cmd)
	fmt.Print(" ")
	fmt.Println(args)

	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		return err
	}
	return nil
}
