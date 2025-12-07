package cmd

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

//go:embed prompt.md
var agentPrompt string

var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Output instructions for AI coding agents",
	Long:  `Outputs a prompt that instructs AI coding agents on how to use the beans CLI to manage project issues.`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print(agentPrompt)

		// Append dynamic issue types section from config
		if cfg != nil && len(cfg.Types) > 0 {
			var sb strings.Builder
			sb.WriteString("\n## Issue Types\n\n")
			sb.WriteString("This project has the following issue types configured. Always specify a type with `-t` when creating beans:\n\n")
			for _, t := range cfg.Types {
				if t.Description != "" {
					sb.WriteString(fmt.Sprintf("- **%s**: %s\n", t.Name, t.Description))
				} else {
					sb.WriteString(fmt.Sprintf("- **%s**\n", t.Name))
				}
			}
			sb.WriteString("\n")
			fmt.Print(sb.String())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)
}
