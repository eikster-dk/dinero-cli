package invoice

import (
	listCmd "github.com/eikc/dinero-cli/internal/cmd/invoice/list"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "invoice <command>",
		Short:   "Manage invoices",
		Long:    "Work with invoices in dinero",
		Example: `todo`,
	}

	cmd.AddCommand(listCmd.New())

	return cmd
}
