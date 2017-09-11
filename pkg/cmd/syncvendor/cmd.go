package syncvendor

import (
	"fmt"

	"io"

	"github.com/spf13/cobra"
)

type SyncBranch struct {
	FromRepo   string
	FromBranch string
	ToRepo     string
	ToBranch   string

	Out    io.Writer
	ErrOut io.Writer
}

func NewCmdSyncBranch(out io.Writer, errOut io.Writer) *cobra.Command {
	o := &SyncBranch{
		Out:    out,
		ErrOut: errOut,
	}

	cmd := &cobra.Command{
		Use:   "sync-vendor --from-repo repo --from-branch branch --to-repo repo --to-branch branch",
		Short: "filters the original repo by directory and pushes to target",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run()
		},
	}

	cmd.Flags().StringVar(&o.FromRepo, "from-repo", o.FromRepo, "source repository")
	cmd.Flags().StringVar(&o.FromBranch, "from-branch", o.FromBranch, "source branch")
	cmd.Flags().StringVar(&o.ToRepo, "to-repo", o.ToRepo, "destination repository")
	cmd.Flags().StringVar(&o.ToBranch, "to-branch", o.ToBranch, "destination branch")

	return cmd
}

func (s *SyncBranch) Run() error {
	fmt.Printf("Starting\n")

	return nil
}
