package vouchers

import (
	"github.com/eikc/dinero-cli/internal/cmd/vouchers/list"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "vouchers <command>",
		Short:   "Work with vouchers (purchases) in dinero regnskab",
		Example: `todo`,
	}

	cmd.AddCommand(list.New())

	return cmd
}
