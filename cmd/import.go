package cmd

import(
	"github.com/spf13/cobra"
)

func importCmd() *cobra.Command {
	return &cobra.Command{
		Use: "import remote_archive [local_archive]",

		Short: "Import an archive from the backup store",

		Args: cobra.RangeArgs(1,2),
		
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
