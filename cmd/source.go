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
		if len(cmdSplit) != 2 {
			return nil, errors.New("invalid command syntax")
		}

		name := strings.TrimSpace(cmdSplit[0])
		cmd := strings.TrimSpace(cmdSplit[1])

		sources[i] = &Source{
			Name: name,
			Command: cmd,
		}
	}

	return sources, nil
}
