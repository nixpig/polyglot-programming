package projector

import (
	"fmt"
	"os"
	"path"
)

type Operation = int

const (
	Print Operation = iota
	Add
	Remove
)

type Config struct {
	Args      []string
	Operation Operation
	Pwd       string
	Config    string
}

func getArgs(opts *Opts) ([]string, error) {
	if len(opts.Args) == 0 {
		return []string{}, nil
	}

	operation, err := getOperation(opts)
	if err != nil {
		return nil, err
	}

	if operation == Print {
		if len(opts.Args) > 1 {
			return []string{}, fmt.Errorf("Expected 0 or 1 arguments but got %v", len(opts.Args))
		}

		return opts.Args, nil
	}

	if operation == Add {
		if len(opts.Args) != 3 {
			return []string{}, fmt.Errorf("Expected 2 arguments but got %v", len(opts.Args)-1)
		}

	}

	if operation == Remove {
		if len(opts.Args) != 2 {
			return []string{}, fmt.Errorf("Expected 2 arguments but got %v", len(opts.Args)-1)
		}
	}

	return opts.Args[1:], nil

}

func getOperation(opts *Opts) (Operation, error) {
	if len(opts.Args) == 0 {
		return 0, fmt.Errorf("expected 1 or more arguments but got %v", len(opts.Args))
	}
	if opts.Args[0] == "add" {
		return Add, nil
	}

	if opts.Args[0] == "remove" {
		return Remove, nil
	}

	return Print, nil
}

func getConfig(opts *Opts) (string, error) {
	if opts.Config != "" {
		return opts.Config, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, ".config", "projector", "projector.json"), nil
}

func getPwd(opts *Opts) (string, error) {
	if opts.Pwd != "" {
		return opts.Pwd, nil
	}

	return os.Getwd()
}

func NewConfig(opts *Opts) (*Config, error) {
	operation, err := getOperation(opts)
	if err != nil {
		return nil, err
	}

	args, err := getArgs(opts)
	if err != nil {
		return nil, err
	}

	config, err := getConfig(opts)
	if err != nil {
		return nil, err
	}

	pwd, err := getPwd(opts)
	if err != nil {
		return nil, err
	}

	return &Config{
		Args:      args,
		Operation: operation,
		Config:    config,
		Pwd:       pwd,
	}, nil

}
