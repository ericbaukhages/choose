package cmd

import (
	"fmt"

	"github.com/ericbaukhages/choose/choose"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open an existing project if possible",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		configFileName, err := homedir.Expand("~/.tmux.sessions.log")
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
		}

		config := choose.Config{
			Location: configFileName,
		}
		config.Parse()

		var name string

		if len(args) == 1 {
			name = args[0]
		} else {
			ui := choose.Interface{
				Config: config,
			}
			name, err = ui.Run()

			if err != nil {
				fmt.Printf("Prompt failed: %v\n", err)
				return
			}
		}

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

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
