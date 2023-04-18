package cmd

import (
	"os"
		
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/laforge-tech/bseal/pkg/bsio"
)

func exportCmd() *cobra.Command {
	var configuration bsio.Configuration

	return &cobra.Command{
		Use: "export [local_archive] remote_archive",

		Short: "Export an archive to the backup store",

		Args: cobra.RangeArgs(1,2),

		PreRunE: func(cmd *cobra.Command, args []string) error {
			return viper.Unmarshal(&configuration)
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				// Input Stream 
				f *os.File

				// Archive Name
				archiveName string
			)

			if len(args) == 2 {
				var err error

				f, err = os.Open(args[0]);
				if err != nil {
					return err
				}
				defer f.Close()
			} else {
				f = os.Stdin
			}
		
			return bsio.NewExporter(configuration).Export(f, archiveName)
		},
	}
}
