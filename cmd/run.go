package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/segmentio/textio"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
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

func Run(cmd *cobra.Command, args []string) error {
	sources, err := ParseSource(os.Stdin)
	if err != nil {
		return err
	}

	eg := errgroup.Group{}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals)

	maxL := 0
	for _, source := range sources {
		l := len(source.Name)
		if l > maxL {
			maxL = l
		}
	}

	cmds := make([]*exec.Cmd, len(sources))
	for i, source := range sources {
		cmdSlice := strings.Split(source.Command, " ")

		cmd := exec.Command(cmdSlice[0], cmdSlice[1:]...)
		cmd.Stdout = textio.NewPrefixWriter(os.Stdout, aurora.Colorize(fmt.Sprintf("%*s", maxL, source.Name)+"| ", getColor()).String())
		cmd.Stderr = textio.NewPrefixWriter(os.Stderr, aurora.Colorize(fmt.Sprintf("%*s", maxL, source.Name)+"| ", getColor()).String())

		err := cmd.Start()
		if err != nil {
			return err
		}

		cmds[i] = cmd

		eg.Go(func() error {
			return cmd.Wait()
		})
	}

	s := <-signals
	for _, cmd := range cmds {
		err = cmd.Process.Signal(s)
		if err != nil {
			return err
		}
	}

	return eg.Wait()
}
