package invoice

import (
	"fmt"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "list & filter invoices from dinero",
		Example: `todo`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("hello world")

			return nil
		},
	}

	return cmd
}
