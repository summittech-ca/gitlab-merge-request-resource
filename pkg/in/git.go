package in

import (
	"log"
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
	l := log.New(os.Stderr, "", 0)
	l.Print(cmd)
	l.Print(" ")
	l.Println(args)

	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		return err
	}
	return nil
}
