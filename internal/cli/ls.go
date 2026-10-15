package cli

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
	"github.com/tomato2718/docker-registry-manager/internal/httpclient"
)

var lsCmd = &cobra.Command{
	Use:   "ls <registry url>",
	Short: "List all images",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		images, err := httpclient.Get(args[0])
		sort.Strings(images)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		for _, image := range images {
			fmt.Println(image)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
