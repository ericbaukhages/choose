package cmd

import (
	"fmt"
	"os"

	"github.com/febvigrail/choose/choose"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "choose",
	Short: "Tmux wrapper and project manager",
	Long: `This project is a work in progress.

I use this to encompass my scripting and project management tools
used with tmux and vim.

If run with no options, the choosing UI will run.`,
	Run: func(cmd *cobra.Command, args []string) {
		configFileName, err := homedir.Expand("~/.tmux.sessions.log")
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
		}

		config := choose.Config{
			Location: configFileName,
		}
		config.Parse()

		ui := choose.Interface{
			Config: config,
		}

		name, err := ui.Run()
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
			return
		}

		_ = name

		path := config.Values[name]

		if path == "" {
			fmt.Printf("No session called: %s. Exiting.\n", name)
			return
		}

		session := choose.Session{
			Path:    path,
			Session: name,
		}

		_, err = session.Start()
		if err != nil {
			fmt.Printf("Session failed: %v\n", err)
			return
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.choose.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".choose" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".choose")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
