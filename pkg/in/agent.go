package in

import (
	"os"
	"os/exec"
	"strings"
)

type AgentRunner interface {
	Start() error
	AddKey(key string) error
}

func NewAgentRunner() AgentRunner {
	return &AgentRunnerImpl{
		sockPath: "/tmp/ssh-agent.sock",
	}
}

type AgentRunnerImpl struct {
	sockPath string
	agent    *exec.Cmd
}

func (r *AgentRunnerImpl) Start() error {
	if r.agent != nil {
		return nil // already running
	}
	agent := exec.Command("ssh-agent", "-a", r.sockPath)
	agent.Stdin = os.Stdin
	agent.Stderr = os.Stderr
	err := agent.Run()
	if err != nil {
		return err
	}
	r.agent = agent
	os.Setenv("SSH_AUTH_SOCK", r.sockPath)
	return nil
}

func (r AgentRunnerImpl) AddKey(key string) error {
	println("AgentRunnerImpl::AddKey: " + key)
	command := exec.Command("ssh-add", "-")
	// stdin, err := command.StdinPipe()
	command.Stderr = os.Stderr
	// command.Stdin = bytes.NewBuffer([]byte(key))
	command.Stdin = strings.NewReader(key + "\n")
	// if err != nil {
	// 	io.WriteString(stdin, key)
	// }
	// stdin.Close()
	// var b bytes.Buffer
	// b.Write([]byte(key))
	// command.Stdin = &b
	err := command.Run()
	if err != nil {
		return err
	}
	return nil
}
