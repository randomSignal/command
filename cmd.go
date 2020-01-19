package command

import (
	"os"
)

type cmdFunc func() error

type command struct {
	funcMap map[string]cmdFunc
}

func NewCommand() *command {
	return &command{
		funcMap: make(map[string]cmdFunc),
	}
}

func (cmd *command) Register(cmdString string, cmdFunc cmdFunc) {
	// check repeat
	if _, ok := cmd.funcMap[cmdString]; ok {
		panic("register repeat")
	}
	cmd.funcMap[cmdString] = cmdFunc
}

func (cmd *command) Run() error {
	if len(os.Args) != 2 {
		return nil
	}

	cmdString := os.Args[1]
	// check registered
	if _, ok := cmd.funcMap[cmdString]; !ok {
		panic("not register")
	}

	return cmd.funcMap[cmdString]()
}
