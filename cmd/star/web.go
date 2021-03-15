package main

import (
	"github.com/damonchen/org-star/cmd/star/require"
	"github.com/damonchen/org-star/internal/action"
	"github.com/damonchen/org-star/internal/config"
	"github.com/spf13/cobra"
	"io"
)

func newWebCmd(cfg *config.Configuration, outer io.Writer, args []string) *cobra.Command {
	server := action.Web{}
	cmd := &cobra.Command{
		Use:   "web",
		Short: "web server",
		Long:  "",
		Args:  require.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := server.Run()
			return err
		},
	}

	f := cmd.Flags()
	f.StringVar(&server.CfgFile, "cfg", "", "config filename path")

	return cmd
}
