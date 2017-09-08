package syncvendor

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "sync-vendor [-b branch] ORIGINATING_REPO TARGET_REPO DIRECTORY",
	Short: "filters the original repo by directory and pushes to target",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting")
	},
}
