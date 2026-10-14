package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tomato2718/docker-registry-manager/internal/httpclient"
)

var lsCmd = &cobra.Command{
	Use:   "ls <registry url>",
	Short: "List all images",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res, err := httpclient.Get(args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		for _, tag := range res {
			fmt.Println(tag)
		}

	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
