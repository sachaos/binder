package cmd

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/segmentio/textio"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/exec"
	"sync"
)

var (
	colorCnt = 0
	colors   = []aurora.Color{
		aurora.RedFg,
		aurora.GreenFg,
		aurora.BlueFg,
		aurora.CyanFg,
		aurora.MagentaFg,
		aurora.YellowFg,
	}
)

func getColor() aurora.Color {
	color := colors[colorCnt]

	colorCnt++
	if colorCnt >= len(colors) {
		colorCnt = 0
	}

	return aurora.BoldFm | color
}

type Bind struct {
	Cmd    *exec.Cmd
	Prefix string
}

func Run(cmd *cobra.Command, args []string) error {
	sources, err := ParseSource(os.Stdin)
	if err != nil {
		return err
	}

	maxL := 0
	for _, source := range sources {
		l := len(source.Name)
		if l > maxL {
			maxL = l
		}
	}

	binds := make([]*Bind, len(sources))
	for i, source := range sources {
		color := getColor()
		prefix := aurora.Colorize(fmt.Sprintf("%*s", maxL, source.Name)+"| ", color).String()

		cmd := exec.Command("sh", "-c", source.Command)
		cmd.Stdout = textio.NewPrefixWriter(os.Stdout, prefix)
		cmd.Stderr = textio.NewPrefixWriter(os.Stderr, prefix)

		binds[i] = &Bind{
			Cmd:    cmd,
			Prefix: prefix,
		}
	}

	var wg sync.WaitGroup
	for _, bind := range binds {
		err := bind.Cmd.Start()
		if err != nil {
			return fmt.Errorf("start cmd: %w", err)
		}

		wg.Add(1)
		go func(bind *Bind) {
			defer wg.Done()
			err := bind.Cmd.Wait()
			if err != nil {
				io.WriteString(os.Stderr, bind.Prefix + err.Error())
			}
		}(bind)
	}

	wg.Wait()
	return nil
}
