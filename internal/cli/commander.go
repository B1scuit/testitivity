package cli

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func (c *Client) commander() *cobra.Command {
	return &cobra.Command{
		Use:   "commander",
		Short: "Run the commander service",
		RunE: func(cmd *cobra.Command, args []string) error {
			c.log.Info("Executing Commander...")

			res, err := c.core.RunCommander()
			if err != nil {
				return err
			}

			c.log.Info("Found items", zap.Int("items", len(res.Items)))

			for _, item := range res.Items {
				c.log.Sugar().Infof("%+v", item.Spec)
			}

			return nil
		},
	}
}
