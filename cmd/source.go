package cmd

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
)

type Source struct {
	Name    string
	Command string
}

func ParseSource(r io.Reader) ([]*Source, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	cmdStrs := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	sources := make([]*Source, len(cmdStrs))
	for i, cmdStr := range cmdStrs {
		cmdSplit := strings.SplitN(cmdStr, ":", 2)

		var name, cmd string
		switch len(cmdSplit) {
		case 1:
			cmd = strings.TrimSpace(cmdSplit[0])
			name = cmd
		case 2:
			cmd = strings.TrimSpace(cmdSplit[1])
			name = strings.TrimSpace(cmdSplit[0])
		default:
			return nil, errors.New("invalid command syntax")
		}

		sources[i] = &Source{
			Name: name,
			Command: cmd,
		}
	}

	return sources, nil
}
