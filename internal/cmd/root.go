package cmd

import (
	"github.com/eikc/dinero-cli/internal/cmd/invoice"
	"github.com/eikc/dinero-cli/internal/cmd/vouchers"
	"github.com/spf13/cobra"
)

func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dinero <command> <subcommand> [flags]",
		Short: "dinero CLI",
		Long:  `Work seamlessly with dinero regnskab from the command line.`,

		SilenceErrors: true,
		SilenceUsage:  true,
		Example:       `dinero cli`,
	}

	cmd.AddCommand(invoice.NewCmd())
	cmd.AddCommand(vouchers.NewCmd())

	return cmd
}
