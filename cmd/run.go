package cmd

import (
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/segmentio/textio"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func Run(cmd *cobra.Command, args []string) error {
	sources, err := ParseSource(os.Stdin)
	if err != nil {
		return err
	}

	eg := errgroup.Group{}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals)

	for _, source := range sources {
		cmdSlice := strings.Split(source.Command, " ")

		cmd := exec.Command(cmdSlice[0], cmdSlice[1:]...)
		cmd.Stdout = textio.NewPrefixWriter(os.Stdout, source.Name + " |")
		cmd.Stderr = textio.NewPrefixWriter(os.Stderr, source.Name + " |")

		err := cmd.Start()
		if err != nil {
			return err
		}

		eg.Go(func() error {
			return cmd.Wait()
		})

		eg.Go(func() error {
			s := <-signals
			return cmd.Process.Signal(s)
		})
	}

	return eg.Wait()
}
