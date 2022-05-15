package cli

import (
	"github.com/spf13/cobra"
)

func (c *Client) rootCmd() *cobra.Command {
	return &cobra.Command{
		Use: "commander",
	}
}

func (c *Client) Run() {
	root := c.rootCmd()
	root.AddCommand(c.commander())

	root.Execute()
}
