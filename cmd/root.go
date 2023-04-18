package cmd

import (
	"github.com/pterm/pterm"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func rootCmd() *cobra.Command {
	var configFile string

	var cmd = &cobra.Command{

		Use: "bseal",
		Short: "An encrypted backup tool",
	
		// No usage on error
		SilenceUsage: true,

		// We re doing the error report
		SilenceErrors: true,

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			viper.SetConfigName("config")
			viper.SetConfigType("yaml")
			viper.AddConfigPath(".")
			viper.AddConfigPath("$HOME/.bseal")
			viper.AddConfigPath("/etc/bseal")

			viper.SetEnvPrefix("bseal")
			viper.AutomaticEnv()

			if configFile != "" {
				viper.SetConfigFile(configFile)
			}

			err := viper.ReadInConfig()
			if err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); ok && configFile == "" {
					// We re fine with no config at all, as long as we did not 
					// specify an explicit config filename
					err = nil
				}
			}

			// Configure pterm
			pterm.Info.Prefix.Text = "INFO    "
			pterm.Warning.Prefix.Text = "WARNING "
			pterm.Error.Prefix.Text = "ERROR   "
			pterm.Success.Prefix.Text = "SUCCESS "

			return err
		},
	}

	// Disable completion command
	cmd.CompletionOptions.DisableDefaultCmd = true
	
	cmd.PersistentFlags().StringVarP(&configFile, "config", "c" , "", "config file")

	return cmd
}

func Execute() {
	cmd := rootCmd()

	cmd.AddCommand(
		importCmd(),
		exportCmd(),
	)

	if err := cmd.Execute() ; err != nil {
		pterm.Error.Printf("%s\n", err)
	}
}