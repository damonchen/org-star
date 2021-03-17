package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/damonchen/org-star/internal/config"
	"github.com/damonchen/org-star/internal/model"
)

var (
	cfgFile string
)

func debug(format string, v ...interface{}) {
	format = fmt.Sprintf("[DEBUG]: %s\n", format)
	_, _ = fmt.Fprintf(os.Stderr, format, v...)
}

func warning(format string, v ...interface{}) {
	format = fmt.Sprintf("[WARNING]: %s\n", format)
	_, _ = fmt.Fprintf(os.Stderr, format, v...)
}

func newRootCmd(args []string) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:          "osv",
		Short:        "oss server for aliyun and tencent oss, etc.",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load(cfgFile)
			if err != nil {
				return err
			}

			// gorm的db处理
			db, err := model.DBInit(&cfg.DB)
			if err != nil {
				return err
			}

			// 迁移 schema
			err = db.AutoMigrate(&model.User{}, &model.Group{},
				&model.Favorite{}, &model.GitHubUser{})
			if err != nil {
				return err
			}

			fmt.Println("migrate db done")
			return nil
		},
	}

	flags := cmd.PersistentFlags()
	flags.ParseErrorsWhitelist.UnknownFlags = true

	flags.StringVar(&cfgFile, "cfg", "", "config filename path")

	err := flags.Parse(args)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func main() {
	cmd, err := newRootCmd(os.Args[1:])
	if err != nil {
		warning("create root command error %s", err)
		return
	}

	if err = cmd.Execute(); err != nil {
		debug("%+v", err)
		os.Exit(1)
	}
}
