package main

import (
	"fmt"
	"os"

	"github.com/deads2k/sync-vendor/pkg/cmd/syncvendor"
	_ "github.com/spf13/cobra"
	_ "gopkg.in/src-d/go-git.v4"
)

func main() {
	if err := syncvendor.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
