package main

import (
	goflags "flag"

	"github.com/spf13/cobra"
)

const (
	defaultAssistantServerUrl = "http://localhost:9999"
)

var (
	server string
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "acli [usage]",
		Short: "Command line tools for assistant",
		Long:  `acli is a command tool for assistant that can be used to call assistant apis`,
	}
	fs := goflags.NewFlagSet("", goflags.PanicOnError)

	cmd.PersistentFlags().AddGoFlagSet(fs)

	cmd.PersistentFlags().StringVar(&server, "server", defaultAssistantServerUrl, "asisstant server url")

	cmd.AddCommand(GenUnitTests())

	return cmd
}

func main() {
	if err := NewCmd().Execute(); err != nil {
		panic(err)
	}
}
